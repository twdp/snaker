package model

import (
	"github.com/emirpasic/gods/lists"
	"github.com/emirpasic/gods/lists/arraylist"
)

type NodeModel struct {
	BaseModel

	Inputs lists.List

	Outputs lists.List

	// 前置局部拦截器实例集合
	PreInterceptors lists.List

	// 后置局部拦截器实例集合
	PostInterceptors lists.List

}

func NewNodeModel(name, displayName string) *NodeModel {
	return &NodeModel{
		BaseModel: BaseModel{
			Name: name,
			DisplayName: displayName,
		},
		Inputs: arraylist.New(),
		Outputs: arraylist.New(),
		PreInterceptors: arraylist.New(),
		PostInterceptors: arraylist.New(),
	}
}