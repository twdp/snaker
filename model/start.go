package model

import "tianwei.pro/snaker/core"

type StartModel struct {
	NodeModel
}

func (s *StartModel) Execute(execution *core.Execution) error {
	return s.runOutTransition(execution)
}