package main

import (
	"fmt"

	"github.com/nikunjy/rules/parser"
)

func main() {
	data := `x.y.a eq 1 and x.y.b eq 2 and x.z eq "c"`
	items := map[string]interface{}{
		"name": "nikunj",
		"age":  10,
		"x": map[string]interface{}{
			"y": map[string]interface{}{
				"a": 1,
				"b": 2,
			},
			"z": "c",
		},
	}
	fmt.Println(parser.Evaluate(data, items))
}
