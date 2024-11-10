package internal

import (
	"net/http"

	"github.com/matrixorigin/wecom-group-bot/internal/utils/requestutils"
	"github.com/matrixorigin/wecom-group-bot/internal/wberr"
)

func Get(url string, header http.Header, options ...requestutils.RequestOption) error {
	resp, err := requestutils.Request("GET", url, header, nil, options...)
	if err != nil {
		return err
	}
	defer requestutils.ReleaseResponse(resp)
	return wberr.NewErrorFromBytes(resp.Body())
}

func Delete(url string, header http.Header, body []byte, options ...requestutils.RequestOption) error {
	resp, err := requestutils.Request("DELETE", url, header, body, options...)
	if err != nil {
		return err
	}
	defer requestutils.ReleaseResponse(resp)
	return wberr.NewErrorFromBytes(resp.Body())
}

func Post(url string, header http.Header, body []byte, options ...requestutils.RequestOption) error {
	resp, err := requestutils.Request("POST", url, header, body, options...)
	if err != nil {
		return err
	}
	defer requestutils.ReleaseResponse(resp)
	return wberr.NewErrorFromBytes(resp.Body())
}

func Put(url string, header http.Header, body []byte, options ...requestutils.RequestOption) error {
	resp, err := requestutils.Request("PUT", url, header, body, options...)
	if err != nil {
		return err
	}
	defer requestutils.ReleaseResponse(resp)
	return wberr.NewErrorFromBytes(resp.Body())
}

func Patch(url string, header http.Header, body []byte, options ...requestutils.RequestOption) error {
	resp, err := requestutils.Request("PATCH", url, header, body, options...)
	if err != nil {
		return err
	}
	defer requestutils.ReleaseResponse(resp)
	return wberr.NewErrorFromBytes(resp.Body())
}
