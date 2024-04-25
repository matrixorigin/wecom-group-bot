package wecom_group_bot

import (
	"fmt"
	"testing"
	"time"

	"github.com/matrixorigin/wecom-group-bot/utils"
)

func Test_TextSend(t *testing.T) {
	message := NewTextMessage(&Text{
		Content:             "test test test",
		MentionedList:       []string{""},
		MentionedMobileList: []string{"17864731129"},
	})
	webhook := utils.MustGetEnv(WebhookEnvName)
	sender := NewSender(webhook)
	id, err := sender.Send(message)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.FailNow()
	}
	fmt.Printf("id: %d\n", id)

	for i := 0; i < 300; i++ {
		result := sender.WaitResult(id)
		fmt.Printf("result.Status: %v\n", result.Status)
		fmt.Printf("result.Err: %v\n", result.Err)
		fmt.Printf("result.Time: %v\n", result.Time)
		if result.Status == SuccessStatus {
			break
		}
		time.Sleep(10 * time.Second)
	}

}
