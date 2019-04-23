package entity

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 历史流程实例实体类
type HistoryOrder struct {

	Id int64

	// 流程定义id
	ProcessId int64

	// 流程实例状态（0：结束；1：活动）
	OrderStatus int8

	// 流程实例创建者ID
	Creator string

	// 流程实例结束时间
	EndAt time.Time

	// 流程实例为子流程时，该字段标识父流程实例ID
	ParentId int64

	// 流程实例期望完成时间
	ExpiredAt time.Time

	// 流程实例优先级
	Priority int64

	// 流程实例编号
	OrderId int64

	// 流程实例附属变量
	Variable string

	CreatedAt orm.TimeField `orm:"auto_now_add"`

	UpdatedAt orm.TimeField `orm:"auto_now"`
}

func int() {
	orm.RegisterModelWithPrefix("snaker_", new(HistoryOrder))
}