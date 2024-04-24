package wecom_group_bot

import (
	"fmt"
	"testing"
)

func Test_TextSend(t *testing.T) {
	sender := NewTextSender(&Text{
		Content:             "test test test",
		MentionedList:       []string{""},
		MentionedMobileList: []string{"17864731129"},
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
