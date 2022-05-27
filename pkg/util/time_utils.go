package putil

import "time"

//QueryYesterdayTime 查询昨天的时间，以20060102的格式
func QueryYesterdayTime() string {
	return time.Now().AddDate(0, 0, -1).Format("2006-01-02")
}