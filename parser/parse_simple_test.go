package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	rule  string
	input map[string]interface{}

	result bool
}

type obj map[string]interface{}

func TestListOfStrings(t *testing.T) {
	tests := []testCase{
		{
			`x IN ["abc" "cde" "fgh"]`,
			obj{
				"x": "abc",
			},
			true,
		},
		{
			`y eq 5 and x IN ["abc" "cde" "fgh"]`,
			obj{
				"x": "abc",
				"y": 5,
			},
			true,
		},
		{
			`y eq 5 and (x IN ["abc" "cde" "fgh"])`,
			obj{
				"x": "abc",
				"y": 5,
			},
			true,
		},
		{
			`y eq 5 and not (x IN ["abc" "cde" "fgh"])`,
			obj{
				"x": "abc",
				"y": 5,
			},
			false,
		},
		{
			`x IN ["abc" "cde" "fgh"]`,
			obj{
				"x": "xyz",
			},
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}

}

func TestListOfDoubles(t *testing.T) {
	tests := []testCase{
		{
			`x IN [1.1 2.2 3.3]`,
			obj{
				"x": 1.1,
			},
			true,
		},
		{
			`y eq 5 and x IN [1.1 2.1 3.3]`,
			obj{
				"x": 1.1,
				"y": 5,
			},
			true,
		},
		{
			`y eq 5 and (x IN [1.4 2.1 3.3])`,
			obj{
				"x": 1.4,
				"y": 5,
			},
			true,
		},
		{
			`y eq 5 and not (x IN [1.1 2.2 3.3])`,
			obj{
				"x": 1.1,
				"y": 5,
			},
			false,
		},
		{
			`x IN [1.0 2.2 3.3]`,
			obj{
				"x": 4.0,
			},
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}

}

func TestListOfInts(t *testing.T) {
	tests := []testCase{
		{
			`x IN [1 2 3]`,
			obj{
				"x": 1,
			},
			true,
		},
		{
			`y eq 5 and x IN [1 2 3]`,
			obj{
				"x": 1,
				"y": 5,
			},
			true,
		},
		{
			`y eq 5 and (x IN [1 2 3])`,
			obj{
				"x": 1,
				"y": 5,
			},
			true,
		},
		{
			`y eq 5 and not (x IN [1 2 3])`,
			obj{
				"x": 1,
				"y": 5,
			},
			false,
		},
		{
			`x IN [1 2 3]`,
			obj{
				"x": 4,
			},
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}

}

func TestPresent(t *testing.T) {
	tests := []testCase{
		{
			`x pr`,
			obj{
				"x": true,
			},
			true,
		},
		{
			`x pr`,
			obj{},
			false,
		},
		{
			`not (x pr)`,
			obj{
				"x": true,
			},
			false,
		},
		{
			`NOT (x pr)`,
			obj{
				"x": true,
			},
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}

}

func TestNull(t *testing.T) {
	tests := []testCase{
		{
			`x eq null`,
			obj{
				"x": true,
			},
			false,
		},
		{
			`x EQ null`,
			obj{
				"x": true,
			},
			false,
		},
		{
			`x ne null`,
			obj{
				"x": true,
			},
			true,
		},
		{
			`x eq null`,
			obj{
				"x": nil,
			},
			true,
		},
		{
			`x eq null`,
			obj{},
			true,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}
}

func TestBool(t *testing.T) {
	tests := []testCase{
		{
			`x eq true`,
			obj{
				"x": true,
			},
			true,
		},
		{
			`x eq false`,
			obj{
				"x": true,
			},
			false,
		},
		{
			`xy eq true`,
			obj{
				"x": true,
			},
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input))
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}
}

func TestInt(t *testing.T) {
	tests := []testCase{
		{
			`x eq 1`,
			obj{
				"x": 1,
			},
			true,
		},
		{
			`x eq 1`,
			obj{
				"y": 1,
			},
			false,
		},
		{
			`x ne 1`,
			obj{
				"x": 1,
			},
			false,
		},
		{
			`x NE 1`,
			obj{
				"x": 1,
			},
			false,
		},
		{
			`x le 1`,
			obj{
				"x": 0,
			},
			true,
		},
		{
			`x LE 1`,
			obj{
				"x": 1,
			},
			true,
		},
		{
			`x le 1`,
			obj{
				"x": 2,
			},
			false,
		},
		{
			`x LT 1`,
			obj{
				"x": 0,
			},
			true,
		},
		{
			`x lt 1`,
			obj{
				"x": 1,
			},
			false,
		},
		{
			`x GT 1`,
			obj{
				"x": 2,
			},
			true,
		},
		{
			`x gt 1`,
			obj{
				"x": 1,
			},
			false,
		},
		{
			`x GE 1`,
			obj{
				"x": 2,
			},
			true,
		},
		{
			`x ge 1`,
			obj{
				"x": 1,
			},
			true,
		},
		{
			`x ge 1`,
			obj{
				"x": 0,
			},
			false,
		},
		{
			`x NE 1`,
			obj{
				"x": 2,
			},
			true,
		},
		{
			`x eq 1`,
			obj{
				"y": 1.0,
			},
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}
}

func TestFloat(t *testing.T) {
	tests := []testCase{
		{
			`x eq 1.0`,
			obj{
				"x": 1.0,
			},
			true,
		},
		{
			`x eq 1.0`,
			obj{
				"y": 1.0,
			},
			false,
		},
		{
			`x ne 1.1`,
			obj{
				"x": 1.1,
			},
			false,
		},
		{
			`x le 1.1`,
			obj{
				"x": 0,
			},
			true,
		},
		{
			`x le 1.1`,
			obj{
				"x": 1.1,
			},
			true,
		},
		{
			`x le 1.1`,
			obj{
				"x": 2.0,
			},
			false,
		},
		{
			`x lt 1.1`,
			obj{
				"x": 0.0,
			},
			true,
		},
		{
			`x lt 1.0`,
			obj{
				"x": 1,
			},
			false,
		},
		{
			`x gt 1.1`,
			obj{
				"x": 2,
			},
			true,
		},
		{
			`x gt 1.0`,
			obj{
				"x": 1,
			},
			false,
		},
		{
			`x ge 1.2`,
			obj{
				"x": 2,
			},
			true,
		},
		{
			`x ge 1.0`,
			obj{
				"x": 1,
			},
			true,
		},
		{
			`x ge 1.0`,
			obj{
				"x": 0,
			},
			false,
		},
		{
			`x ne 1.0`,
			obj{
				"x": 2,
			},
			true,
		},
		{
			`x eq 1.0`,
			obj{
				"y": 1.0,
			},
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}
}

func TestString(t *testing.T) {
	tests := []testCase{
		{
			`x eq "abc"`,
			obj{
				"x": "abc",
			},
			true,
		},
		{
			`x eq "abc"`,
			obj{
				"y": "abc",
			},
			false,
		},
		{
			`x ne "abc"`,
			obj{
				"x": "abc",
			},
			false,
		},
		{
			`x le "abc"`,
			obj{
				"x": "abc",
			},
			true,
		},
		{
			`x le "cde"`,
			obj{
				"x": "abc",
			},
			true,
		},
		{
			`x le "cde"`,
			obj{
				"x": "efg",
			},
			false,
		},
		{
			`x lt "cde"`,
			obj{
				"x": "abc",
			},
			true,
		},
		{
			`x lt "cde"`,
			obj{
				"x": "cde",
			},
			false,
		},
		{
			`x gt "cde"`,
			obj{
				"x": "def",
			},
			true,
		},
		{
			`x gt "cde"`,
			obj{
				"x": "abc",
			},
			false,
		},
		{
			`x ge "cde"`,
			obj{
				"x": "cde",
			},
			true,
		},
		{
			`x ge "cde"`,
			obj{
				"x": "def",
			},
			true,
		},
		{
			`x ge "cde"`,
			obj{
				"x": "abc",
			},
			false,
		},
		{
			`x ne "abc"`,
			obj{
				"x": "cde",
			},
			true,
		},
		{
			`x co "ab"`,
			obj{
				"x": "abc",
			},
			true,
		},
		{
			`x co "ab"`,
			obj{
				"x": "bbc",
			},
			false,
		},
		{
			`x CO "ab"`,
			obj{
				"x": "bbc",
			},
			false,
		},
		{
			`x sw "ab"`,
			obj{
				"x": "abc",
			},
			true,
		},
		{
			`x SW "ab"`,
			obj{
				"x": "abc",
			},
			true,
		},
		{
			`x sw "ab"`,
			obj{
				"x": "bbc",
			},
			false,
		},
		{
			`x ew "ab"`,
			obj{
				"x": "bab",
			},
			true,
		},
		{
			`x EW "ab"`,
			obj{
				"x": "bbc",
			},
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}
}
