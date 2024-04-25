package wecom_group_bot

import (
	"encoding/json"
	"time"

	"github.com/matrixorigin/wecom-group-bot/utils"
)

const (
	NoticePrefix      = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"
	UploadMediaPrefix = "https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media"
)

type Sender struct {
	key string
}

func NewSender(key string) *Sender {
	return &Sender{
		key: key,
	}
}

func (s *Sender) Send(message Messager) error {
	if err := message.Validate(); err != nil {
		return err
	}
	if err := s.Validate(); err != nil {
		return err
	}

	url := utils.URL{
		Params: map[string]string{
			"key": s.key,
		},
	}
	switch message.GetType() {
	case MediaType:
		url.Endpoint = UploadMediaPrefix
		return utils.UploadMedia(url)
	case TextType, MarkdownType, ImageType, NewsType, FileType, VoiceType, CardType:
		url.Endpoint = NoticePrefix
		payload, err := json.Marshal(message)
		if err != nil {
			return err
		}
		_, err = utils.PostWithRetry(url, payload, 5, 10*time.Second)
		return err
	default:
		return ErrInvalidType
	}
}

func (s *Sender) Validate() error {
	if s.key == "" {
		return ErrEmptyWebhookKey
	}
	return nil
}

type Messager interface {
	SetType(messageType string)
	GetType() string
	DeepCopy() Messager
	Validate() error
}
