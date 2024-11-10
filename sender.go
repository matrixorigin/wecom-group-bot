package wecom_group_bot

import (
	"encoding/json"

	"github.com/matrixorigin/wecom-group-bot/internal"
	"github.com/matrixorigin/wecom-group-bot/internal/wberr"
	"github.com/matrixorigin/wecom-group-bot/utils"
)

type Sender struct {
	key            string
	messageSendURL string
	mediaSendURL   string
}

func NewSender(key string) *Sender {
	return &Sender{
		key:            key,
		messageSendURL: internal.GenURL(NoticePrefix, "?key=", key),
		mediaSendURL:   internal.GenURL(UploadMediaPrefix, "?key=", key),
	}
}

func (s *Sender) Send(message Messager) error {
	if err := message.Validate(); err != nil {
		return err
	}
	switch message.GetType() {
	case MessageTypeMedia:
		return utils.UploadMedia(s.mediaSendURL)
	case MessageTypeText, MessageTypeMarkdown, MessageTypeImage,
		MessageTypeNews, MessageTypeFile, MessageTypeVoice, MessageTypeCard:
		payload, err := json.Marshal(message)
		if err != nil {
			return wberr.NewErrorFromError(err)
		}
		return internal.Post(s.messageSendURL, internal.DefaultRestRequiredHeader(), payload)
	default:
		return wberr.ErrInvalidType
	}
}

func (s *Sender) Validate() error {
	if s.key == "" {
		return wberr.ErrEmptyWebhookKey
	}
	return nil
}

type Messager interface {
	SetType(messageType MessageType)
	GetType() MessageType
	DeepCopy() Messager
	Validate() error
}
