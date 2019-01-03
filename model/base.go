package model

import (
	"tianwei.pro/snaker/core"
	"tianwei.pro/snaker/handler"
)

type BaseModel struct {

	// 元素名称
	Name string

	// 显示名称
	Display string
}

// 将执行对象execution交给具体的处理器处理
func (b *BaseModel) Fire(handler handler.IHandler, execution *core.Execution) {
	handler.Handler(execution)
}