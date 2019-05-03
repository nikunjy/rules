package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
