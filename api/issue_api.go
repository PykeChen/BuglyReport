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
)

const (
	TopIssueUrl = "https://api.bugly.qq.com/openapi/stat/issue/top?app_id=%s&app_key=%s&platform_id=%d&date=%s&version=%s&limit=%d&type=%d"
)

// TopIssue 查询每日的top问题版
// type : 1 crash ,2 anr/卡顿,3 error
// limit 查询top的量级 按照crash_user进行排序
func TopIssue(appKey string, appId string, platform int, queryDate string, version string) []entity.IssueInfo {

	url := fmt.Sprintf(TopIssueUrl, appId, appKey, platform, queryDate, version, Limit, TypeCrash)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("TopIssue request error : %v\n", request)
		return nil
	}
	request.Header.Set("content-type", "application/json")
	request.Header.Set("cache-control", "no-cache")
	resp, err := mhttp.DoClient.Do(request)
	if err != nil {
		log.Errorf("TopIssue response error:%v\n", err)
		return nil
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("TopIssue read response body error:%v\n", err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	log.Printf("TopIssue read response:%s\n", all)
	data := &entity.TopIssue{}
	err2 := json.Unmarshal(all, data)
	if err2 != nil {
		log.Errorf("TopIssue json unmarshal error:%v\n", err2)
		return nil
	}

	var issueInfos []entity.IssueInfo
	if data.RtCode == 0 {
		//log.Printf("requestDayOfTrend response json:%+v\n", data)
		issueInfos = data.Data.Infos
		sortIssueArray(issueInfos)
		log.Printf("TopIssue last sort %+v\n", issueInfos)
	} else {
		log.Errorf("TopIssue error, rtCode(%d), rtMsg(%s) !! response:%s\n", data.RtCode, data.Msg, all)
	}

	return issueInfos
}

// sortCrashArray  返回true 代表是i排在前，排序就是ij, false表示i排在后面 ji
func sortIssueArray(crashInfos []entity.IssueInfo) {
	sort.SliceStable(crashInfos, func(i, j int) bool {
		// 使用版本号大的排序
		return sortByCrashNum(crashInfos[i], crashInfos[j])
	})
}

func sortByCrashNum(i entity.IssueInfo, j entity.IssueInfo) bool {
	return i.CrashNum >= j.CrashNum
}
