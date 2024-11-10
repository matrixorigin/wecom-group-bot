package wberr

type Option func(e *Error)

func WithMessage(message string) Option {
	return func(e *Error) {
		if message != "" {
			e.Message = message
		}
	}
}
