package snaker

import "tianwei.pro/snaker/entity"

// 流程定义业务类
type IProcessService interface {

	// 检查流程定义对象
	Check(process *entity.Process, idOrName string) error

	// 根据主键ID获取流程定义对象
	GetProcessById(id int64) *entity.Process

	// 根据key获取流程定义对象
	GetProcessByKey(key string) *entity.Process

	// 根据key\version获取流程定义对象
	GetProcessByKeyAndVersion(key string) *entity.Process
}