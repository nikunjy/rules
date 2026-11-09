package parser

import (
	"encoding/json"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedError(t *testing.T) {
	t.Parallel()

	base := errors.New("some random error")
	tests := []struct {
		name     string
		err      *NestedError
		original error
		want     string
		wantVals ErrVals
	}{
		{
			name:     "nested",
			err:      newNestedError(newNestedError(base, "a"), "b"),
			original: base,
			want:     `{"err":"{\"err\":\"some random error\",\"msg\":\"a\"}","msg":"b"}`,
		},
		{
			name:     "vals",
			err:      newNestedError(base, "m").Set(ErrVals{"attr_path": "x.y", "rule_operand": []int{1, 2}}),
			original: base,
			want:     `{"attr_path":"x.y","err":"some random error","msg":"m","rule_operand":[1,2]}`,
			wantVals: ErrVals{"attr_path": "x.y", "rule_operand": []int{1, 2}},
		},
		{
			name:     "vals named err and msg are overridden",
			err:      newNestedError(base, "m").Set(ErrVals{"err": "e", "msg": "x"}),
			original: base,
			want:     `{"err":"some random error","msg":"m"}`,
			wantVals: ErrVals{"err": "e", "msg": "x"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.EqualValues(t, tt.original, tt.err.Original())
			m := make(map[string]interface{})
			assert.NoError(t, json.Unmarshal([]byte(tt.err.Error()), &m))
			// concurrent Error calls must not race (run with -race) and
			// must not write to Vals
			var wg sync.WaitGroup
			for g := 0; g < 4; g++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for i := 0; i < 100; i++ {
						assert.Equal(t, tt.want, tt.err.Error())
					}
				}()
			}
			wg.Wait()
			assert.Equal(t, tt.wantVals, tt.err.Vals)
		})
	}
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
