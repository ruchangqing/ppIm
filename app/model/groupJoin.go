package model

// 群组添加临时表
type GroupJoin struct {
	Id      int
	GroupId int
	UserId  int
	JoinAt  int64
}
