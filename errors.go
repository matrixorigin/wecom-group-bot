package wecom_group_bot

import (
	"errors"
	"time"
)

var (
	ErrEmptyWebhook = errors.New("webhook is empty")
	ErrInvalidType  = errors.New("invalid sender type")
	ErrNilChannel   = errors.New("nil channel")
)

const (
	SuccessStatus = iota
	FailedStatus
	BackOffStatus
)

type Result struct {
	ID     int32
	Status int
	Err    error

	Time time.Time
}
