package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	rule  string
	input map[string]interface{}

	result   bool
	hasError bool
}

type obj map[string]interface{}

func eval(t *testing.T, rule string, input obj) (bool, error) {
	ev, err := NewEvaluator(rule)
	assert.NoError(t, err)
	return ev.Process(input)
}

func TestInvalidRule(t *testing.T) {
	invalidRules := []string{
		`x eq 1 and y ~ 1`,
		`y ~ 1`,
		`y ~ 1 and x eq true`,
	}

	for _, rule := range invalidRules {
		result, err := eval(t, rule, obj{"x": 1})
		assert.Error(t, err, rule)
		assert.False(t, result)

	}
}

func TestVersions(t *testing.T) {
	tests := []testCase{
		{
			`x EQ 1.0.0`,
			obj{
				"x": "1.0.0",
			},
			true,
			false,
		},
		{
			`x EQ 1.0.0`,
			obj{
				"x": 2,
			},
			false,
			false,
		},
		{
			`x EQ 1.0.0`,
			obj{
				"y": "1.0.0",
			},
			false,
			false,
		},
		{
			`x EQ 1.0.0`,
			obj{},
			false,
			false,
		},
		{
			`x EQ 1.0.0`,
			obj{
				"x": "1.1.1",
			},
			false,
			false,
		},
		{
			`x NE 1.0.0`,
			obj{
				"x": "1.0.0.",
			},
			false,
			false,
		},
		{
			`x NE 1.0.0`,
			obj{
				"x": 2.3,
			},
			false,
			false,
		},
		{
			`x NE 1.0.0`,
			obj{
				"x": "1.1.0",
			},
			true,
			false,
		},
		{
			`x LT 1.1.0`,
			obj{
				"x": "1.0.0",
			},
			true,
			false,
		},
		{
			`x LT 1.1.0`,
			obj{
				"x": "2.0.0",
			},
			false,
			false,
		},
		{
			`x LT 1.1.0`,
			obj{
				"x": 1,
			},
			false,
			false,
		},
		{
			`x LE 1.1.0`,
			obj{
				"x": "1.1.0",
			},
			true,
			false,
		},
		{
			`x LE 1.1.0`,
			obj{
				"x": "1.0.0",
			},
			true,
			false,
		},
		{
			`x LE 1.1.0`,
			obj{
				"x": "2.0.0",
			},
			false,
			false,
		},
		{
			`x LE 1.1.0`,
			obj{
				"x": 2,
			},
			false,
			false,
		},
		{
			`x GT 1.1.0`,
			obj{
				"x": "2.0.0",
			},
			true,
			false,
		},
		{
			`x GT 1.1.0`,
			obj{
				"x": "1.0.0",
			},
			false,
			false,
		},
		{
			`x GT 1.1.0`,
			obj{
				"x": 1,
			},
			false,
			false,
		},
		{
			`x GE 1.1.0`,
			obj{
				"x": "2.0.0",
			},
			true,
			false,
		},
		{
			`x GE 1.1.0`,
			obj{
				"x": "1.0.0",
			},
			false,
			false,
		},
		{
			`x GE 1.1.0`,
			obj{
				"x": "1.1.0",
			},
			true,
			false,
		},
		{
			`x GE 1.1.0`,
			obj{
				"x": 1,
			},
			false,
			false,
		},
		{
			`x CO 1.1.0`,
			obj{
				"x": "1.0.0",
			},
			true,
			true,
		},
		{
			`x SW 1.1.0`,
			obj{
				"x": "1.0.0",
			},
			true,
			true,
		},
		{
			`x EW 1.1.0`,
			obj{
				"x": "1.0.0",
			},
			true,
			true,
		},
		{
			`x SW 1.1.0`,
			obj{
				"x": "1.0.0",
			},
			true,
			true,
		},
	}

	for _, tt := range tests {
		result, err := eval(t, tt.rule, tt.input)
		if tt.hasError {
			assert.Error(t, err)
			continue
		} else {
			assert.NoError(t, err)
			assert.Equal(t, result, tt.result, fmt.Sprintf("rule :%s, input :%v", tt.rule, tt.input))
		}
		assert.Equal(t, false, Evaluate(tt.rule, obj{"x": 4.5}), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}
}

func TestListOfStrings(t *testing.T) {
	tests := []testCase{
		{
			`x IN ["abc", "cde", "fgh"]`,
			obj{
				"x": "abc",
			},
			true,
			false,
		},
		{
			`y eq 5 and x IN ["abc", "cde", "fgh"]`,
			obj{
				"x": "abc",
				"y": 5,
			},
			true,
			false,
		},
		{
			`y eq 5 and (x IN ["abc","cde","fgh"])`,
			obj{
				"x": "cde",
				"y": 5,
			},
			true,
			false,
		},
		{
			`y eq 5 and not (x IN ["abc","cde","fgh"])`,
			obj{
				"x": "abc",
				"y": 5,
			},
			false,
			false,
		},
		{
			`x IN ["abc", "cde", "fgh"]`,
			obj{
				"x": "xyz",
			},
			false,
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
			`x IN [1.1, 2.2, 3.3]`,
			obj{
				"x": 1.1,
			},
			true,
			false,
		},
		{
			`y eq 5 and x IN [1.1, 2.1, 3.3]`,
			obj{
				"x": 1.1,
				"y": 5,
			},
			true,
			false,
		},
		{
			`y eq 5 and (x IN [1.4, 2.1, 3.3])`,
			obj{
				"x": 1.4,
				"y": 5,
			},
			true,
			false,
		},
		{
			`y eq 5 and not (x IN [1.1, 2.2, 3.3])`,
			obj{
				"x": 1.1,
				"y": 5,
			},
			false,
			false,
		},
		{
			`x IN [1.0, 2.2, 3.3]`,
			obj{
				"x": 4.0,
			},
			false,
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, false, Evaluate(tt.rule, obj{"x": "abc"}), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}

}

func TestListOfInts(t *testing.T) {
	tests := []testCase{
		{
			`x IN [1, 2, 3]`,
			obj{
				"x": 1,
			},
			true,
			false,
		},
		{
			`y eq 5 and x IN [1, 2, 3]`,
			obj{
				"x": 1,
				"y": 5,
			},
			true,
			false,
		},
		{
			`y eq 5 and (x IN [1, 2, 3])`,
			obj{
				"x": 1,
				"y": 5,
			},
			true,
			false,
		},
		{
			`y eq 5 and not (x IN [1, 2, 3])`,
			obj{
				"x": 1,
				"y": 5,
			},
			false,
			false,
		},
		{
			`x IN [1, 2, 3]`,
			obj{
				"x": 4,
			},
			false,
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, false, Evaluate(tt.rule, obj{"x": "123"}), tt.rule)
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
			false,
		},
		{
			`x pr`,
			obj{},
			false,
			false,
		},
		{
			`not (x pr)`,
			obj{
				"x": true,
			},
			false,
			false,
		},
		{
			`NOT (x pr)`,
			obj{
				"x": true,
			},
			false,
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
			false,
		},
		{
			`x EQ null`,
			obj{
				"x": true,
			},
			false,
			false,
		},
		{
			`x ne null`,
			obj{
				"x": true,
			},
			true,
			false,
		},
		{
			`x eq null`,
			obj{
				"x": nil,
			},
			true,
			false,
		},
		{
			`x eq null`,
			obj{},
			true,
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}
}

func TestNullInvalids(t *testing.T) {
	badRules := []string{
		`x lt null`,
		`x le null`,
		`x ge null`,
		`x gt null`,
		`x co null`,
		`x sw null`,
		`x ew null`,
		`x in null`,
	}
	for _, rule := range badRules {
		_, err := eval(t, rule, obj{
			"x": nil,
		})
		assert.Error(t, err)
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
			false,
		},
		{
			`x ne true`,
			obj{
				"x": false,
			},
			true,
			false,
		},
		{
			`x eq true`,
			obj{
				"x": 1,
			},
			false,
			false,
		},
		{
			`x ne true`,
			obj{
				"x": 2,
			},
			false,
			false,
		},
		{
			`x eq false`,
			obj{
				"x": true,
			},
			false,
			false,
		},
		{
			`xy eq true`,
			obj{
				"x": true,
			},
			false,
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input))
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}
}

func TestBoolInvalids(t *testing.T) {
	badRules := []string{
		`x le true`,
		`x ge true`,
		`x gt true`,
		`x lt true`,
		`x le true`,
		`x co true`,
		`x sw true`,
		`x ew true`,
		`x in true`,
	}

	for _, rule := range badRules {
		result, err := eval(t, rule, obj{
			"x": true,
		})
		assert.Error(t, err)
		assert.False(t, result)
	}
}

func TestIntInvalids(t *testing.T) {
	badRules := []string{
		`x co 1`,
		`x sw 1`,
		`x ew 1`,
	}

	for _, rule := range badRules {
		result, err := eval(t, rule, obj{
			"x": 2,
		})
		assert.Error(t, err)
		assert.False(t, result)
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
			false,
		},
		{
			`x == 1`,
			obj{
				"x": 1,
			},
			true,
			false,
		},
		{
			`x eq 1`,
			obj{
				"x": "1",
			},
			false,
			false,
		},
		{
			`x eq 1`,
			obj{
				"y": 1,
			},
			false,
			false,
		},
		{
			`x ne 1`,
			obj{
				"x": 1,
			},
			false,
			false,
		},
		{
			`x != 1`,
			obj{
				"x": 1,
			},
			false,
			false,
		},
		{
			`x ne 1`,
			obj{
				"x": "1",
			},
			false,
			false,
		},
		{
			`x != 1`,
			obj{
				"x": "1",
			},
			false,
			false,
		},
		{
			`x NE 1`,
			obj{
				"x": 1,
			},
			false,
			false,
		},
		{
			`x le 1`,
			obj{
				"x": 0,
			},
			true,
			false,
		},
		{
			`x <= 1`,
			obj{
				"x": 0,
			},
			true,
			false,
		},
		{
			`x le 1`,
			obj{
				"x": "1",
			},
			false,
			false,
		},
		{
			`x LE 1`,
			obj{
				"x": 1,
			},
			true,
			false,
		},
		{
			`x le 1`,
			obj{
				"x": 2,
			},
			false,
			false,
		},
		{
			`x LT 1`,
			obj{
				"x": 0,
			},
			true,
			false,
		},
		{
			`x < 1`,
			obj{
				"x": 0,
			},
			true,
			false,
		},
		{
			`x lt 1`,
			obj{
				"x": 1,
			},
			false,
			false,
		},
		{
			`x lt 1`,
			obj{
				"x": "1",
			},
			false,
			false,
		},
		{
			`x < 1`,
			obj{
				"x": "1",
			},
			false,
			false,
		},
		{
			`x GT 1`,
			obj{
				"x": 2,
			},
			true,
			false,
		},
		{
			`x > 1`,
			obj{
				"x": 2,
			},
			true,
			false,
		},
		{
			`x gt 1`,
			obj{
				"x": 1,
			},
			false,
			false,
		},
		{
			`x gt 1`,
			obj{
				"x": "1",
			},
			false,
			false,
		},
		{
			`x > 1`,
			obj{
				"x": "1",
			},
			false,
			false,
		},
		{
			`x GE 1`,
			obj{
				"x": 2,
			},
			true,
			false,
		},
		{
			`x >= 1`,
			obj{
				"x": 2,
			},
			true,
			false,
		},
		{
			`x ge 1`,
			obj{
				"x": 1,
			},
			true,
			false,
		},
		{
			`x ge 1`,
			obj{
				"x": "1",
			},
			false,
			false,
		},
		{
			`x ge 1`,
			obj{
				"x": 0,
			},
			false,
			false,
		},
		{
			`x >= 1`,
			obj{
				"x": 0,
			},
			false,
			false,
		},
		{
			`x NE 1`,
			obj{
				"x": 2,
			},
			true,
			false,
		},
		{
			`x != 1`,
			obj{
				"x": 2,
			},
			true,
			false,
		},
		{
			`x eq 1`,
			obj{
				"y": 1.0,
			},
			false,
			false,
		},
		{
			`x == 1`,
			obj{
				"y": 1.0,
			},
			false,
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
			false,
		},
		{
			`x eq 1.0`,
			obj{
				"y": 1.0,
			},
			false,
			false,
		},
		{
			`x ne 1.1`,
			obj{
				"x": 1.1,
			},
			false,
			false,
		},
		{
			`x le 1.1`,
			obj{
				"x": 0,
			},
			true,
			false,
		},
		{
			`x le 1.1`,
			obj{
				"x": 1.1,
			},
			true,
			false,
		},
		{
			`x le 1.1`,
			obj{
				"x": 2.0,
			},
			false,
			false,
		},
		{
			`x lt 1.1`,
			obj{
				"x": 0.0,
			},
			true,
			false,
		},
		{
			`x lt 1.0`,
			obj{
				"x": 1,
			},
			false,
			false,
		},
		{
			`x gt 1.1`,
			obj{
				"x": 2,
			},
			true,
			false,
		},
		{
			`x gt 1.0`,
			obj{
				"x": 1,
			},
			false,
			false,
		},
		{
			`x ge 1.2`,
			obj{
				"x": 2,
			},
			true,
			false,
		},
		{
			`x ge 1.0`,
			obj{
				"x": 1,
			},
			true,
			false,
		},
		{
			`x ge 1.0`,
			obj{
				"x": 0,
			},
			false,
			false,
		},
		{
			`x ne 1.0`,
			obj{
				"x": 2,
			},
			true,
			false,
		},
		{
			`x eq 1.0`,
			obj{
				"y": 1.0,
			},
			false,
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
		assert.Equal(t, false, Evaluate(tt.rule, obj{"x": "2.4"}), tt.rule)
	}
}

func TestString(t *testing.T) {
	tests := []testCase{
		{
			`x eq ""`,
			obj{},
			false,
			false,
		},
		{
			`x eq ""`,
			obj{
				"x": "",
			},
			true,
			false,
		},
		{
			`x eq "abc"`,
			obj{
				"x": "abc",
			},
			true,
			false,
		},
		{
			`x eq "abc"`,
			obj{
				"y": "abc",
			},
			false,
			false,
		},
		{
			`x ne "abc"`,
			obj{
				"x": "abc",
			},
			false,
			false,
		},
		{
			`x le "abc"`,
			obj{
				"x": "abc",
			},
			true,
			false,
		},
		{
			`x le "cde"`,
			obj{
				"x": "abc",
			},
			true,
			false,
		},
		{
			`x le "cde"`,
			obj{
				"x": "efg",
			},
			false,
			false,
		},
		{
			`x lt "cde"`,
			obj{
				"x": "abc",
			},
			true,
			false,
		},
		{
			`x lt "cde"`,
			obj{
				"x": "cde",
			},
			false,
			false,
		},
		{
			`x gt "cde"`,
			obj{
				"x": "def",
			},
			true,
			false,
		},
		{
			`x gt "cde"`,
			obj{
				"x": "abc",
			},
			false,
			false,
		},
		{
			`x ge "cde"`,
			obj{
				"x": "cde",
			},
			true,
			false,
		},
		{
			`x ge "cde"`,
			obj{
				"x": "def",
			},
			true,
			false,
		},
		{
			`x ge "cde"`,
			obj{
				"x": "abc",
			},
			false,
			false,
		},
		{
			`x ne "abc"`,
			obj{
				"x": "cde",
			},
			true,
			false,
		},
		{
			`x co "ab"`,
			obj{
				"x": "abc",
			},
			true,
			false,
		},
		{
			`x co "ab"`,
			obj{
				"x": "bbc",
			},
			false,
			false,
		},
		{
			`x CO "ab"`,
			obj{
				"x": "bbc",
			},
			false,
			false,
		},
		{
			`x sw "ab"`,
			obj{
				"x": "abc",
			},
			true,
			false,
		},
		{
			`x SW "ab"`,
			obj{
				"x": "abc",
			},
			true,
			false,
		},
		{
			`x sw "ab"`,
			obj{
				"x": "bbc",
			},
			false,
			false,
		},
		{
			`x ew "ab"`,
			obj{
				"x": "bab",
			},
			true,
			false,
		},
		{
			`x EW "ab"`,
			obj{
				"x": "bbc",
			},
			false,
			false,
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule)
		assert.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
		assert.Equal(t, false, Evaluate(tt.rule, obj{"x": 3}), tt.rule)
	}
}
