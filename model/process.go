package model

import (
	"github.com/emirpasic/gods/lists"
	"github.com/emirpasic/gods/lists/arraylist"
)

type ProcessModel struct {

	BaseModel

	// 节点元素集合
	Nodes lists.List

	TaskModels lists.List

}

func NewProcess(name, displayName string) *ProcessModel {
	return &ProcessModel {
		BaseModel: BaseModel {
			Name: name,
			DisplayName: displayName,
		},
		Nodes: arraylist.New(),
		TaskModels: arraylist.New(),
	}
}