package model

import (
	"container/list"
	"errors"
	"fmt"
)

type ProcessModel struct {

	// 节点元素集合
	Nodes list.List

	TaskModes list.List

	InstanceUrl string


}

func (p *ProcessModel) GetWorkModels() list.List {
	r := list.New()
	for e := p.Nodes.Front(); e != nil; e = e.Next() {
		if v, ok := e.Value.(*WorkModel); ok {
			r.PushBack(v)
		}
	}
	return *r
}

func (p *ProcessModel) GetStart() (*StartModel, error) {
	for e := p.Nodes.Front(); e != nil; e = e.Next() {
		if v, ok := e.Value.(*StartModel); ok {
			return v, nil
		}
	}
	return nil, errors.New("没有start节点")
}

func (p *ProcessModel) GetNode(nodeName string) (*NodeModel, error) {
	for e := p.Nodes.Front(); e != nil; e = e.Next() {
		if v, ok := e.Value.(*NodeModel); ok {
			if v.Name == nodeName {
				return v, nil
			}
		}
	}
	return nil, errors.New(fmt.Sprintf("没有[%s]节点", nodeName))
}