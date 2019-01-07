package model

import "tianwei.pro/snaker/core"

type EndModel struct {

	NodeModel
}

// todo::
func (e *EndModel) exec(execution *core.Execution) error {
	//e.Fire(, execution)
	return nil
}