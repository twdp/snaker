package model

import "tianwei.pro/snaker/core"

type ForkModel struct {
	NodeModel

}

func (f *ForkModel) exec(execution *core.Execution)error {
	return f.runOutTransition(execution)
}