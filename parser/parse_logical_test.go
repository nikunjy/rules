package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogicalExpWithAnd(t *testing.T) {
	tests := []testCase{
		{
			`x eq 1 and y gt 1`,
			obj{
				"x": 1,
				"y": 2,
			},
			true,
		},
		{
			`x eq 1 and y gt 1`,
			obj{
				"x": 1,
				"y": 1,
			},
			false,
		},
		{
			`x eq 1 and not (y gt 1)`,
			obj{
				"x": 1,
				"y": 1,
			},
			true,
		},
		{
			`x eq 1 and not (y gt 1)`,
			obj{
				"x": 1,
				"y": 2,
			},
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
		},
		{
			`(x eq 1 and y gt 1) and z eq 3 or a gt 4`,
			obj{
				"x": 1,
				"y": 2,
				"a": 5,
			},
			true,
		},
		{
			`(x eq 1 and y gt 1) and (z eq 3 or a gt 4)`,
			obj{
				"x": 1,
				"y": 2,
				"a": 5,
			},
			true,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}
}
