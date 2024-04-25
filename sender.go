package wecom_group_bot

import (
	"encoding/json"
	"time"

	"github.com/matrixorigin/wecom-group-bot/utils"
)

type Sender struct {
	webhook string
}

func NewSender(webhook string) *Sender {
	return &Sender{
		webhook: webhook,
	}
}

func (s *Sender) Send(message Messager) error {
	if err := message.Validate(); err != nil {
		return err
	}
	if err := s.Validate(); err != nil {
		return err
	}

	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}
	url := utils.URL{
		RawURL: s.webhook,
	}
	_, err = utils.PostWithRetry(url, payload, 5, 10*time.Second)
	return err
}

func (s *Sender) Validate() error {
	if s.webhook == "" {
		return ErrEmptyWebhook
	}
	return nil
}

type Messager interface {
	SetType(messageType string)
	GetType() string
	DeepCopy() Messager
	Validate() error
}
