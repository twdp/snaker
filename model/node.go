package model

import (
	"github.com/emirpasic/gods/lists"
	"github.com/emirpasic/gods/lists/arraylist"
	"tianwei.pro/snaker/core"
	"tianwei.pro/snaker/entity"
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

//  对执行逻辑增加前置、后置拦截处理
func (n *NodeModel) Execute(execution *core.Execution) error {
	//if err := n.intercept(n.PreInterceptors, execution); err != nil {
	//	return err
	//} else if err = n.exec(execution); err != nil {
	//	return err
	//} else if err = n.intercept(n.PostInterceptors, execution); err != nil {
	//	return err
	//}
	return nil
}

// 具体节点模型需要完成的执行逻辑
func (n *NodeModel) exec(execution *core.Execution) error {
	panic("子模型需要实现exec方法")
}

// 运行变迁继续执行
func (n *NodeModel) runOutTransition(execution *core.Execution) error {
	for _, e := range n.Outputs.Values() {
		tm := e.(*TransitionModel)
		tm.Enable = true
		if err := tm.Execute(execution); err != nil {
			return err
		}
	}

	return nil
}

//// 拦截方法
//func (n *NodeModel) intercept(interceptors lists.List, execution *core.Execution) error {
//	for _, e := range n.Outputs.Values() {
//		interceptor := e.(*snaker.Interceptor)
//		if err := (*interceptor).Intercept(execution); err != nil {
//			return err
//		}
//	}
//	return nil
//}

/**
 * 根据父节点模型、当前节点模型判断是否可退回。可退回条件：
 * 1、满足中间无fork、join、subprocess模型
 * 2、满足父节点模型如果为任务模型时，参与类型为any
 */
func (n *NodeModel) CanRejected(current *NodeModel, parent *NodeModel) bool {

	switch t := (interface{})(parent).(type) {
	case *TaskModel:
		return t.PerformType == entity.PerformtypeAll
	}
	result := false

	for _, e := range n.Outputs.Values() {
		tm := e.(*TransitionModel)
		source := tm.Source
		if source == parent {
			return true
		}
		switch _ := (interface{})(source).(type) {
		//case *ForkModel:
		//	continue
		//case *JoinModel:
		//	continue
		//case *SubProcessModel:
		//	continue
		case *StartModel:
			continue
		}
		result = result || n.CanRejected(source, parent);
	}
	return result
}