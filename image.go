package wecom_group_bot

import (
	"errors"
	"strings"
	"time"
)

type ImageMessage struct {
	id      int32
	Msgtype string `json:"msgtype,omitempty"`
	Image   *Image `json:"image,omitempty"`
}

func (i *ImageMessage) SetType(messageType string) {
	i.Msgtype = messageType
}

func (i *ImageMessage) GetType() string {
	return i.Msgtype
}

func (i *ImageMessage) DeepCopy() Messager {
	return &ImageMessage{
		id:      i.id,
		Msgtype: strings.Clone(i.Msgtype),
		Image:   i.Image.DeepCopy(),
	}
}

func (i *ImageMessage) Validate() error {
	if i.Msgtype != ImageType {
		return errors.Join(ErrInvalidType, errors.New("need is "+ImageType+" but got "+i.Msgtype))
	}
	return i.Image.Validate()
}

func (i *ImageMessage) SetID(id int32) {
	i.id = id
}

func (i *ImageMessage) GetID() int32 {
	return i.id
}

func (i *ImageMessage) SetBackOff(t time.Time) {
	//TODO implement me
	panic("implement me")
}

func (i *ImageMessage) GetBackOff() *time.Time {
	//TODO implement me
	panic("implement me")
}

type Image struct {
	Base64 string `json:"base64,omitempty"`
	Md5    string `json:"md5,omitempty"`
}

func (i *Image) DeepCopy() *Image {
	return &Image{
		Base64: strings.Clone(i.Base64),
		Md5:    strings.Clone(i.Md5),
	}
}

func (i *Image) Validate() error {
	return nil
}
