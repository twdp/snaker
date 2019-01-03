package snaker

import "tianwei.pro/snaker/parser"

type Engine interface {

}

type SnakerEngine struct {

	// xml 元素解析容器
	elementParserContainer parser.SnakerParserContainer
}