package wecom_group_bot

import (
	"fmt"
	"testing"

	"github.com/matrixorigin/wecom-group-bot/utils"
)

func Test_TextSender(t *testing.T) {
	sender := NewSender(utils.MustGetEnv(WebhookKeyEnvName))
	message := NewTextMessage(&Text{
		Content:             "test test test",
		MentionedList:       []string{""},
		MentionedMobileList: []string{"17864731129"},
	})
	err := sender.Send(message)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.FailNow()
	}
}
