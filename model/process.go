package model

import (
	"container/list"
	"errors"
	"fmt"
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


func (p *ProcessModel) GetWorkModels() list.List {
	r := list.New()
	for _, e := range p.Nodes.Values() {
		if v, ok := e.(*WorkModel); ok {
			r.PushBack(v)
		}
	}
	return *r
}

func (p *ProcessModel) GetStart() (*StartModel, error) {
	for _, e := range p.Nodes.Values() {
		if v, ok := e.(*StartModel); ok {
			return v, nil
		}
	}
	return nil, errors.New("没有start节点")
}

func (p *ProcessModel) GetNode(nodeName string) (*NodeModel, error) {
	for _, e := range p.Nodes.Values() {
		if v, ok := e.(*NodeModel); ok {
			if v.Name == nodeName {
				return v, nil
			}
		}
	}

	return nil, errors.New(fmt.Sprintf("没有[%s]节点", nodeName))
}