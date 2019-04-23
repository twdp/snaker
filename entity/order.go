package entity

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 流程工作单实体类（一般称为流程实例）
type Order struct {
	Id int64

	// 版本
	Version int8

	// 流程定义id
	ProcessId int64

	// 流程实例创建者ID
	Creator string

	// 流程实例为子流程时，该字段标识福流程实例ID
	ParentId int64

	// 流程实例为子流程时，该字段标识福流程哪个节点模型启动的子流程
	ParentNodeName string

	// 流程实例期望完成时间
	ExpireTime time.Time

	// 流程实例优先级
	Priority int64

	// 流程实例附属变量
	Variable string

	CreatedAt orm.TimeField `orm:"auto_now_add"`

	UpdatedAt orm.TimeField `orm:"auto_now"`
}
