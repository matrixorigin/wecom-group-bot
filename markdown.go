package wecom_group_bot

import (
	"fmt"
	"strings"

	"github.com/matrixorigin/wecom-group-bot/internal/wberr"
)

func NewMarkdownMessage(message *Markdown, mentionedList []string) Messager {
	message = message.DeepCopy()
	message.SetMentionedList(mentionedList)
	return &MarkdownMessage{
		Msgtype:  MessageTypeMarkdown,
		Markdown: message,
	}
}

type MarkdownMessage struct {
	Msgtype  MessageType `json:"msgtype,omitempty"`
	Markdown *Markdown   `json:"markdown,omitempty"`
}

func (m *MarkdownMessage) SetType(messageType MessageType) {
	m.Msgtype = messageType
}

func (m *MarkdownMessage) GetType() MessageType {
	return m.Msgtype
}

func (m *MarkdownMessage) DeepCopy() Messager {
	return &MarkdownMessage{
		Msgtype: m.Msgtype,
		Markdown: &Markdown{
			Content: strings.Clone(m.Markdown.Content),
		},
	}
}

func (m *MarkdownMessage) Validate() error {
	if m.Msgtype != MessageTypeMarkdown {
		return wberr.OverrideError(wberr.ErrInvalidType,
			wberr.WithMessage("need is "+MessageTypeMarkdown.String()+" but got "+m.Msgtype.String()),
		)
	}
	return m.Markdown.Validate()
}

type Markdown struct {
	Content string `json:"content,omitempty"`
}

func (m *Markdown) DeepCopy() *Markdown {
	return &Markdown{
		Content: strings.Clone(m.Content),
	}
}

func (m *Markdown) Validate() error {
	return nil
}

func (m *Markdown) SetMentionedList(mentionedList []string) {
	if len(mentionedList) == 0 {
		return
	}
	for _, user := range mentionedList {
		m.Content += fmt.Sprintf("<@%s>", user)
	}
}
