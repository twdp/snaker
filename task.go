package snaker

import (
	"container/list"
	"tianwei.pro/snaker/core"
	"tianwei.pro/snaker/entity"
	"tianwei.pro/snaker/model"
)

/**
 * 任务业务类，包括以下服务：
 * 1、创建任务
 * 2、添加、删除参与者
 * 3、完成任务
 * 4、撤回任务
 * 5、回退任务
 */
type ITaskService interface {

	// 根据任务模型、执行对象创建新的任务
	CreateTask(task *model.TaskModel, execution *core.Execution) (*list.List,  error)

	// 完成指定的任务，删除活动任务记录，创建历史任务
	CompleteById(id int64) (*entity.Task, error)

	// 完成指定的任务，删除活动任务记录，创建历史任务
	CompleteByIdAndOperator(id int64, operator string) (*entity.Task, error)

	// 根据任务主键ID，操作人ID完成任务
	CompleteByIdAndOperatorAndArgs(id int64, operator string, args map[string]interface{}) (*entity.Task, error)

	// 更新任务对象
	UpdateTask(task *entity.Task) error

	// 根据执行对象、自定义节点模型创建历史任务记录
	History(execution *core.Execution, model CustomModel) (*entity.HistoryTask, error)

	// 根据历史任务主键id，操作人唤醒历史任务
	// 该方法会导致流程状态不可控，请慎用
	Resume(taskId int64, operator string) (*entity.Task, error)

	// 向指定的任务id添加参与者
	AddTaskActor(taskId int64, actors ...string) error

	// 向指定的任务id添加参与者
	AddTaskActorByPerformType(task int64, performType int8, actors ...string) error

	// 对指定的任务id删除参与者
	RemoveTaskActor(taskId int64, actor ...string)

	// 根据任务主键id、操作人撤回任务
	WithdrawTask(taskId int64, operator string) (*entity.Task, error)

	// 根据当前任务对象驳回至上一步处理
	RejectTask(model *model.ProcessModel, currentTask *entity.Task) (*entity.Task, error)

	// 根据taskId、operator，判断操作人operator是否允许执行任务
	IsAllowed(task *entity.Task, operator string)

	// 根据已有任务id、任务类型、参与者创建新的任务
	CreateNewTask(taskId int64, taskType int8, actors ...string) (list.List, error)

	// 根据任务id获取任务模型
	GetTaskModel(taskId int64) (*model.TaskModel, error)
}