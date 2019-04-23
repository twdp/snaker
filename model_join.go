package snaker


//XML合并节点
type JoinModel struct {
	NodeModel
}

//合并分叉节点
func (p *JoinModel) MergeBranchHandle(execution *Execution) {
	activeNodes := FindActiveNodes(p)
	MergeHandle(execution, activeNodes)
}

//执行
func (p *JoinModel) Exec(execution *Execution) {
	p.MergeBranchHandle(execution)
	if execution.IsMerged {
		p.RunOutTransition(execution)
	}
}

//递归查找分叉节点
func FindForkTaskNames(node INodeModel) []string {
	ret := make([]string, 0)
	switch node.(type) {
	case *ForkModel:
	default:
		for _, tm := range node.GetInputs() {
			switch tm.Source.(type) {
			case *SubProcessModel:
				ret = append(ret, tm.Source.(*SubProcessModel).Name)
			case *TaskModel:
				ret = append(ret, tm.Source.(*TaskModel).Name)
			default:
				ret = append(ret, FindForkTaskNames(tm.Source)...)
			}
		}
	}
	return ret
}

//查找分叉节点
func FindActiveNodes(node INodeModel) []string {
	return FindForkTaskNames(node)
}
