package entity

import "tianwei.pro/snaker"

// 流程工作单实体类（一般称为流程实例）
type Instance struct {

	Base

	// 版本
	Version int

	// 流程定义id
	ProcessId int64

	// 创建人
	CreatedBy *snaker.Logger

	// 程实例为子流程时，该字段标识父流程实例ID
	ParentId int64

	// 流程实例为子流程时，该字段标识父流程哪个节点模型启动的子流程
	ParentNodeName string

	// 流程实例附属变量
	Variable map[string]interface{}

	// 流程优先级
	//Priority int

	// 流程实例期望完成时间
	//ExpireTime time.Time
}