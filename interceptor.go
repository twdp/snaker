package snaker

import "tianwei.pro/snaker/core"

// 任务拦截器，对产生的任务结果进行拦截
type Interceptor interface {

	// 拦截方法，参数为执行对象
	Intercept(execution *core.Execution) error
}