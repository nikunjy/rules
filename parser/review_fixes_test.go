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
}

func TestInvalidRuleRejectedAtParse(t *testing.T) {
	_, err := NewEvaluator(`this is not a valid rule !!!`)
	require.Error(t, err)

	_, err = NewEvaluator(`x eq 1 leftover`)
	require.Error(t, err)
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
