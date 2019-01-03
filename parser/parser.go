package parser

// xml 节点信息
const (
	// 变迁节点名称
	NODE_TRANSITION = "transition"

	// 节点属性名称
	ATTR_NAME = "name"
	ATTR_DISPLAYNAME = "displayName"
	ATTR_INSTANCEURL = "instanceUrl"
	ATTR_INSTANCENOCLASS = "instanceNoClass"
	ATTR_EXPR = "expr"
	ATTR_HANDLECLASS = "handleClass"
	ATTR_FORM = "form"
	ATTR_FIELD = "field"
	ATTR_VALUE = "value"
	ATTR_ATTR = "attr"
	ATTR_TYPE= "type"
	ATTR_ASSIGNEE = "assignee"
	ATTR_ASSIGNEE_HANDLER = "assignmentHandler"
	ATTR_PERFORMTYPE = "performType"
	ATTR_TASKTYPE = "taskType"
	ATTR_TO = "to"
	ATTR_PROCESSNAME = "processName"
	ATTR_VERSION = "version"
	ATTR_EXPIRETIME = "expireTime"
	ATTR_AUTOEXECUTE = "autoExecute"
	ATTR_CALLBACK = "callback"
	ATTR_REMINDERTIME = "reminderTime"
	ATTR_REMINDERREPEAT = "reminderRepeat"
	ATTR_CLAZZ = "clazz"
	ATTR_METHODNAME = "methodName"
	ATTR_ARGS = "args"
	ATTR_VAR = "var"
	ATTR_LAYOUT = "layout"
	ATTR_G = "g"
	ATTR_OFFSET = "offset"
	ATTR_PREINTERCEPTORS = "preInterceptors"
	ATTR_POSTINTERCEPTORS = "postInterceptors"
)

// 节点解析接口
type NodeParser interface {

	// 节点dom元素解析方法，由实现类完成解析
	Parse(element map[string]interface{}) (*NodeModel)


}