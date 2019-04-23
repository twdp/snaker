package entity

import "github.com/astaxie/beego/orm"

// 抄送实例实体
type CCOrder struct {

	Id        int64

	ActorId   string

	Creator   string

	FinishAt  orm.TimeField

	Status    int8

	CreatedAt orm.TimeField `orm:"auto_now_add"`

	UpdatedAt orm.TimeField `orm:"auto_now"`
}

func init() {
	orm.RegisterModelWithPrefix("snaker_", new(CCOrder))
}
