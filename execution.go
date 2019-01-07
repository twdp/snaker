package snaker

import (
	"github.com/emirpasic/gods/lists"
	"tianwei.pro/snaker/entity"
)

// 流程执行过程中所传递的执行对象，其中包含流程定义、流程模型、流程实例对象、执行参数、返回的任务列表
type Execution struct {

	// @Deprecated
	Process *entity.Process

	// 当前流程执行的流程模型
	ProcessModel *ProcessModel

	Engine Engine

	// 执行时传递的参数
	Args map[string]interface{}

	Operator string

	ParentInstance *entity.Instance

	Instance *entity.Instance

	ParentNodeName string

	Tasks lists.List

	//Task *entity.Task
}