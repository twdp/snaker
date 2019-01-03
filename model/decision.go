package model

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"tianwei.pro/snaker"
	"tianwei.pro/snaker/core"
)

// 决策定义decision元素
type DecisionModel struct {
	NodeModel

	// 决策选择表达式串（需要表达式引擎解析）
	Expr string

	// 表达式解析器
	Expression snaker.Expression
}

func (d *DecisionModel) exec(execution *core.Execution) error {
	logs.Info("%d->decision execution.getArgs():%v", execution.Instance.Id, execution.Args)

	if nil == d.Expression {
		return errors.New("表达式解析器为空")
	}
	isFound := false
	for e := d.Outputs.Front(); e != nil; e = e.Next() {
		tm := e.Value.(*TransitionModel)

		if "" != tm.Expr && d.Expression.Eval(tm.Expr, execution.Args) {
			tm.Enable = true
			tm.Execute(execution)
			isFound = true
		}
	}

	if !isFound {
		return errors.New(fmt.Sprintf("%d->decision节点无法确定下一步执行路线", execution.Instance.Id))
	}
	return nil
}