package entity

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 历史任务实体类
type HistoryTask struct {

	Id int64

	// 流程实例ID
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
	Operator string

	// 任务完成时间
	FinishAt time.Time

	// 期望任务完成时间
	ExpireAt time.Time

	// 任务关联的表单url
	ActionUrl string

	// 任务参与者列表
	ActorIds string

	// 父任务id
	ParentTaskId int64

	// 任务负数变量
	Variable string

	CreatedAt orm.TimeField `orm:"auto_now_add"`

	UpdatedAt orm.TimeField `orm:"auto_now"`
}