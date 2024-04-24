package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	ErrHttpResponse     = errors.New("github api server response is not 2xx")
	ErrHttpRequest      = errors.New("request failed")
	ErrHttpNewRequest   = errors.New("new http request failed")
	ErrRateLimited      = errors.New("api rate limit is reached")
	ErrInvalidMessage   = errors.New("invalid message type")
	ErrInvalidMd5       = errors.New("invalid md5 value of content")
	ErrInvalidImageSize = errors.New("invalid image size")
	ErrMissTitle        = errors.New(" missing title")
)

const (
	SuccessCode                = 0
	RateLimitedErrorCode       = 45009
	InvalidContentErrorCode    = 40008
	InvalidImageSizedErrorCode = 40009
	InvalidFileMd5ErrorCode    = 301019
	TitleMissedErrorCode       = 41016
)

type Reply struct {
	Body         []byte
	StatusCode   int
	Header       http.Header
	RateResource string
}

type Message struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

type URL struct {
	RawURL   string
	Endpoint string
	Path     string
	Params   map[string]string
}

func (u URL) toRawURL() string {
	if len(u.RawURL) != 0 {
		return u.RawURL
	}
	url := u.Endpoint

	if len(u.Path) != 0 {
		pathStart, pathEnd := 0, len(u.Path)
		if u.Path[0] == '/' {
			pathStart++
		}
		if u.Path[len(u.Path)-1] == '/' {
			pathEnd--
		}

		url += "/" + u.Path[pathStart:pathEnd]
	}

	params := ""
	for key, value := range u.Params {
		params += "&" + key + "=" + value
	}

	if len(params) != 0 {
		url += "?" + params[1:]
	}
	return url
}

func PostWithRetry(url URL, body []byte, retryTimes int, retryDuration time.Duration) (reply *Reply, err error) {
	for i := 0; i < retryTimes; i++ {
		reply, err = Post(url, body)
		if err == nil {
			break
		}
		if err != nil && !errors.Is(err, ErrHttpRequest) {
			break
		}
		time.Sleep(retryDuration)
	}
	return reply, err
}

func GetWithRetry(url URL, retryTimes int, retryDuration time.Duration) (reply *Reply, err error) {
	for i := 0; i < retryTimes; i++ {
		reply, err = Get(url)
		if err == nil {
			break
		}
		if err != nil && !errors.Is(err, ErrHttpRequest) {
			break
		}
		time.Sleep(retryDuration)
	}
	return reply, err
}

func Get(url URL) (reply *Reply, err error) {
	req, err := http.NewRequest(
		"GET",
		url.toRawURL(),
		nil,
	)
	if err != nil {
		return reply, errors.Join(ErrHttpNewRequest, err)
	}
	return resolveReply(do(req))
}

func Post(url URL, body []byte) (reply *Reply, err error) {
	req, err := http.NewRequest(
		"POST",
		url.toRawURL(),
		bytes.NewReader(body),
	)
	if err != nil {
		return reply, errors.Join(ErrHttpNewRequest, err)
	}
	return resolveReply(do(req))
}

func do(req *http.Request) (reply *Reply, err error) {
	reply = &Reply{}

	timeOut := 20 * time.Second

	// do http request
	resp, err := newHttpClient(timeOut).Do(req)
	if err != nil {
		return reply, errors.Join(ErrHttpRequest, err)
	}

	// read http status code from resp
	reply.StatusCode = resp.StatusCode

	// read headers from resp
	reply.Header = resp.Header

	// read reply from resp
	reply.Body, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return reply, err
}

func newHttpClient(timeOut time.Duration) *http.Client {
	c := http.Client{}

	// transport settings
	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}

	c.Transport = t
	c.Timeout = timeOut

	return &c
}

func resolveReply(reply *Reply, err error) (*Reply, error) {
	if err != nil {
		return reply, err
	}

	// return error if reply.statuscode != 2xx
	if reply.StatusCode >= 300 {
		return reply, errors.Join(ErrHttpResponse, errors.New(fmt.Sprintf("status: %d, resp message: %s", reply.StatusCode, string(reply.Body))))
	}

	// check reply error code and related error message
	if err = checkReplyErrorCode(reply); err != nil {
		return nil, err
	}

	return reply, nil
}

func checkReplyErrorCode(reply *Reply) error {

	m := &Message{}
	if err := json.Unmarshal(reply.Body, m); err != nil {
		return err
	}
	switch m.ErrorCode {
	case SuccessCode:
		return nil
	case RateLimitedErrorCode:
		return warpError(ErrRateLimited, m.ErrorMessage)
	case InvalidContentErrorCode:
		return warpError(ErrInvalidMessage, m.ErrorMessage)
	case InvalidFileMd5ErrorCode:
		return warpError(ErrInvalidMd5, m.ErrorMessage)
	case InvalidImageSizedErrorCode:
		return warpError(ErrInvalidImageSize, m.ErrorMessage)
	case TitleMissedErrorCode:
		return warpError(ErrMissTitle, m.ErrorMessage)
	default:
		return errors.New(fmt.Sprintf("unknown response error type: %s", m.ErrorMessage))
	}
}
