package robot

import (
	"base.bugly/pkg/robot/ding_talk"
)

// IRobotType https://oapi.dingtalk.com/robot/send?access_token=replace-ding-ding-token-here
type IRobotType int32

const (
	DingDing = iota
	Wechat
)

type IRobotPoxy struct {
	Token string
	Secret string
}

type IRobotInterface interface {

	SendMessage(msg interface{}) error

	SendMarkdownMessage(title string, text string, atMobiles []string, isAtAll bool) error
}

func NewRobotPoxy(token, secret string, rType IRobotType) IRobotInterface {
	switch rType {
	case DingDing:
		return ding_talk.NewRobot(token, secret)
	case Wechat:
		return nil
	default:
		return nil
	}
}






