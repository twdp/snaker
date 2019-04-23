package entity

import "github.com/astaxie/beego/orm"

type TaskActor struct {

	Id int64

	TaskId int64

	ActorId string

	CreatedAt orm.TimeField `orm:"auto_now_add"`

	UpdatedAt orm.TimeField `orm:"auto_now"`
}