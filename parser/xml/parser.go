package xml

import (
	"errors"
	"fmt"
	"github.com/clbanning/mxj"
	"reflect"
	"tianwei.pro/snaker/model"
	"tianwei.pro/snaker/parser"
)



type XmlParser struct {
	// xml 元素解析容器
	elementParserContainer parser.SnakerParserContainer
}

// 解析流程定义文件，并将解析后的对象放入模型容器中
func (x *XmlParser) ParseXml(content string) (*model.ProcessModel, error) {
	if c, err := mxj.NewMapXml([]byte(content)); err != nil {
		return nil, errors.New(fmt.Sprintf("解析xml文件出错, content: %s", content))
	} else {
		// 根元素
		root := c.Old()[parser.RootElement].(map[string]interface{})
		process := model.NewProcess(root[parser.AttrName].(string), root[parser.AttrDisplayName].(string))

		for k, v := range root {
			vv := reflect.ValueOf(v)
			switch vv.Kind() {
			case reflect.Map:
				vvv := v.(map[string]interface{})
				vvv[parser.ElementType] = k
				if m, err := x.parseModel(vvv); err != nil {
					return nil, err
				} else {
					process.Nodes.Add(m)
				}
			case reflect.Slice:
				// 节点类型多个时
				// 是slice类型
				for _, k := range v.([]interface{}) {
					vvv := k.(map[string]interface{})
					vvv[parser.ElementType] = k
					if m, err := x.parseModel(vvv); err != nil {
						return nil, err
					} else {
						process.Nodes.Add(m)
					}
				}
			}
		}

		for _, node := range process.Nodes.Values() {
			nodeModel := node.(*model.NodeModel)
			for _, t := range nodeModel.Outputs.Values() {
				transition := t.(*model.TransitionModel)
				to := transition.To
				for _, node2 := range process.Nodes.Values() {
					nodeModel2 := node2.(*model.NodeModel)
					if to == nodeModel2.Name {
						nodeModel2.Inputs.Add(transition)
						transition.Target = nodeModel2
					}
				}
			}
		}

		return process, nil
	}
	return nil, fmt.Errorf("解析xml失败")
}

// 对流程定义xml的节点，根据其节点对应的解析器解析节点内容
func (x *XmlParser) parseModel(node map[string]interface{}) (*model.NodeModel, error) {
	return x.elementParserContainer.GetNodeParserFactory(node[parser.ElementType].(string)).NewParse().Parse(node)
}
