package snaker

import (
	"errors"
	"snaker/core"
)

type AssignmentHandler interface {

	// 分配参与者方法，可获取到当前的执行对象
	// @param execution 执行对象
	// @return 参与者对象
	assign(execution *core.Execution) (interface{}, error)
}

// 分配参与者的处理抽象类
type Assignment struct {

}

func (a *Assignment) assign(execution *core.Execution) (interface{}, error) {
	return a.modelAssign(nil, execution)
}

func (a *Assignment) modelAssign(model *TaskModel, execution *core.Execution) (interface{}, error) {
	return nil, errors.New("没有实现Assignment.modelAssign")
}