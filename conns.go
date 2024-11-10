package wecom_group_bot

const (
	NoticePrefix      = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"
	UploadMediaPrefix = "https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media"
)

type MessageType string

const (
	MessageTypeText     MessageType = "text"
	MessageTypeMarkdown MessageType = "markdown"
	MessageTypeImage    MessageType = "image"
	MessageTypeNews     MessageType = "news"
	MessageTypeFile     MessageType = "file"
	MessageTypeVoice    MessageType = "voice"
	MessageTypeCard     MessageType = "template_card"

	MessageTypeMedia MessageType = "media"
)

func (nt MessageType) String() string {
	return string(nt)
}
