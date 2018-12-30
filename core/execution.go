package core

import (
	"container/list"
	"tianwei.pro/snaker"
	"tianwei.pro/snaker/entity"
)

type Execution struct {

	Engine snaker.SnakerEngine

	// 流程定义对象
	Process *entity.Process

	// 流程实例对象
	Instance *entity.Instance

	// 父流程实例
	ParentInstance *entity.Instance


	// 父流程实例节点名称
	ParentNodeName string

	// 子流程实例节点名称
	ChildInstanceId int64

	// 执行参数
	Args map[string]interface{}

	// 操作人
	Operator string

	// 任务
	Task *entity.Task

	// 返回的任务列表
	Tasks *list.List

	// 是否已合并
	// 对于join节点的处理
	isMerged bool
}