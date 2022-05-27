package api

import (
	"base.bugly/pkg/mhttp"
	log "base.bugly/pkg/plog"
	putil "base.bugly/pkg/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const (
	IssueDetailUrl = "https://api.bugly.qq.com/openapi/crash/detail?app_key=%s&platform_id=%d"
	// IssueLinkUrl 第一个appId, 第二个%s issueId
	IssueLinkUrl = "https://bugly.qq.com/v2/crash-reporting/crashes/%s/%s?pid=%d"
)

// IssueDetail 查询Issue的详细信息
// type : 1 crash ,2 anr/卡顿,3 error
// limit 查询top的量级 按照crash_user进行排序
func IssueDetail(appKey string, appId string, platform int, issueId float32) string {

	url := fmt.Sprintf(IssueDetailUrl, appKey, platform)
	requestJson := fmt.Sprintf("{\"api_version\":1,\"app_id\":\"%s\",\"issueId\":\"%f\",\"pid\":%d,\"isNeedOsTop\":false,\"isNeedVersionRatio\":false}",
		appId, issueId, platform)
	requestBody := strings.NewReader(requestJson)

	request, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		log.Errorf("IssueDetail request error : %v\n", request)
		return ""
	}
	request.Header.Set("content-type", "application/json")
	request.Header.Set("cache-control", "no-cache")
	resp, err := mhttp.DoClient.Do(request)
	if err != nil {
		log.Errorf("IssueDetail response error:%v\n", err)
		return ""
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("IssueDetail read response body error:%v\n", err)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	issDetailStr := string(all)
	//log.Printf("IssueDetail read response:%s\n", issDetailStr)
	crossIssueId := analysisCrossIssueIdByReg(issDetailStr)
	if putil.StringValid(crossIssueId) {
		return GenerateIssueLink(appId, crossIssueId, platform)
	}

	return ""

}

func analysisCrossIssueIdByReg(content string) string {
	crossReg := regexp.MustCompile(`"crossIssueId\\":(?P<bugID>\d+)`)
	if crossReg == nil {
		log.Printf("regexp err")
		return ""
	}
	//matchString := crossReg.FindString(content)
	//subNames := crossReg.SubexpNames()
	//index := crossReg.SubexpIndex("bugID")
	matchStr := crossReg.FindStringSubmatch(content)

	if matchStr != nil && len(matchStr) > 0 {
		return matchStr[1]
	}

	return ""

}

func GenerateIssueLink(appId string, crossIssueId string, platform int) string {
	//fmt.Sprintf(IssueLinkUrl, appId, strconv.FormatFloat(float64(issueId), 'f', 0, 32))
	return fmt.Sprintf(IssueLinkUrl, appId, crossIssueId, platform)
}
