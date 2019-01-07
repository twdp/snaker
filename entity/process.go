package entity

import (
	"tianwei.pro/snaker"
)

const (
	ProcessInit = iota

)
type Process struct {

	Base

	// 版本
	Version int

	// 流程定义名称
	Name string

	// 流程定义显示名称
	DisplayName string

	// 流程定义类型（预留字段）
	Type int8

	// 预留字段
	// 当前流程的实例url（一般为流程第一步的url）
	// 该字段可以直接打开流程申请的表单
	InstanceUrl string

	// 当前状态
	Status int8

	// 流程定义模型
	//Model *model.ProcessModel

	CreatedBy *snaker.Logger

	// 流程定义xml
	Content string
}

//func (p *Process) SetModel(processModel *model.ProcessModel) {
//	//p.Model = processModel
//	p.Name = processModel.Name
//	p.DisplayName = processModel.DisplayName
//	p.InstanceUrl = processModel.InstanceUrl
//}