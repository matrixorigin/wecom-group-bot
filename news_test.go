package wecom_group_bot

import (
	"fmt"
	"testing"
)

func Test_NewsSender(t *testing.T) {
	sender := NewNewsSender(&News{
		Articles: Articles{
			{
				Title:       "中秋节礼品领取",
				Description: "今年中秋节公司有豪礼相送",
				URL:         "www.qq.com",
				PicURL:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
			},
		},
	})
	err := sender.InitFromENV()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.FailNow()
	}
	err = sender.Send()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.FailNow()
	}
}
