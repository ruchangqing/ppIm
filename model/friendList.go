package model

type FriendList struct {
	Id        int
	Uid       int
	FUid      int
	Channel   int
	Reason    string
	CreatedAt string `gorm:"default:''"`
}
