package model

import "tianwei.pro/snaker/core"

type CustomModel struct {
	WorkModel

	// 实例名称
	Clazz string

	// 传入参数
	//Args string

}

func (c *CustomModel) exec(execution *core.Execution) error {
	// 从di容器中查找指定的实例
	return nil
}