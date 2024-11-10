package wecom_group_bot

import (
	"strings"

	"github.com/matrixorigin/wecom-group-bot/internal/wberr"
)

func NewFileMessage(mediaID string) Messager {
	return &FileMessage{
		Msgtype: MessageTypeFile,
		File: &File{
			MediaId: mediaID,
		},
	}
}

type FileMessage struct {
	Msgtype MessageType `json:"msgtype,omitempty"`
	File    *File       `json:"file,omitempty"`
}

func (f *FileMessage) SetType(messageType MessageType) {
	f.Msgtype = messageType
}

func (f *FileMessage) GetType() MessageType {
	return f.Msgtype
}

func (f *FileMessage) DeepCopy() Messager {
	return &FileMessage{
		Msgtype: f.Msgtype,
		File:    f.File.DeepCopy(),
	}
}

func (f *FileMessage) Validate() error {
	if f.Msgtype != MessageTypeFile {
		return wberr.OverrideError(wberr.ErrInvalidType,
			wberr.WithMessage("need is "+MessageTypeFile.String()+" but got "+f.Msgtype.String()),
		)
	}
	return f.File.Validate()
}

type File struct {
	MediaId string `json:"media_id,omitempty"`
}

func (f *File) Validate() error {
	return nil
}

func (f *File) DeepCopy() *File {
	return &File{
		MediaId: strings.Clone(f.MediaId),
	}
}
