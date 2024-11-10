package wberr

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/matrixorigin/wecom-group-bot/internal/utils/poolutils"
)

type Error struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
}

func (e Error) Error() string {
	builder := poolutils.GetStringBuilder()
	defer poolutils.PutStringBuilder(builder)

	builder.WriteString(`{`)
	builder.WriteString(`"errmsg":"` + e.Message + `",`)
	builder.WriteString(`"errcode":` + strconv.Itoa(e.Code) + `}`)
	return builder.String()
}

func NewError(code int, message string, options ...Option) error {
	e := &Error{
		Code:    code,
		Message: message,
	}
	for _, opt := range options {
		opt(e)
	}
	return e
}

func NewErrorFromError(err error) error {
	if err == nil {
		return nil
	}
	var e *Error
	if !errors.As(err, &e) {
		e = &Error{
			Code:    -1,
			Message: err.Error(),
		}
	}
	return e
}

func NewErrorFromBytes(bytes []byte) error {
	if len(bytes) == 0 {
		return nil
	}
	e := &Error{}
	err := json.Unmarshal(bytes, e)
	if err != nil {
		e.Code = -1
		e.Message = "data: " + string(bytes) + ", error: " + err.Error()
	}
	if errors.Is(e, ErrOK) {
		return nil
	}
	return e
}

func NewErrorFromString(str string) error {
	return NewErrorFromBytes([]byte(str))
}

func OverrideError(err error, options ...Option) error {
	var e *Error
	if !errors.As(err, &e) {
		return NewErrorFromError(err)
	}
	for _, opt := range options {
		opt(e)
	}
	return e
}

func (e Error) Is(target error) bool {
	var ne *Error
	if errors.As(target, &ne) {
		return ne.Code == e.Code
	}
	return false
}
