package model

type FieldNode struct {
	BaseModel

	AttrMap map[string]interface{}


}

func (f *FieldNode) AddAttr(key string, value interface{}) {
	f.AttrMap[key] = value
}