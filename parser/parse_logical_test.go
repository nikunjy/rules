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

func TestIssue42VersionAndLogicalOrder(t *testing.T) {
	input := obj{
		"x": "abc",
		"y": "11.3.0",
	}
	parenInput := obj{
		"x": "2.2.0",
		"y": "z",
	}
	tests := []testCase{
		{
			`y gt 11.1.0-0 and not(x IN ["abc","cde","fgh"])`,
			input,
			false,
			false,
		},
		{
			`not(x IN ["abc","cde","fgh"]) and y gt 11.1.0-0`,
			input,
			false,
			false,
		},
		{
			`y gt 11.1.0-0 and not(x IN ["cde","fgh"])`,
			input,
			true,
			false,
		},
		{
			`not(x IN ["cde","fgh"]) and y gt 11.1.0-0`,
			input,
			true,
			false,
		},
		{
			`x gt 2.1.0-0 and not(y in ["a", "b", "c"])`,
			parenInput,
			true,
			false,
		},
		{
			`(x gt 2.1.0-0) and not(y in ["a", "b", "c"])`,
			parenInput,
			true,
			false,
		},
		{
			`((x gt 2.1.0-0) and not(y in ["a", "b", "c"]))`,
			parenInput,
			true,
			false,
		},
		{
			`(x gt 2.1.0-0 and not(y in ["a", "b", "c"]))`,
			parenInput,
			true,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.rule, func(t *testing.T) {
			result, err := eval(t, tt.rule, tt.input)
			assert.NoError(t, err, tt.rule)
			assert.Equal(t, tt.result, result, tt.rule)
			assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
			assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
		})
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
		t.Run(tt.rule, func(t *testing.T) {
			assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
			assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
		})
	}
}

func TestLogicalExpWithAndWithExtraSpaces(t *testing.T) {
	tests := []testCase{
		{
			`x < 2 and ( y > 4 or z == true )`,
			obj{
				"x": 1,
				"y": 3,
				"z": true,
			},
			true,
			false,
		},
		{
			`x < 2 and ( y > 4 or z == true )`,
			obj{
				"x": 1,
				"y": 3,
			},
			false,
			false,
		},
		{
			`x < 2 and ( y > 4 or z == true )`,
			obj{
				"x": 1,
				"y": 5,
			},
			true,
			false,
		},
		{
			`x < 2 and ( y > 4 or z == true)`,
			obj{
				"x": 1,
				"y": 3,
				"z": true,
			},
			true,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.rule, func(t *testing.T) {
			assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
			assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
		})
	}
}
