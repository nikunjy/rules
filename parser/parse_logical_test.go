package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedObject(t *testing.T) {
	tests := []testCase{
		{
			`x.a eq 1 and x.b.c eq 2`,
			map[string]interface{}{
				"x": map[string]interface{}{
					"a": 1,
					"b": map[string]interface{}{
						"c": 2.0,
					},
				},
			},
			true,
			false,
		},
		{
			`x.a eq 1 and x.b.c eq 2`,
			map[string]interface{}{
				"x": map[string]interface{}{
					"a": 1,
					"b": map[string]interface{}{
						"c": 2,
					},
				},
			},
			true,
			false,
		},
		{
			`x.a eq 1 and x.b.c eq 2`,
			map[string]interface{}{
				"x": map[string]interface{}{
					"b": map[string]interface{}{
						"c": 2,
					},
				},
			},
			false,
			false,
		},
		{
			`x.a eq 1 or x.b.c eq 2`,
			map[string]interface{}{
				"x": map[string]interface{}{
					"b": map[string]interface{}{
						"c": 2,
					},
				},
			},
			true,
			false,
		},
		{
			`x.a eq 1 or x.b.c eq 2`,
			nil,
			false,
			false,
		},
		{
			`x.a eq 1 or x.b.c eq 2`,
			obj{},
			false,
			false,
		},
		{
			`x.a eq 1 or x.b.c gt true`,
			obj{},
			false,
			true,
		},
		{
			`x.a eq 1 or x.b.c gt true`,
			map[string]interface{}{
				"x": map[string]interface{}{
					"a": 1,
				},
			},
			true,
			false,
		},
		{
			`x.b.c gt true or x.a eq 1`,
			map[string]interface{}{
				"x": map[string]interface{}{
					"a": 1,
				},
			},
			false,
			true,
		},
	}

	for _, tt := range tests {
		result, err := eval(t, tt.rule, tt.input)
		if tt.hasError {
			assert.Error(t, err)
			continue
		}
		assert.NoError(t, err)
		assert.Equal(t, tt.result, result)
	}
}

func TestLogicalExpWithAnd(t *testing.T) {
	tests := []testCase{
		{
			`x eq 1 or y gt 1`,
			obj{
				"x": 1,
				"y": 0,
			},
			true,
			false,
		},
		{
			`x eq 1 and y gt 1`,
			obj{
				"x": 1,
				"y": 2,
			},
			true,
			false,
		},
		{
			`x eq 1 and y gt 1`,
			obj{
				"x": 1,
				"y": 1,
			},
			false,
			false,
		},
		{
			`x eq 1 and not (y gt 1)`,
			obj{
				"x": 1,
				"y": 1,
			},
			true,
			false,
		},
		{
			`x eq 1 and not (y gt 1)`,
			obj{
				"x": 1,
				"y": 2,
			},
			false,
			false,
		},
		{
			`(x eq 1 and y gt 1) and z eq 3`,
			obj{
				"x": 1,
				"y": 2,
				"z": 3,
			},
			true,
			false,
		},
		{
			`(x eq 1 and y gt 1) and z eq 3 or a gt 4`,
			obj{
				"x": 1,
				"y": 2,
				"a": 5,
			},
			true,
			false,
		},
		{
			`(x eq 1 and y gt 1) and (z eq 3 or a gt 4)`,
			obj{
				"x": 1,
				"y": 2,
				"a": 5,
			},
			true,
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}
}
