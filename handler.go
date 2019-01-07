package snaker


// 流程各模型操作处理接口
type IHandler interface {

	// 处理具体的操作
	Handle(execution *Execution) error
}