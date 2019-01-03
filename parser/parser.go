package parser

import (
	"errors"
	"fmt"
	"github.com/clbanning/mxj"
	"tianwei.pro/snaker/model"
)

// xml 节点信息
const (
	// 变迁节点名称
	NODE_TRANSITION = "transition"

	// 节点属性名称
	ATTR_NAME = "-name"
	ATTR_DISPLAYNAME = "-displayName"
	ATTR_INSTANCEURL = "-instanceUrl"
	ATTR_INSTANCENOCLASS = "-instanceNoClass"
	ATTR_EXPR = "-expr"
	ATTR_HANDLECLASS = "-handleClass"
	ATTR_FORM = "-form"
	ATTR_FIELD = "-field"
	ATTR_VALUE = "-value"
	ATTR_ATTR = "-attr"
	ATTR_TYPE= "-type"
	ATTR_ASSIGNEE = "-assignee"
	ATTR_ASSIGNEE_HANDLER = "-assignmentHandler"
	ATTR_PERFORMTYPE = "-performType"
	ATTR_TASKTYPE = "-taskType"
	ATTR_TO = "-to"
	ATTR_PROCESSNAME = "-processName"
	ATTR_VERSION = "-version"
	ATTR_EXPIRETIME = "-expireTime"
	ATTR_AUTOEXECUTE = "-autoExecute"
	ATTR_CALLBACK = "-callback"
	ATTR_REMINDERTIME = "-reminderTime"
	ATTR_REMINDERREPEAT = "-reminderRepeat"
	ATTR_CLAZZ = "-clazz"
	ATTR_METHODNAME = "-methodName"
	ATTR_ARGS = "-args"
	ATTR_VAR = "-var"
	ATTR_LAYOUT = "-layout"
	ATTR_G = "-g"
	ATTR_OFFSET = "-offset"
	ATTR_PREINTERCEPTORS = "-preInterceptors"
	ATTR_POSTINTERCEPTORS = "-postInterceptors"
)

//
type SnakerParserContainer interface {

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

// 解析流程定义文件，并将解析后的对象放入模型容器中
func ParseXml(content string) (*model.ProcessModel, error) {
	if c, err := mxj.NewMapXml([]byte(content)); err != nil {
		return nil, errors.New(fmt.Sprintf("解析xml文件出错, content: %s", content))
	} else {
		fmt.Println(c)
	}
	return nil, nil
}

// 对流程定义xml的节点，根据其节点对应的解析器解析节点内容
func parseMode(node map[string]interface{}) (*model.NodeModel, error) {

}

