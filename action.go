package snaker

import "snaker/core"

type Action interface {

	// 根据当前的执行对象所维持的process、order、model、args对所属流程实例进行执行
	execute(execution core.Execution)
}