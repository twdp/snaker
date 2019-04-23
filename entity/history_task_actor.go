package entity

import "github.com/astaxie/beego/orm"

// 历史任务参与者实体类
type HistoryTaskActor struct {
	Id int64

	// 关联的任务ID
	TaskId int64

	// 关联的参与者ID（参与者可以为用户、部门、角色）
	ActorId string

	CreatedAt orm.TimeField `orm:"auto_now_add"`

	UpdatedAt orm.TimeField `orm:"auto_now"`
}
