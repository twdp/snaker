package parse

import (
	"github.com/clbanning/mxj"
	"tianwei.pro/snaker/model"
)

func Parse(content string) (*model.ProcessModel, error) {
	m, err := mxj.NewMapXml([]byte(content))
	if err != nil {
		return nil, err
	}


	return nil, nil
}