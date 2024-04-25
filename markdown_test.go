package wecom_group_bot

import (
	"fmt"
	"testing"

	"github.com/matrixorigin/wecom-group-bot/utils"
)

func Test_MarkDownSender(t *testing.T) {
	sender := NewSender(utils.MustGetEnv(WebhookKeyEnvName))
	message := NewMarkdownMessage(&Markdown{
		Content: `# test test
		> ref test

## title 2
`,
	}, []string{"JiejieJia"})

	err := sender.Send(message)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.FailNow()
	}
}
