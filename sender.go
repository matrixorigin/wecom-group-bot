package wecom_group_bot

import (
	"encoding/json"
	"time"

	"github.com/matrixorigin/wecom-group-bot/utils"
)

type Sender struct {
	webhook string
	message Messager
}

func (s *Sender) InitFromENV() error {
	s.webhook = utils.MustGetEnv(WebhookEnvName)
	return nil
}

func (s *Sender) InitFromFile(path string) error {
	// TODO: generate sender from config file
	return nil
}

func (s *Sender) Send() error {
	if err := s.Validate(); err != nil {
		return err
	}

	payload, err := json.Marshal(s.message)
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
	return s.message.Validate()
}

type Messager interface {
	SetType(messageType string)
	GetType() string
	DeepCopy() Messager
	Validate() error
}
