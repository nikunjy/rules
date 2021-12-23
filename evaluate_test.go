package rules

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvaluateBasic(t *testing.T) {
	res, err := Evaluate(`x eq "abc"`, map[string]interface{}{
		"x": "abc",
	})
	require.NoError(t, err)
	require.True(t, res)

	res, err = Evaluate(`x eq abc`, map[string]interface{}{
		"x": "abc",
	})
	require.Error(t, err)
	require.False(t, res)
}

func TestSum(t *testing.T) {
	res, err := Evaluate(`SUM ( x,y,z ) eq 10`, map[string]interface{}{
		"x": 5,
		"y": 5,
		"z": 0,
	})
	require.NoError(t, err)
	require.True(t, res)
}

func TestMul(t *testing.T) {
	res, err := Evaluate(`MLP (x,y,a) > 24`, map[string]interface{}{
		"x": 5,
		"y": 5,
		"a": 5,
	})
	require.NoError(t, err)
	require.True(t, res)
}
