package entity

type HistoryOrder struct {

	// 主键id
	Id int64

	// 流程定义id
	ProcessId int64

	// 流程实例状态(0: 结束, 1: 活动)
	OrderStatus int8


}