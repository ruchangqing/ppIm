package model

type FriendList struct {
	Id        int
	Uid       int
	FUid      int
	Channel   string
	Reason    string
	Role      int
	CreatedAt string `gorm:"default:''"`
}
