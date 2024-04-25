package wecom_group_bot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/matrixorigin/wecom-group-bot/utils"
)

const (
	APIRateLimit         = 20
	APIRateLimitWaitTime = 1 * time.Minute
)

type Sender struct {
	id      atomic.Int32
	webhook string
	message chan Messager

	result        chan Result
	resultMap     map[int32]Result
	resultMapLock *sync.RWMutex

	ctx    context.Context
	cancel context.CancelFunc
}

func NewSender(webhook string) *Sender {
	ctx, cancel := context.WithCancel(context.Background())
	s := &Sender{
		id:            atomic.Int32{},
		webhook:       webhook,
		message:       make(chan Messager, APIRateLimit*2),
		result:        make(chan Result, APIRateLimit*2),
		resultMap:     make(map[int32]Result, APIRateLimit),
		resultMapLock: &sync.RWMutex{},
		ctx:           ctx,
		cancel:        cancel,
	}
	err := s.startServer()
	if err != nil {
		panic(err)
	}
	return s
}

func (s *Sender) startServer() error {
	if err := s.Validate(); err != nil {
		return err
	}
	// start send server
	go func(ctx context.Context, webhook string, messageChan chan Messager, resultChan chan Result) {
		url := utils.URL{
			RawURL: webhook,
		}
		queue := make(chan Messager, 1)
		backoff := make(chan Messager, APIRateLimit*2)

		go func(ctx context.Context) {
			needExit := false
			for {
				select {
				case <-ctx.Done():
					needExit = true
				case message := <-backoff:
					fmt.Printf("message.GetID(): %v\n", message.GetID())
					queue <- message
				case message := <-messageChan:
					backoff <- message
				default:
					if len(backoff) == 0 && len(messageChan) == 0 && needExit {
						return
					}
				}
			}
		}(ctx)

		needExit := false
		for {
			select {
			case <-ctx.Done():
				needExit = true
			case message := <-queue:
				result := sendMessage(url, message, 5, 5*time.Second)
				// if err is rate limited error
				if result.Status == BackOffStatus {
					backoff <- message
				}
				resultChan <- result
			default:
				if len(queue) == 0 && needExit {
					return
				}
			}
		}
	}(s.ctx, s.webhook, s.message, s.result)

	return nil
}

func (s *Sender) StopServer() error {
	s.cancel()
	return nil
}

func (s *Sender) Send(message Messager) (id int32, err error) {
	if err = message.Validate(); err != nil {
		return -1, err
	}
	id = s.id.Add(1) - 1
	message.SetID(id)
	s.message <- message
	return id, nil
}

func (s *Sender) WaitResult(id int32) *Result {
	for i := 0; i < 300; i++ {
		if result := s.GetResult(id); result != nil {
			return result
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (s *Sender) GetResult(id int32) *Result {
	s.resultMapLock.Lock()
	for {
		result := s.readResultWithSelect()
		if result != nil {
			s.resultMap[result.ID] = *result
			continue
		}
		break
	}
	s.resultMapLock.Unlock()

	s.resultMapLock.RLock()
	if value, ok := s.resultMap[id]; ok {
		return &value
	}
	s.resultMapLock.RUnlock()
	return nil
}

func (s *Sender) readResultWithSelect() *Result {
	timer := time.NewTimer(500 * time.Millisecond)
	select {
	case <-timer.C:
		return nil
	case result := <-s.result:
		fmt.Printf("result.Status: %v\n", result.Status)
		return &result
	default:
		return nil
	}
	return nil
}

func (s *Sender) Validate() error {
	if s.webhook == "" {
		return ErrEmptyWebhook
	}
	if s.message == nil {
		return ErrNilChannel
	}
	return nil
}

func sendMessage(url utils.URL, message Messager, retryTimes int, retryDuration time.Duration) Result {
	result := Result{
		ID:   message.GetID(),
		Err:  nil,
		Time: time.Now(),
	}
	payload, err := json.Marshal(message)
	if err != nil {
		result.Err = err
		return result
	}
	_, err = utils.PostWithRetry(url, payload, retryTimes, retryDuration)
	fmt.Printf("err: %v\n", err)
	result.Err = err
	switch {
	case errors.Is(err, utils.ErrRateLimited):
		result.Status = BackOffStatus
	case err == nil:
		result.Status = SuccessStatus
	default:
		result.Status = FailedStatus
	}
	return result
}

type Messager interface {
	SetType(messageType string)
	GetType() string
	DeepCopy() Messager
	Validate() error

	SetID(int32)
	GetID() int32
}
