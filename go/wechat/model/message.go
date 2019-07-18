package model

const (
	TEXT = "text"
)

type Message struct {
	XMLName struct{} `xml:"xml""`

	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}
