package entity

type TrendData struct {
	InnerData  TrendInnerData `json:"data"`
	StatusCode string         `json:"status_code"`
	StatusDesc string `json:"status_desc"`
	ReceivedTimestamp string `json:"received_timestamp"`
	ResponseTimestamp string `json:"response_timestamp"`
}

type TrendCrashInfo struct {
	Date string `json:"date"`              //日期
	CrashNum string `json:"crash_num"` //崩溃次数
	CrashUser string `json:"crash_user"` //崩溃影响用户数
	AccessUser string `json:"access_user"` //联网用户数
	AccessNum string `json:"access_num"`   //联网次数
	AppVersion string `json:"app_version"` //版本
}

type TrendInnerData struct {
	AppId string `json:"appId"`
	PlatformId string      `json:"platformId"`
	Infos []TrendCrashInfo `json:"crashInfos"`
}

type Trend struct {
	RtCode int32 `json:"rtcode"`
	Msg  string    `json:"msg"`
	Data TrendData `json:"data"`
}
