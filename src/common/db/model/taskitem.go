package model

import "time"

type TaskItem struct {
	Id          uint
	Name        string    //名称
	Url         string    //任务地址
	ProjectName string    //项目名称
	Worker      string    //任务处理者
	DeadLine    time.Time //要求完成时间
	DoneFlag    bool      //是否已完成
}
