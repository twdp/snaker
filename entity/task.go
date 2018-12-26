package entity

import (
	"snaker/model"
	"time"
)

// 任务实体类
type Task struct {

	// 主键id
	Id int64

	// 版本
	Version int8

	// 流程实例Id
	OrderId int64

	// 任务名称
	TaskName string

	// 任务显示名称
	DisplayName string

	// 参与方式(0: 普通任务 1: 参与者会签任务)
	PerformType int8

	// 任务类型(0: 主办任务 1: 协办任务)
	TaskType int8

	// 任务处理者Id
	OperatorId string

	// 任务创建时间
	CreatedAt time.Time

	// 任务完成时间
	FinishTime time.Time

	// 期望任务完成时间
	ExpireTime time.Time

	// 期望的完成时间date类型
	ExpireDate time.Time

	// 提醒时间
	RemindDate time.Time

	// 任务关联的表单url
	ActionUrl string

	// 任务参与者列表
	ActorIds []int64

	// 父任务id
	ParentTaskId int64

	// 任务附属变量
	Variable string

	// 保持模型对象
	TaskModel *model.TaskModel
}