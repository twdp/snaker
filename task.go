package snaker

import (
	"container/list"
	"tianwei.pro/snaker/core"
	"tianwei.pro/snaker/model"
)

type ITaskService interface {

	// 根据任务模型、执行对象创建新的任务
	CreateTask(task *model.TaskModel, execution *core.Execution) (*list.List,  error)
}