package model

//记录用户的行为日志
type ActionLog struct {
	Id     uint
	Worker string //任务处理者
}
