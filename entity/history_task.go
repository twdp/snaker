package entity

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"snaker/model"
	"time"
)

const (
	Normal = iota
	Fork
)
type HistoryTask struct {

	// 主键ID
	Id int64

	// 流程实例id
	OrderId int64

	// 任务名称
	TaskName string

	// 任务显示名称
	DisplayName string

	// 参与方式（0：普通任务；1：参与者fork任务[即：如果10个参与者，需要每个人都要完成，才继续流转]）
	PerformType int8

	// 任务类型
	TaskType int8

	// 任务状态（0：结束；1：活动）
	TaskStatus int8

	// 任务处理者ID
	OperatorId int64

	// 任务创建时间
	CreatedAt time.Time

	// 任务修改时间
	UpdatedAt time.Time

	// 任务完成时间
	FinishAt time.Time

	// 期望任务完成时间
	ExpireTime	time.Time

	// 任务关联的表单url
	ActionUrl string

	// 任务参与者列表
	ActorIds []int64

	// 父任务id
	ParentTaskId int64

	// 任务附属变量
	Variable string
	// todo:: task???
	TaskId int64
}

func NewHistoryTaskFromTask(task *Task) *HistoryTask {
	return &HistoryTask{
		TaskId: task.Id,
		OrderId: task.OrderId,
		DisplayName: task.DisplayName,
		TaskName: task.TaskName,
		TaskType: task.TaskType,
		ExpireTime: task.ExpireTime,
		ActionUrl: task.ActionUrl,
		ActorIds: task.ActorIds,
		ParentTaskId: task.ParentTaskId,
		Variable: task.Variable,
		PerformType: task.PerformType,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// 根据历史任务产生撤回的任务对象
func (h *HistoryTask) UndoTask() *Task {
	return &Task{
		OrderId:      h.OrderId,
		DisplayName:  h.DisplayName,
		TaskName:     h.TaskName,
		TaskType:     h.TaskType,
		ExpireTime:   h.ExpireTime,
		ActionUrl:    h.ActionUrl,
		ActorIds:     h.ActorIds,
		ParentTaskId: h.ParentTaskId,
		Variable:     h.Variable,
		PerformType:  h.PerformType,
	}
}

func (h *HistoryTask) isPerformAny() bool {
	return h.PerformType == model.Any
}

func (h *HistoryTask) GetVariableMap() map[string]interface{} {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(h.Variable), &m)
	if err != nil {
		logs.Error("history task unmarshal failed. variable: %s, err: %v", h.Variable, err)
	}
	return make(map[string]interface{})
}