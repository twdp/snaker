package snaker

//流程执行体
type Execution struct {
	Engine         *Engine                //引擎
	Process        *Process               //流程定义对象
	Order          *Order                 //流程实例对象
	ParentOrder    *Order                 //父流程实例
	ParentNodeName string                 //父流程实例节点名称
	ChildOrderId   int64                 //子流程实例节点名称
	Args           map[string]interface{} //执行参数
	Operator       string                 //操作人
	Task           *Task                  //任务
	Tasks          []*Task                //返回的任务列表
	IsMerged       bool                   //是否已合并,针对join节点的处理
}
