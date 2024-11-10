package wecom_group_bot

const (
	TextCardType = "text_notice"
	NewsCardType = "news_notice"
)

type CardMessage struct {
	Msgtype      string `json:"msgtype,omitempty"`
	TemplateCard Carder `json:"template_card,omitempty"`
}

type Carder interface {
	SetCardType(cardType string)
	GetCardType() string
}

type NewsCard struct {
	commonCard
	CardImage           *CardImage          `json:"card_image,omitempty"`
	ImageTextArea       *ImageTextArea      `json:"image_text_area,omitempty"`
	VerticalContentList VerticalContentList `json:"vertical_content_list,omitempty"`
}

func (n *NewsCard) SetCardType(cardType string) {
	n.CardType = cardType
}

func (n *NewsCard) GetCardType() string {
	return n.CardType
}

type TextCard struct {
	commonCard
	EmphasisContent *EmphasisContent `json:"emphasis_content,omitempty"`
	SubTitleText    string           `json:"sub_title_text,omitempty"`
}

func (t *TextCard) SetCardType(cardType string) {
	t.CardType = cardType
}

func (t *TextCard) GetCardType() string {
	return t.CardType
}

type Source struct {
	IconURL   string `json:"icon_url,omitempty"`
	Desc      string `json:"desc,omitempty"`
	DescColor int    `json:"desc_color,omitempty"`
}
type MainTitle struct {
	Title string `json:"title,omitempty"`
	Desc  string `json:"desc,omitempty"`
}
type EmphasisContent struct {
	Title string `json:"title,omitempty"`
	Desc  string `json:"desc,omitempty"`
}
type QuoteArea struct {
	Type      int    `json:"type,omitempty"`
	URL       string `json:"url,omitempty"`
	AppID     string `json:"appid,omitempty"`
	PagePath  string `json:"pagepath,omitempty"`
	Title     string `json:"title,omitempty"`
	QuoteText string `json:"quote_text,omitempty"`
}
type HorizontalContent struct {
	KeyName string `json:"keyname,omitempty"`
	Value   string `json:"value,omitempty"`
	Type    int    `json:"type,omitempty,omitempty"`
	URL     string `json:"url,omitempty,omitempty"`
	MediaID string `json:"media_id,omitempty,omitempty"`
}
type HorizontalContentList []HorizontalContent

type Jump struct {
	Type     int    `json:"type,omitempty"`
	URL      string `json:"url,omitempty,omitempty"`
	Title    string `json:"title,omitempty"`
	AppID    string `json:"appid,omitempty,omitempty"`
	PagePath string `json:"pagepath,omitempty,omitempty"`
}
type JumpList []Jump

type CardAction struct {
	Type     int    `json:"type,omitempty"`
	URL      string `json:"url,omitempty"`
	AppID    string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`
}

type commonCard struct {
	CardType              string                 `json:"card_type,omitempty"`
	Source                *Source                `json:"source,omitempty"`
	MainTitle             *MainTitle             `json:"main_title,omitempty"`
	QuoteArea             *QuoteArea             `json:"quote_area,omitempty"`
	HorizontalContentList *HorizontalContentList `json:"horizontal_content_list,omitempty"`
	JumpList              *JumpList              `json:"jump_list,omitempty"`
	CardAction            *CardAction            `json:"card_action,omitempty"`
}

type CardImage struct {
	URL         string  `json:"url,omitempty"`
	AspectRatio float64 `json:"aspect_ratio,omitempty"`
}
type ImageTextArea struct {
	Type     int    `json:"type,omitempty"`
	URL      string `json:"url,omitempty"`
	Title    string `json:"title,omitempty"`
	Desc     string `json:"desc,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
}

type VerticalContent struct {
	Title string `json:"title,omitempty"`
	Desc  string `json:"desc,omitempty"`
}
type VerticalContentList []VerticalContent
