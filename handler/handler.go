package handler

import "tianwei.pro/snaker/core"

// 流程各模型操作处理接口
type IHandler interface {

	// 处理具体的操作
	Handle(execution *core.Execution) error
}