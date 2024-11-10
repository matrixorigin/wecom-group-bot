package wberr

var (
	ErrOK              = NewError(0, "ok")
	ErrUnknown         = NewError(-1, "unknown error")
	ErrRequestFailed   = NewError(-2, "http request failed")
	ErrRequestMaxTimes = NewError(-3, "http request max times")
	ErrEmptyWebhookKey = NewError(-4, "empty webhook key")
	ErrInvalidType     = NewError(-5, "invalid type")
)
