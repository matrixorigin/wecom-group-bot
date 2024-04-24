package wecom_group_bot

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

type BytesWriter struct {
	buff []byte
}

func (w *BytesWriter) Write(p []byte) (n int, err error) {
	w.buff = append(w.buff, p...)
	return len(p), nil
}

func Test_ImageSender(t *testing.T) {
	content, err := os.ReadFile("/Users/jia/Desktop/repos/guguducken/wecom-group-bot/cat.jpg")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.FailNow()
	}

	sender := NewImageSender(&Image{
		Base64: base64.StdEncoding.EncodeToString(content),
		Md5:    fmt.Sprintf("%x", md5.Sum(content)),
	})
	sender.InitFromENV()
	err = sender.Send()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.FailNow()
	}
}
