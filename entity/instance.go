package entity

// 流程工作单实体类（一般称为流程实例）
type Instance struct {

	Id int64

	ProcessId int64

	// 流程实例内容
	Content string

	// 流程实例附属变量
	Variable map[string]interface{}
}