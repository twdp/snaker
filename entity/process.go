package entity

// 流程定义实体类
type Process struct {

	// 主键
	Id int64

	// 版本
	Version int

	// 流程定义名称，根据此字段启动流程
	Name string

	// 页面上展示的名称
	DisplayName string

	// 当前状态
	Status int8

	// 流程定义内容
	Content string
}

