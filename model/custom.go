package model

import "tianwei.pro/snaker/core"

// 用户自定义处理
// snaker对外提供一个di容器
// 实现接口并注入到容器中
// snaker 处理时调用
type CustomModel struct {
	WorkModel

	// 实例名称
	Clazz string

	// 传入参数
	//Args string

}

// todo::
func (c *CustomModel) exec(execution *core.Execution) error {
	// 从di容器中查找指定的实例
	return nil
}