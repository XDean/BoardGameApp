package model

const (
	TEXT = "text"
)

type Message struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}
