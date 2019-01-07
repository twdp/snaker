package model

import "tianwei.pro/snaker/core"

type TransitionModel struct {
	BaseModel

	// 当前转移路径是否可用
	Enable bool

	// 变迁的目标节点应用
	Target *NodeModel

	// 变迁的源节点引用
	Source *NodeModel

	// 变迁的目标节点name名称
	To string

	//  变迁的条件表达式，用于decision
	Expr string

	// 转折点图形数据
	// G string
}
func (t *TransitionModel) Execute(execution *core.Execution) error {
	return nil
}
