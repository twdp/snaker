package parser

import (
	"fmt"
	"tianwei.pro/snaker/model"
)

//
type SnakerParserContainer interface {

	// 添加解析
	AddParserFactory(elementName string, f NodeParserFactory)

	// 根据element名称获取对应的工厂
	GetNodeParserFactory(elementName string) NodeParserFactory
}

// engine上挂载factory
type NodeParserFactory interface {

	// 根据elementName查找使用哪个parser
	NewParse() NodeParser
}

// 节点解析接口
type NodeParser interface {

	// 节点dom元素解析方法，由实现类完成解析
	Parse(element map[string]interface{}) (*model.NodeModel, error)
}

type DefaultSnakerParserContainer struct {
	container map[string]NodeParserFactory
}

func (d *DefaultSnakerParserContainer) AddParserFactory(elementName string, f NodeParserFactory) {
	d.container[elementName] = f
}

func (d *DefaultSnakerParserContainer) GetNodeParserFactory(elementName string) NodeParserFactory {
	if f, ok := d.container[elementName]; ok {
		return f
	} else {
		panic(fmt.Sprintf("[%s]没有对应的解析工厂类", elementName))
	}
}

func NewDefaultSnakerParserContainer() *DefaultSnakerParserContainer {
	return &DefaultSnakerParserContainer{
		container: make(map[string]NodeParserFactory),
	}
}

type AbstractNodeParser struct {
	model *model.NodeModel
}

func (a *AbstractNodeParser) Parse(element map[string]interface{}) (*model.NodeModel, error) {
	panic("implement me")
}
