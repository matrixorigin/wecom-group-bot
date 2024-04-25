# WeCom Group Bot

## example
~~~bash
go get github.com/matrixorigin/wecom-group-bot
~~~

~~~go
package main

import (
	wbot "github.com/matrixorigin/wecom-group-bot"
	"github.com/matrixorigin/wecom-group-bot/utils"
)

func main() {
	sender := wbot.NewSender(utils.MustGetEnv(wbot.WebhookEnvName))
	textMessage := wbot.NewTextMessage(&wbot.Text{
		Content:             "test test test",
		MentionedList:       []string{"user_id"},
		MentionedMobileList: []string{"xxxxxxxxxx"},
	})
	err := sender.Send(textMessage)
	if err != nil {
		panic(err)
	}

	message := wbot.NewNewsMessage(&wbot.News{
		Articles: wbot.Articles{
			{
				Title:       "中秋节礼品领取",
				Description: "今年中秋节公司有豪礼相送",
				URL:         "www.qq.com",
				PicURL:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
			},
		},
	})
	err = sender.Send(message)
	if err != nil {
		panic(err)
	}
}

~~~
