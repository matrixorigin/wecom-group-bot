package wecom_group_bot

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	"github.com/matrixorigin/wecom-group-bot/utils"
)

func Test_ImageSender(t *testing.T) {
	sender := NewSender(utils.MustGetEnv(WebhookKeyEnvName))

	content, err := os.ReadFile("./cat.jpg")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.FailNow()
	}

	message := NewImageMessage(&Image{
		Base64: base64.StdEncoding.EncodeToString(content),
		Md5:    fmt.Sprintf("%x", md5.Sum(content)),
	})
	err = sender.Send(message)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.FailNow()
	}
}
