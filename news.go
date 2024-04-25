package wecom_group_bot

import (
	"errors"
	"strings"
	"time"
)

type NewsMessage struct {
	id      int32
	Msgtype string `json:"msgtype,omitempty"`
	News    *News  `json:"news,omitempty"`
}

func (n *NewsMessage) SetType(messageType string) {
	n.Msgtype = messageType
}

func (n *NewsMessage) GetType() string {
	return n.Msgtype
}

func (n *NewsMessage) DeepCopy() Messager {
	return &NewsMessage{
		id:      n.id,
		Msgtype: strings.Clone(n.Msgtype),
		News:    n.News.DeepCopy(),
	}
}

func (n *NewsMessage) Validate() error {
	if n.Msgtype != NewsType {
		return errors.Join(ErrInvalidType, errors.New("need is "+NewsType+" but got "+n.Msgtype))
	}
	return n.News.Validate()
}

func (n *NewsMessage) SetID(id int32) {
	n.id = id
}

func (n *NewsMessage) GetID() int32 {
	return n.id
}

type News struct {
	Articles Articles `json:"articles,omitempty"`
}

func (n *News) DeepCopy() *News {
	return &News{
		Articles: n.Articles.DeepCopy(),
	}
}

func (n *News) Validate() error {
	return n.Articles.Validate()
}

func (n *NewsMessage) SetBackOff(t time.Time) {
	//TODO implement me
	panic("implement me")
}

func (n *NewsMessage) GetBackOff() *time.Time {
	//TODO implement me
	panic("implement me")
}

type Article struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	PicURL      string `json:"picurl,omitempty"`
}

type Articles []Article

func (as Articles) Validate() error {
	return nil
}

func (as Articles) DeepCopy() Articles {
	dst := make(Articles, 0, len(as))
	for _, article := range as {
		dst = append(dst, Article{
			Title:       strings.Clone(article.Title),
			Description: strings.Clone(article.Description),
			URL:         strings.Clone(article.URL),
			PicURL:      strings.Clone(article.PicURL),
		})
	}
	return dst
}
