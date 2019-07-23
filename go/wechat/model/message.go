package model

const (
	TEXT  = "text"
	IMAGE = "image"
)

type Message struct {
	XMLName struct{} `xml:"xml""`

	MsgId        int
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string

	// text
	Content string

	// image
	PicUrl string
}

func NewText(s string) Message {
	return Message{
		MsgType: TEXT,
		Content: s,
	}
}
