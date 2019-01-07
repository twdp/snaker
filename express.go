package snaker

// 表达式解析接口
type Expression interface {

	// 根据表达式串、参数解析表达式并返回bool
	Eval(expr string, args map[string]interface{}) bool
}