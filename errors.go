package wecom_group_bot

import "errors"

var (
	ErrEmptyWebhookKey = errors.New("webhook key is empty")
	ErrInvalidType     = errors.New("invalid sender type")
)
