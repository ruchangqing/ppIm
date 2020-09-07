package model

// 已添加的好友列表表
type FriendList struct {
	Id        int
	Uid       int
	FUid      int
	Channel   string
	Reason    string
	Role      int
	CreatedAt string `gorm:"default:''"`
}
