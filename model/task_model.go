package model

// performType 参与类型
const (
	Any = iota
	ALL
)

// 任务类型(Major:主办的,Aidant:协助的,Record:仅仅作为记录的)
const (
	Major = iota
	Aidant
	Record
)
type TaskModel struct {

}