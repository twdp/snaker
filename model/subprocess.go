package model

import "tianwei.pro/snaker/core"

type SubProcessModel struct {
	WorkModel

	ProcessName string

	Version int

	SubProcess *ProcessModel
}

func (s *SubProcessModel) exec(execution *core.Execution) error {
	return s.runOutTransition(execution)
}