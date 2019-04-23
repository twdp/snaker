package entity

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	PerformtypeAll = iota

)
// 提醒时间
// 多长时间自动过期
// 延迟执行时间

type Task struct {
	Id int64

	Version int8

	OrderId int64

	TaskName string

	DisplayName string

	PerformType int8

	TaskType int8

	Operator string

	FinishAt time.Time

	ExpireAt time.Time

	RemindAt time.Time

	ActionUrl string

	ActorIds string

	ParentTaskId int64

	Variable string

	CreatedAt orm.TimeField `orm:"auto_now_add"`

	UpdatedAt orm.TimeField `orm:"auto_now"`
}