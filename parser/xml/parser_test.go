package xml

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"testing"
	"tianwei.pro/snaker/model"
	"tianwei.pro/snaker/parser"
)

type AParserFactory struct {


}

func (a *AParserFactory) NewParse() parser.NodeParser {
	return &AParser{}
}

type AParser struct {
	parser.NodeParser
}



func (a *AParser) Parse(element map[string]interface{}) (*model.NodeModel, error) {
	return &model.NodeModel{
		BaseModel: model.BaseModel{
			Name: element[attrName].(string),
			DisplayName: element[attrDisplayName].(string),
		},
		Inputs:arraylist.New(),
		Outputs:arraylist.New(),
	}, nil
}

func TestXmlParser_ParseXml(t *testing.T) {
	x:= `<process name="t" displayName="tt">
		<a name="test" displayName="测试一下">
		<b/>
	</a></process>`
	parse := &XmlParser{
		elementParserContainer: parser.NewDefaultSnakerParserContainer(),
	}
	parse.elementParserContainer.AddParserFactory("a", &AParserFactory{})

	model, err := parse.ParseXml(x)
	fmt.Println(model, err)
	fmt.Println(model.Nodes.Size())

}