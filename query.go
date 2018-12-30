package snaker

import "tianwei.pro/snaker/entity"

// 流程相关查询服务
type IQueryService interface {

	// 根据流程实例id获取流程实例对象
	GetInstance(instanceId int64) (*entity.Instance, error)
}