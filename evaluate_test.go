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

func TestSubtractSuccessfully(t *testing.T) {
	t.Run("when subtract  (5-9) the result should be -4", func(t *testing.T) {
		res, err := Evaluate(`SUBTRACT (x,y) EQ -4`, map[string]interface{}{
			"x": 5,
			"y": 9,
		})
		require.NoError(t, err)
		require.True(t, res)
	})
	t.Run("when subtract  (-5- (-9)) the result should be 4", func(t *testing.T) {
		res, err := Evaluate(`SUBTRACT (x,y) EQ 4`, map[string]interface{}{
			"x": -5,
			"y": -9,
		})
		require.NoError(t, err)
		require.True(t, res)
	})
	t.Run("when subtract  (5-4) the result should be -1", func(t *testing.T) {
		res, err := Evaluate(`SUBTRACT (x,y) EQ 0`, map[string]interface{}{
			"x": 5,
			"y": 4,
		})
		require.NoError(t, err)
		require.False(t, res)
	})
}

func TestSubtractUnSuccessfully(t *testing.T) {
	t.Run("when parameters are not enough should return an error", func(t *testing.T) {
		res, err := Evaluate(`SUBTRACT (x) EQ 0`, map[string]interface{}{
			"x": 5,
		})
		require.Error(t, err)
		require.False(t, res)
	})
}
