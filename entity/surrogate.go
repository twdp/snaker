package entity

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 委托代理实体类
type Surrogate struct {
	//
	Id int64

	// 流程name
	ProcessName string

	// 授权人
	Operator string

	// 代理人
	Surrogate string

	// 操作时间
	ODate string

	// 开始时间
	StartAt time.Time

	// 结束时间
	EndAt time.Time

	// 状态
	Status int8

	CreatedAt orm.TimeField `orm:"auto_now_add"`

	UpdatedAt orm.TimeField `orm:"auto_now"`
}
