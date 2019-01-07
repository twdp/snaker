package entity

import (
	"tianwei.pro/snaker"
	"time"
)

const (
	Major = iota
	Aidant
	Record
)

const (
	PerformtypeAny = iota
	PerformtypeAll
)

// 任务实体类
type Task struct {

	Base

	Version int

	// 流程实例id
	InstanceId int64

	// 任务名称
	TaskName string

	// 任务显示名称
	DisplayName string

	// 参与方式（0：普通任务；1：参与者会签任务）
	PerformType int8

	// 任务类型（0：主办任务；1：协办任务）
	TaskType int8

	// 任务处理者信息
	Operator *snaker.Logger

	// 任务完成时间
	FinishAt time.Time

	// 期望任务完成时间
	ExpireAt time.Time

	// 提醒时间
	RemindAt time.Time

	// 任务关联的表单url
	ActionUrl string

	// 参与者列表
	Actors []*snaker.Logger

	// 父任务id
	ParentTaskId int64

	// 任务附属变量
	Variable map[string]interface{}

	// 模型对象
	//model *model.TaskModel
}

func (t *Task) IsMajor() bool {
	return t.TaskType == Major
}