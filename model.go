package snaker

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/emirpasic/gods/lists"
	"github.com/emirpasic/gods/lists/arraylist"
	"tianwei.pro/snaker/entity"
)

type BaseModel struct {

	// 元素名称
	Name string

	// 显示名称
	DisplayName string
}

// 将执行对象execution交给具体的处理器处理
func (b *BaseModel) fire(handler IHandler, execution *Execution) error {
	return handler.Handle(execution)
}


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
func (n *NodeModel) Execute(execution *Execution) error {
	if err := n.intercept(n.PreInterceptors, execution); err != nil {
		return err
	} else if err = n.exec(execution); err != nil {
		return err
	} else if err = n.intercept(n.PostInterceptors, execution); err != nil {
		return err
	}
	return nil
}

// 具体节点模型需要完成的执行逻辑
func (n *NodeModel) exec(execution *Execution) error {
	panic("子模型需要实现exec方法")
}

// 运行变迁继续执行
func (n *NodeModel) runOutTransition(execution *Execution) error {
	for _, v := range n.Outputs.Values() {
		tm := v.(*TransitionModel)
		tm.Enable = true
		if err := tm.Execute(execution); err != nil {
			return err
		}
	}
	return nil
}

// 拦截方法
func (n *NodeModel) intercept(interceptors lists.List, execution *Execution) error {
	for _, v := range interceptors.Values() {
		interceptor := v.(*Interceptor)
		if err := (*interceptor).Intercept(execution); err != nil {
			return err
		}
	}
	return nil
}

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
		switch s := (interface{})(source).(type) {
		//case *ForkModel:
		//	continue
		//case *JoinModel:
		//	continue
		//case *SubProcessModel:
		//	continue
		case *StartModel:
			logs.Debug(s)
			continue
		}
		result = result || n.CanRejected(source, parent);
	}
	return result
}

type StartModel struct {
	NodeModel
}

func (s *StartModel) Execute(execution *Execution) error {
	return s.runOutTransition(execution)
}

type EndModel struct {

	NodeModel
}

// todo::
func (e *EndModel) exec(execution *Execution) error {
	//e.Fire(, execution)
	return nil
}

// 工作元素
type WorkModel struct {
	NodeModel

	Form string
}

// 用户自定义处理
// snaker对外提供一个di容器
// 实现接口并注入到容器中
// snaker 处理时调用
type CustomModel struct {
	WorkModel

	// 实例名称
	Clazz string

	// 传入参数
	//Args string

}

// todo::
func (c *CustomModel) exec(execution *Execution) error {
	// 从di容器中查找指定的实例
	return nil
}


// 决策定义decision元素
type DecisionModel struct {
	NodeModel

	// 决策选择表达式串（需要表达式引擎解析）
	Expr string

	// 表达式解析器
	Expression Expression
}

func (d *DecisionModel) exec(execution *Execution) error {
	logs.Info("%d->decision execution.getArgs():%v", 11, execution.Args)

	if nil == d.Expression {
		return errors.New("表达式解析器为空")
	}
	isFound := false
	for _, e := range d.Outputs.Values() {
		tm := e.(*TransitionModel)

		if "" != tm.Expr && d.Expression.Eval(tm.Expr, execution.Args) {
			tm.Enable = true
			tm.Execute(execution)
			isFound = true
		}
	}

	if !isFound {
		return errors.New(fmt.Sprintf("%d->decision节点无法确定下一步执行路线", 11))
	}
	return nil
}

type ForkModel struct {
	NodeModel

}

func (f *ForkModel) exec(execution *Execution)error {
	return f.runOutTransition(execution)
}

type JoinModel struct {
	NodeModel

}

// todo::
//func (j *JoinModel) exec(execution *Execution) error {
//
//}


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

type SubProcessModel struct {
	WorkModel

	ProcessName string

	Version int

	SubProcess *ProcessModel
}

func (s *SubProcessModel) exec(execution *Execution) error {
	return s.runOutTransition(execution)
}

type TaskModel struct {

	PerformType int8

}

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
func (t *TransitionModel) Execute(execution *Execution) error {
	return nil
}
