package parser

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"reflect"
	"tianwei.pro/snaker/model"
)

// xml 节点信息
const (
	RootElement = "process"
	// map中节点类型
	ElementType = "elementType"
	// 变迁节点名称
	NodeTransition = "transition"

	// 节点属性名称
	AttrName = "-name"
	AttrDisplayName = "-displayName"
	AttrInstanceUrl = "-instanceUrl"
	AttrInstanceNoClazz = "-instanceNoClass"
	AttrExpr = "-expr"
	AttrHandleClazz = "-handleClass"
	AttrForm = "-form"
	AttrField = "-field"
	AttrValue = "-value"
	AttrAttr = "-attr"
	attrType= "-type"
	AttrAssignee = "-assignee"
	AttrAssignmentHandler = "-assignmentHandler"
	AttrPerormType = "-performType"
	AttrTaskType = "-taskType"
	AttrTo = "-to"
	AttrProcessName = "-processName"
	AttrVersion = "-version"
	AttrExpireTime = "-expireTime"
	AttrAutoExecute = "-autoExecute"
	AttrCallback = "-callback"
	AttrReminderTime = "-reminderTime"
	AttrReminderRepeat = "-reminderRepeat"
	AttrClazz = "-clazz"
	AttrMethodName = "-methodName"
	AttrArgs = "-args"
	AttrVar = "-var"
	AttrLayout = "-layout"
	AttrG = "-g"
	AttrOffset = "-offset"
	AttrPreInterceptors = "-preInterceptors"
	AttrPostInterceptors = "-postInterceptors"
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

type ModelGen interface {
	newModel() *model.NodeModel
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
	//model *model.NodeModel
	Parent ModelGen
}

func (a *AbstractNodeParser) Parse(element map[string]interface{}) (*model.NodeModel, error) {
	m := a.Parent.newModel()
	//a.model = model
	m.Name = element[AttrName].(string)
	m.DisplayName = element[AttrDisplayName].(string)
	// interceptor

	v := element[NodeTransition]
	tms := arraylist.New()

	if  v != nil {
		vv := reflect.ValueOf(v)

		switch vv.Kind() {
		case reflect.Map:
			tms.Add(vv)
		case reflect.Slice:
			for _, k := range v.([]interface{}) {
				tms.Add(k)
			}
		}
	}


	for _,  te := range tms.Values() {
		tte := te.(map[string]interface{})
		transition := &model.TransitionModel{
			BaseModel: model.BaseModel{
				Name: tte[AttrName].(string),
				DisplayName: tte[AttrDisplayName].(string),
			},
			To: tte[AttrTo].(string),
			Expr: tte[AttrExpr].(string),
			Source: m,
		}
		m.Outputs.Add(transition)
	}

	a.parseNode(m, element)

	return m, nil
}

// 子类可覆盖此方法，完成特定的解析
func (a *AbstractNodeParser) parseNode(model *model.NodeModel, element map[string]interface{}) error {
	return nil
}

func (a *AbstractNodeParser) newModel() *model.NodeModel {
	panic("未实现此方法")
}