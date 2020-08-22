package model

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
