package internal

import "net/http"

type HeaderOption func(http.Header)

func DefaultRestRequiredHeader(options ...HeaderOption) http.Header {
	return NewRestRequiredHeader("")
}

func NewRestRequiredHeader(
	contentType string,
	options ...HeaderOption,
) http.Header {
	if contentType == "" {
		contentType = "application/json"
	}
	header := make(http.Header, 3)

	header.Set("Content-Type", contentType)
	for _, opt := range options {
		opt(header)
	}
	return header
}

func OverrideHeader(key string, value string) HeaderOption {
	return func(header http.Header) {
		if value != "" {
			header.Set(key, value)
		}
	}
}
