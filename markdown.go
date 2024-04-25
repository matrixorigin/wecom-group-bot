package wecom_group_bot

import (
	"errors"
	"fmt"
	"strings"
)

type MarkdownMessage struct {
	id       int32
	Msgtype  string    `json:"msgtype,omitempty"`
	Markdown *Markdown `json:"markdown,omitempty"`
}

func (m *MarkdownMessage) SetType(messageType string) {
	m.Msgtype = messageType
}

func (m *MarkdownMessage) GetType() string {
	return m.Msgtype
}

func (m *MarkdownMessage) DeepCopy() Messager {
	return &MarkdownMessage{
		id:      m.id,
		Msgtype: m.Msgtype,
		Markdown: &Markdown{
			Content: strings.Clone(m.Markdown.Content),
		},
	}
}

func (m *MarkdownMessage) Validate() error {
	if m.Msgtype != MarkdownType {
		return errors.Join(ErrInvalidType, errors.New("need is "+MarkdownType+" but got "+m.Msgtype))
	}
	return m.Markdown.Validate()
}

func (m *MarkdownMessage) SetID(id int32) {
	m.id = id
}

func (m *MarkdownMessage) GetID() int32 {
	return m.id
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
