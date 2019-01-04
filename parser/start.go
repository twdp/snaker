package parser

import (
	"tianwei.pro/snaker/model"
	"unsafe"
)

type StartParser struct {
	AbstractNodeParser
}

type StartParserFactory struct {

}

func (s *StartParserFactory) NewParse() NodeParser {
	ss := new(StartParser)
	ss.AbstractNodeParser.Parent = ss
	return ss
}

func (s *StartParser) newModel() *model.NodeModel {
	newNode := model.NewNodeModel("", "")
	return (*model.NodeModel) (unsafe.Pointer(&model.StartModel{
		NodeModel: *newNode,
	}))
}
