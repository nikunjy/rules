package parser

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCapitalizedLogicalOperators(t *testing.T) {
	input := obj{"x": 1, "y": 99}

	andTrue, err := eval(t, `x eq 1 AND y eq 99`, input)
	require.NoError(t, err)
	assert.True(t, andTrue)

	andFalse, err := eval(t, `x eq 1 AND y eq 2`, input)
	require.NoError(t, err)
	assert.False(t, andFalse)

	orTrue, err := eval(t, `x eq 0 OR y eq 99`, input)
	require.NoError(t, err)
	assert.True(t, orTrue)

	mixed, err := eval(t, `x eq 1 And y eq 99`, input)
	require.NoError(t, err)
	assert.True(t, mixed)
}

func TestLogicalKeywordsInsideStringsAndNames(t *testing.T) {
	result, err := eval(t, `x eq "AND"`, obj{"x": "AND"})
	require.NoError(t, err)
	assert.True(t, result)

	result, err = eval(t, `ANDROID eq 1`, obj{"ANDROID": 1})
	require.NoError(t, err)
	assert.True(t, result)

	result, err = eval(t, `order eq 1`, obj{"order": 1})
	require.NoError(t, err)
	assert.True(t, result)

	result, err = eval(t, `x.AND eq 1`, obj{"x": map[string]interface{}{"AND": 1}})
	require.NoError(t, err)
	assert.True(t, result)

	result, err = eval(t, `x.OR eq 1`, obj{"x": map[string]interface{}{"OR": 1}})
	require.NoError(t, err)
	assert.True(t, result)

	result, err = eval(t, `AND eq 1`, obj{"AND": 1})
	require.NoError(t, err)
	assert.True(t, result)

	result, err = eval(t, `OR eq 1`, obj{"OR": 1})
	require.NoError(t, err)
	assert.True(t, result)
}

func TestInvalidRuleRejectedAtParse(t *testing.T) {
	_, err := NewEvaluator(`this is not a valid rule !!!`)
	require.Error(t, err)

	_, err = NewEvaluator(`x eq 1 leftover`)
	require.Error(t, err)
}

func TestTrailingWhitespaceAccepted(t *testing.T) {
	for _, rule := range []string{
		`x eq 1 `,
		"x eq 1\n",
		"x eq 1 \n",
	} {
		result, err := eval(t, rule, obj{"x": 1})
		require.NoError(t, err, rule)
		assert.True(t, result, rule)
	}
}

func TestNormalizeLogicalKeywordsLeavesAttributeSegments(t *testing.T) {
	assert.Equal(t, `x eq 1 and y eq 2`, normalizeLogicalKeywords(`x eq 1 AND y eq 2`))
	assert.Equal(t, `x eq 1 or y eq 2`, normalizeLogicalKeywords(`x eq 1 OR y eq 2`))
	assert.Equal(t, `x eq 1 and y eq 2`, normalizeLogicalKeywords(`x eq 1 And y eq 2`))
	assert.Equal(t, `x.AND eq 1`, normalizeLogicalKeywords(`x.AND eq 1`))
	assert.Equal(t, `x.OR eq 1`, normalizeLogicalKeywords(`x.OR eq 1`))
	assert.Equal(t, `AND.x eq 1`, normalizeLogicalKeywords(`AND.x eq 1`))
	assert.Equal(t, `AND eq 1`, normalizeLogicalKeywords(`AND eq 1`))
	assert.Equal(t, `x eq "AND"`, normalizeLogicalKeywords(`x eq "AND"`))
	assert.Equal(t, `ANDROID eq 1`, normalizeLogicalKeywords(`ANDROID eq 1`))
}

func TestIntINAcceptsJSONAndWiderInts(t *testing.T) {
	var decoded map[string]interface{}
	require.NoError(t, json.Unmarshal([]byte(`{"x": 1}`), &decoded))

	result, err := eval(t, `x IN [1, 2, 3]`, decoded)
	require.NoError(t, err)
	assert.True(t, result)

	result, err = eval(t, `x IN [1, 2, 3]`, obj{"x": int64(2)})
	require.NoError(t, err)
	assert.True(t, result)

	result, err = eval(t, `x IN [1, 2, 3]`, obj{"x": int32(3)})
	require.NoError(t, err)
	assert.True(t, result)

	result, err = eval(t, `x IN [1, 2, 3]`, obj{"x": 4})
	require.NoError(t, err)
	assert.False(t, result)
}

func TestStringINIsCaseInsensitive(t *testing.T) {
	result, err := eval(t, `x IN ["ABC", "DEF"]`, obj{"x": "abc"})
	require.NoError(t, err)
	assert.True(t, result)

	result, err = eval(t, `x IN ["ABC"]`, obj{"x": "xyz"})
	require.NoError(t, err)
	assert.False(t, result)
}

func TestFloatLEReportsTypeError(t *testing.T) {
	ev, err := NewEvaluator(`x le 1.1`)
	require.NoError(t, err)
	result, err := ev.Process(obj{"x": "not-a-number"})
	assert.False(t, result)
	assert.NoError(t, err)
	assert.Error(t, ev.LastDebugErr())
}

func TestFloatAcceptsWiderInts(t *testing.T) {
	result, err := eval(t, `x eq 1.0`, obj{"x": int64(1)})
	require.NoError(t, err)
	assert.True(t, result)

	result, err = eval(t, `x eq 1.0`, obj{"x": int32(1)})
	require.NoError(t, err)
	assert.True(t, result)
}

func TestNestedNonObjectPath(t *testing.T) {
	result, err := eval(t, `x.a eq 1`, obj{"x": "not-a-map"})
	assert.False(t, result)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "cannot read attribute")
}

func TestFloatINMissingLeft(t *testing.T) {
	ev, err := NewEvaluator(`x IN [1.0, 2.0]`)
	require.NoError(t, err)
	result, err := ev.Process(obj{})
	assert.False(t, result)
	assert.NoError(t, err)
	require.Error(t, ev.LastDebugErr())
	assert.NotPanics(t, func() { _ = ev.LastDebugErr().Error() })
}

func TestErrInvalidOperandNilVal(t *testing.T) {
	err := newErrInvalidOperand(nil, float64(0))
	var msg string
	assert.NotPanics(t, func() { msg = err.Error() })
	assert.Contains(t, msg, "nil")
}
