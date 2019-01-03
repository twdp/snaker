package handler

import "tianwei.pro/snaker/core"

// 流程各模型操作处理接口
type IHandler interface {

	// 子类需要实现的方法，来处理具体的操作
	Handler(execution *core.Execution) error
}