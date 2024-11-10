package wecom_group_bot

import (
	"github.com/matrixorigin/wecom-group-bot/internal/wberr"
)

func NewVoiceMessage(mediaID string) Messager {
	return &VoiceMessage{
		Msgtype: MessageTypeVoice,
		Voice: &File{
			MediaId: mediaID,
		},
	}
}

type VoiceMessage struct {
	Msgtype MessageType `json:"msgtype,omitempty"`
	Voice   *File       `json:"voice,omitempty"`
}

func (v *VoiceMessage) SetType(messageType MessageType) {
	v.Msgtype = messageType
}

func (v *VoiceMessage) GetType() MessageType {
	return v.Msgtype
}

func (v *VoiceMessage) DeepCopy() Messager {
	return &VoiceMessage{
		Msgtype: v.Msgtype,
		Voice:   v.Voice.DeepCopy(),
	}
}

func (v *VoiceMessage) Validate() error {
	if v.Msgtype != MessageTypeVoice {
		return wberr.OverrideError(wberr.ErrInvalidType,
			wberr.WithMessage("need is "+MessageTypeVoice.String()+" but got "+v.Msgtype.String()),
		)
	}
	return v.Voice.Validate()
}
