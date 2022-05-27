package api

import (
	"base.bugly/entity"
	"base.bugly/pkg/mhttp"
	log "base.bugly/pkg/plog"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const (
	TrendUrl = "https://api.bugly.qq.com/openapi/stat/crashOnVersion/trend"
)

// RequestDayOfTrend 查询每日的崩溃趋势, 获取版本前三
// appKey 对应的appKey
func RequestDayOfTrend(appKey string, appId string, queryDate string, platform int) []entity.TrendCrashInfo {
	var url = fmt.Sprintf(TrendUrl+"?app_key=%v", appKey)
	var requestJson = fmt.Sprintf("{\"pid\":\"%d\",\"api_version\":\"1\",\"app_id\":\"%s\",\"start_date\":\"%s\",\"end_date\":\"%s\"}",
		platform, appId, queryDate, queryDate)
	requestBody := strings.NewReader(requestJson)
	request, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		log.Errorf("requestDayOfTrend request error:%v\n", err)
		return nil
	}
	request.Header.Set("content-type", "application/json")
	request.Header.Set("cache-control", "no-cache")

	resp, err := mhttp.DoClient.Do(request)
	if err != nil {
		log.Errorf("requestDayOfTrend response error:%v\n", err)
		return nil
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("requestDayOfTrend read response body error:%v\n", err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	log.Printf("requestDayOfTrend read response:%s\n", all)
	data := &entity.Trend{}
	err2 := json.Unmarshal(all, data)
	if err2 != nil {
		log.Errorf("requestDayOfTrend json unmarshal error:%v\n", err2)
		return nil
	}

	var crashInfos []entity.TrendCrashInfo
	if data.RtCode == 0 {
		//log.Printf("requestDayOfTrend response json:%+v\n", data)
		crashInfos = data.Data.InnerData.Infos
		sortCrashArray(crashInfos)

		crashInfos = filterWantedVersions(crashInfos)

		log.Printf("requestDayOfTrend last sort %+v\n", crashInfos)
	} else {
		log.Errorf("requestDayOfTrend error, rtCode(%d), rtMsg(%s) !! response:%s\n", data.RtCode, data.Msg, all)
	}
	return crashInfos

}

func filterWantedVersions(crashInfos []entity.TrendCrashInfo) []entity.TrendCrashInfo{
	if len(crashInfos) < 3 {
		return crashInfos
	}
	var limitInfo []entity.TrendCrashInfo
	// 只取一个alpha版本
	var hasAlphaVersion = false
	for _, info := range crashInfos {
		isAlpha := strings.Contains(info.AppVersion, "alpha")
		addToArray := false
		if isAlpha && !hasAlphaVersion {
			hasAlphaVersion = true
			addToArray = true
		}  else if !isAlpha{
			addToArray = true
		}
		if addToArray {
			limitInfo = append(limitInfo, info)
			if len(limitInfo) >= 3 {
				break
			}
		}
	}
	return limitInfo
}

// sortCrashArray  返回true 代表是i排在前，排序就是ij, false表示i排在后面 ji
func sortCrashArray(crashInfos []entity.TrendCrashInfo) {
	sort.SliceStable(crashInfos, func(i, j int) bool {
		// 使用版本号大的排序
		return sortByVersion(crashInfos[i], crashInfos[j])
	})
}

func sortByVersion(i entity.TrendCrashInfo, j entity.TrendCrashInfo) bool {
	var versionI = strings.Replace(strings.Replace(i.AppVersion, ".", "", -1), "_alpha", "", -1)
	var versionJ = strings.Replace(strings.Replace(j.AppVersion, ".", "", -1), "_alpha", "", -1)
	numI, err1 := strconv.Atoi(versionI)
	if err1 != nil {
		log.Errorf("sortByVersion error %v, value: %v\n", err1, versionI)
		numI = 0
		return false
	}
	numJ, err2 := strconv.Atoi(versionJ)
	if err1 != nil {
		log.Errorf("sortByVersion error %v, value: %v\n", err2, versionJ)
		numJ = 0
		return false
	}
	return numI >= numJ
}

func sortByAccessNum(i entity.TrendCrashInfo, j entity.TrendCrashInfo) bool {
	numI, err1 := strconv.Atoi(i.AccessNum)
	if err1 != nil {
		log.Errorf("sortByAccessNum error %v, value: %v\n", err1, i.AccessNum)
		numI = 0
		return false
	}
	numJ, err2 := strconv.Atoi(j.AccessNum)
	if err1 != nil {
		log.Errorf("sortByAccessNum error %v, value: %v\n", err2, j.AccessNum)
		numJ = 0
		return false
	}

	return numI >= numJ
}
