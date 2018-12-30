package model

type TransitionModel struct {

	// 当前转移路径是否可用
	Enable bool

	// 变迁的目标节点应用
	Target *NodeModel

	// 变迁的源节点引用
	Source *NodeModel
}