package snaker

import "tianwei.pro/snaker/entity"

// 流程定义业务类
type IProcessService interface {

	// 检查流程定义对象
	Check(process *entity.Process, idOrName string) error

	// 保存流程定义
	SaveProcess(process *entity.Process) error

	// 更新流程定义的类型
	// tt 类型
	UpdateType(id int64, tt int8) error

	// 根据主键ID获取流程定义对象
	GetProcessById(id int64) *entity.Process

	// 根据key获取流程定义对象
	GetProcessByKey(key string) *entity.Process

	// 根据key\version获取流程定义对象
	GetProcessByKeyAndVersion(key string) *entity.Process

	// 部署流程定义
	// return:  id -> 主键
	Deploy(input string) (int64, error)

	// 重新部署流程
	ReDeploy(id int64, input string) error

	// 卸载流程，更新状态
	UnDeploy(id int64) error

	// 删除关联关系
	CascadeRemove(id int64) error
}
