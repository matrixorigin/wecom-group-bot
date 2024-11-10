package wecom_group_bot

import (
	"strings"

	"github.com/matrixorigin/wecom-group-bot/internal/wberr"
)

func NewImageMessage(image *Image) Messager {
	return &ImageMessage{
		Msgtype: MessageTypeImage,
		Image:   image.DeepCopy(),
	}
}

type ImageMessage struct {
	Msgtype MessageType `json:"msgtype,omitempty"`
	Image   *Image      `json:"image,omitempty"`
}

func (i *ImageMessage) SetType(messageType MessageType) {
	i.Msgtype = messageType
}

func (i *ImageMessage) GetType() MessageType {
	return i.Msgtype
}

func (i *ImageMessage) DeepCopy() Messager {
	return &ImageMessage{
		Msgtype: i.Msgtype,
		Image:   i.Image.DeepCopy(),
	}
}

func (i *ImageMessage) Validate() error {
	if i.Msgtype != MessageTypeImage {
		return wberr.OverrideError(wberr.ErrInvalidType,
			wberr.WithMessage("need is "+MessageTypeImage.String()+" but got "+i.Msgtype.String()),
		)
	}
	return i.Image.Validate()
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
