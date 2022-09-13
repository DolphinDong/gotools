package wechat

import (
	client "github.com/DolphinDong/gotools/http"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

const (
	robotUrl = "/cgi-bin/webhook/send"
)

type RobotMessage struct {
	MsgType string `mapstructure:"msgtype"`
	Text    struct {
		Content             string   `mapstructure:"content"`
		MentionedList       []string `mapstructure:"mentioned_list"`
		MentionedMobileList []string `mapstructure:"mentioned_mobile_list"`
	} `mapstructure:"text"`
}

// 发送机器人消息
func SendRobotMessage(robotKey, weChatUrl string, message RobotMessage) error {
	c := client.Client{BaseUrl: weChatUrl}
	parm := map[string]string{"key": robotKey}
	header := map[string]string{"Content-Type": client.ContentJson}

	msg := map[string]interface{}{}
	err := mapstructure.Decode(message, &msg)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	response, err := c.Post(robotUrl, parm, header, msg)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	responseStr := string(response)
	if gjson.Get(responseStr, "errcode").Int() != 0 {
		err = errors.Errorf("send robot message  failed:%v", responseStr)
		return err
	}
	return nil
}
