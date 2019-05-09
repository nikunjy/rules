package parser

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedError(t *testing.T) {
	err := errors.New("some random error")
	ne := newNestedError(newNestedError(err, "a"), "b")
	assert.EqualValues(t, ne.Original(), err)
	m := make(map[string]interface{})
	assert.NoError(t, json.Unmarshal([]byte(ne.Error()), &m))
	assert.True(t, len(m) > 0)
}

func TestEvaluatorPanic(t *testing.T) {
	ev, err := NewEvaluator(`x eq 1`)
	assert.NoError(t, err)
	ev.testHookPanic = func() {
		panic("wait what")
	}
	ret, err := ev.Process(map[string]interface{}{"x": 1})
	assert.False(t, ret)
	assert.Error(t, err)
}
