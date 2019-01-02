package snaker

import "tianwei.pro/snaker/entity"

type IInstanceService interface {

	// 根据流程、操作人员、福流程实例ID创建流程实例
	CreateInstance(process *entity.Process, operator string, args map[string]interface{}) (*entity.Instance, error)

	// 根据流程、操作人员、父流程实例ID创建流程实例
	CreateInstanceUseParentInfo(process *entity.Process, operator string, args map[string]interface{}, parentId int64, parentNodeName string) (*entity.Instance, error)

	// 向指定实例id添加全局变量数据
	AddVariable(instanceId int64, args map[string]interface{}) error

	// 创建抄送实例
	CreateCCInstance(instanceId int64, creator string, actors ...string) error

	// 流程实例正常完成
	Complete(instanceId int64) error

	// 保存流程实例
	SaveInstance(instance *entity.Instance) error

	// 流程实例强制终止
	Terminate(instanceId int64) error

	// 流程实例强制终止
	TerminateByOperator(instanceId int64, operator string) error

	// 唤醒历史流程实例
	Resume(instanceId int64) error

	// 更新流程实例
	UpdateInstance(instance *entity.Instance) error

	// 更新抄送记录为已阅
	UpdateCCStatus(instanceId int64, actors ...string) error

	// 删除抄送记录
	DeleteCCInstance(instanceId int64, actor string) error

	// 级联删除所有数据
	CascadeRemove(id int64) error
}