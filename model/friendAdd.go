package model

// 添加好友请求表
type FriendAdd struct {
	Id        int
	Uid       int
	FUid      int
	Channel   string
	Reason    string
	Status    int
	RequestAt string `gorm:"default:''"`
	PassAt    string `gorm:"default:''"`
}
