package entity


type IssueData struct {
	Infos []IssueInfo 		`json:"data"`
	StatusCode string        `json:"status_code"`
	StatusDesc string `json:"status_desc"`
	ReceivedTimestamp string `json:"received_timestamp"`
	ResponseTimestamp string `json:"response_timestamp"`
}

type IssueInfo struct {
	Date string `json:"date"`         //日期
	Version string `json:"version"` //版本
	CrashNum float32 `json:"crash_num"` //崩溃次数
	CrashUser float32 `json:"crash_user"` //崩溃影响用户数
	IssueId float32 `json:"issue_id"` // bugId
}

type TopIssue struct {
	RtCode int32 `json:"rtcode"`
	Msg string `json:"msg"`
	Data IssueData `json:"data"`
}
