package wecom_group_bot

import "errors"

var (
	ErrEmptyWebhook = errors.New("webhook is empty")
	ErrInvalidType  = errors.New("invalid sender type")
)
