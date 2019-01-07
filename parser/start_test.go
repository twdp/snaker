package parser

import (
	"fmt"
	"testing"
)

func TestStartParserFactory_NewParse(t *testing.T) {
	startParser := &StartParser{}
	startParser.AbstractNodeParser.Parent = startParser

	fmt.Println(startParser.Parse(map[string]interface{}{
		"-name": "dd",
		"-displayName": "fff",
	}))
}