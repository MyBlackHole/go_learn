package main

import (
	"io/ioutil"
	"log"

	"github.com/CodyGuo/dingtalk"
)

// 主要代码
func SendDingMsg(title, text string) {
	webHook := "https://oapi.dingtalk.com/robot/send?access_token=014748917f39fdb64e3f47ee05f44b038c115fcb00b2e1281886712ad4b24186"
	// 密钥，机器人安全设置页面，加签一栏勾选之后下面显示的SEC开头的字符串
	secret := "SECdfcaacc9f7d4993dd4a94c70707b68aa8d859cdbcf0efd843fcc0f913a3f82ae"
	dt := dingtalk.New(webHook, dingtalk.WithSecret(secret))

	if err := dt.RobotSendMarkdown(title, text); err != nil {
		log.Print(err)
	}
	// printResult(dt)
}

func printResult(dt *dingtalk.DingTalk) {
	response, err := dt.GetResponse()
	if err != nil {
		log.Print(err)
	}
	reqBody, err := response.Request.GetBody()
	if err != nil {
		log.Print(err)
	}
	reqData, err := ioutil.ReadAll(reqBody)
	if err != nil {
		log.Print(err)
	}
	log.Printf("发送消息成功, message: %s", reqData)
}

func testSendDingMsg() {
	markdownTitle := "markdown"
	markdownText := "#### 杭州天气 @176XXXXXXXX\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n" +
		"> ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n" +
		"> ###### 10点20分发布 [天气](https://www.dingtalk.com)\n"
	SendDingMsg(markdownTitle, markdownText)
}
