package pipleline

import (
	"base.bugly/api"
	"base.bugly/entity"
	log "base.bugly/pkg/plog"
	"base.bugly/pkg/robot"
	utime "base.bugly/pkg/util"
	"fmt"
	"strconv"
)

func QueryBuglyAndNotify(appName string, AppKey string, AppID string, platformId int, dingTalkToken string) {
	var queryDate = utime.QueryYesterdayTime()
	crashInfos := api.RequestDayOfTrend(AppKey, AppID, queryDate, platformId)
	if crashInfos == nil || len(crashInfos) <= 0 {
		log.Errorf("queryBuglyAndNotify error AppID(%v)", AppID)
		return
	}

	//整理消息格式 markDown格式
	var dailyContent = "### " + appName + " " + queryDate + "日报\n"

	for _, info := range crashInfos {
		dailyContent += "### >>" + info.AppVersion + "<<\n"
		dailyContent += "#### 概览(崩溃率:" + crashRate(info) + ")\n"
		dailyContent += "* 崩溃次数:" + info.CrashNum + "\n"
		dailyContent += "* 崩溃用户数:" + info.CrashUser + "\n"
		dailyContent += "* 联网次数:" + info.AccessNum + "\n"
		dailyContent += "* 联网用户数:" + info.AccessUser + "\n"
		TopIssueByVersion(AppKey, AppID, platformId, info.AppVersion, &dailyContent)
	}

	fmt.Printf("Last == >\n%v", dailyContent)

	robotPoxy := robot.NewRobotPoxy(dingTalkToken, "", robot.DingDing)

	err := robotPoxy.SendMarkdownMessage(appName+"日报", dailyContent, []string{}, false)
	if err != nil {
		return
	}

}

func crashRate(info entity.TrendCrashInfo) string {
	//strconv.FormatFloat(info.AccessNum, "f", 6, )
	numI, err2 := strconv.ParseFloat(info.CrashNum, 64)
	if err2 != nil {
		return "error"
	}
	numJ, err1 := strconv.ParseFloat(info.AccessNum, 64)
	if err1 != nil || numJ <= 0{
		return "error"
	}
	crashRate := (numI / numJ) * 100
	return strconv.FormatFloat(crashRate, 'f', 2, 64) + "%"
}

// TopIssueByVersion
// type : 1 crash ,2 anr/卡顿,3 error
// limit 查询top的量级 按照crash_user进行排序
func TopIssueByVersion(appKey string, appId string, platform int, version string, dailyContentPtr *string) {
	var queryDate = utime.QueryYesterdayTime()
	infos := api.TopIssue(appKey, appId, platform, queryDate, version)

	if infos == nil || len(infos) <= 0 {
		log.Printf("request top issue of version(%v) obtain nil", version)
		return
	}

	*dailyContentPtr += "#### Top Issue\n"

	for _, info := range infos {
		issueDetail := api.IssueDetail(appKey, appId, platform, info.IssueId)
		if len(issueDetail) > 0 {
			*dailyContentPtr += fmt.Sprintf("* 次数:%d, 用户数:%d, [Bug链接](%s)\n", int(info.CrashNum), int(info.CrashUser), issueDetail)
		} else {
			*dailyContentPtr += fmt.Sprintf("* 次数:%d, 用户数:%d\n", int(info.CrashNum), int(info.CrashUser))
		}
	}

}
