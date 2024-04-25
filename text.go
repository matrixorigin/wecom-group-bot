package wecom_group_bot

import "strings"

func NewTextMessage(message *Text) Messager {
	return &TextMessage{
		Msgtype: TextType,
		Text:    message.DeepCopy(),
	}
}

type TextMessage struct {
	Msgtype string `json:"msgtype,omitempty"`
	Text    *Text  `json:"text,omitempty"`
}

func (t *TextMessage) GetType() string {
	return TextType
}

func (t *TextMessage) DeepCopy() Messager {
	dst := &TextMessage{
		Msgtype: t.Msgtype,
		Text:    t.Text.DeepCopy(),
	}
	return dst
}

func (t *TextMessage) SetType(messageType string) {
	t.Msgtype = messageType
}

func (t *TextMessage) Validate() error {
	return nil
}

type Text struct {
	Content             string   `json:"content,omitempty"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

func (t *Text) DeepCopy() *Text {
	dst := &Text{
		Content:             strings.Clone(t.Content),
		MentionedList:       make([]string, len(t.MentionedList)),
		MentionedMobileList: make([]string, len(t.MentionedMobileList)),
	}
	copy(dst.MentionedList, t.MentionedList)
	copy(dst.MentionedMobileList, t.MentionedMobileList)
	return dst
}
