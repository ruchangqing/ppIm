package model

// 消息记录
type ChatMessage struct {
	Id        int
	FromId    int
	ToId     int
	Ope       int 	 //消息通道：0 好友消息，1 群消息
	Type      int 	 //消息类型：0 文本消息，1 图片，2 语音，3 视频，4 地理位置信息，6 文件，10 提示消息
	Body      string //消息内容
	Status    int    //消息状态：0 未读，1 已读，-1 已撤回
	CreatedAt int64
}
