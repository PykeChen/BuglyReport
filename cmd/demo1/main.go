package main

import (
	"base.bugly/api"
	"base.bugly/pipleline"
)

const (
	AppName = "Demo"
	AppKey = "Bugly-App-key"
	AppID = "Bugly-App-ID"
	DingToken = "Ding-Ding-token"
)

func main() {
	var platformId = api.AndroidPlatform
	pipleline.QueryBuglyAndNotify(AppName, AppKey, AppID, platformId, DingToken)

}
