package rules

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvaluateBasic(t *testing.T) {
	res, err := Evaluate(`x.c.d eq "abc"`, map[string]interface{}{
		"x": map[string]interface{}{
			"c": map[string]interface{}{
				"d": "abc",
			},
		},
	})
	require.NoError(t, err)
	require.True(t, res)

}

func TestSum(t *testing.T) {
	res, err := Evaluate(`SUM (x,y,z) eq 10`, map[string]interface{}{
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
	t.Run("when subtract  (5-9) the result should be -7", func(t *testing.T) {
		res, err := Evaluate(`SUBTRACT (x.c,y) EQ -7`, map[string]interface{}{
			"x": map[string]interface{}{
				"c": 2,
			},
			"y": 9,
		})
		require.NoError(t, err)
		require.True(t, res)
	})
	t.Run("when subtract  (5-9) the result should be 7", func(t *testing.T) {
		res, err := Evaluate(`SUBTRACT (y,x.c.d) EQ 7`, map[string]interface{}{
			"x": map[string]interface{}{
				"c": map[string]interface{}{
					"d": 2,
				},
			},
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
	t.Run("when subtract  (5-4) the result not should be 1", func(t *testing.T) {
		res, err := Evaluate(`SUBTRACT (x,y) EQ -1`, map[string]interface{}{
			"x": 5,
			"y": 4,
		})
		require.NoError(t, err)
		require.False(t, res)
	})
}

func TestSubtractUnSuccessfully(t *testing.T) {
	t.Run("when parameters are not enough should return an error", func(t *testing.T) {
		res, err := Evaluate(`SUBTRACT (y.c) EQ 1`, map[string]interface{}{
			"y": map[string]interface{}{
				"c": 4,
			},
		})
		require.Error(t, err)
		require.False(t, res)
	})
}

func TestDiv(t *testing.T) {
	t.Run("when division works successfully", func(t *testing.T) {
		res, err := Evaluate(`DIV (x,y) EQ 10`, map[string]interface{}{
			"x": 100,
			"y": 10,
		})
		require.NoError(t, err)
		require.True(t, res)
	})
	t.Run("when dividing by 0 should return 0", func(t *testing.T) {
		res, err := Evaluate(`DIV (x,y) EQ 0`, map[string]interface{}{
			"x": 100,
			"y": 0,
		})
		require.NoError(t, err)
		require.True(t, res)
	})
}
