package model

import "time"

//Tower的任务信息
type Task struct {
	Id          uint
	Name        string    //名称
	Content     string    //描述信息
	Url         string    //Tower地址
	ProjectName string    //项目名称
	Worker      string    //任务处理者
	DeadLine    time.Time //要求完成时间
	DoneFlag    bool      //是否已完成
}
