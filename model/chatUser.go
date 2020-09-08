package model

// 用户私聊信息记录表
type ChatUser struct {
	Id          int
	SendUid     int
	RecvUid     int
	MessageType int
	Content     string
	Status      int
	CreatedAt   int64
}
