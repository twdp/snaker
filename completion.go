package snaker

import "snaker/entity"

// 任务、实例完成时触发动作的接口
type Completion interface {

	// 任务完成触发执行
	completeTask(task entity.HistoryTask) error

	// 实例完成触发执行
	completeOrder(order entity.HistoryOrder) error
}