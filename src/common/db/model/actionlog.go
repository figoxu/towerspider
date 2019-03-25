package model

import "github.com/jinzhu/gorm"

//记录用户的行为日志
type ActionLog struct {
	gorm.Model
	EventId     string
	Actor       string //任务处理者
	Action      string
	Text        string
	SysName     string
	ResourceUrl string
}
