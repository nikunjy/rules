package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIssue41RuleValidator(t *testing.T) {
	tests := []struct {
		name    string
		rule    string
		wantErr bool
	}{
		{
			name:    "valid rule",
			rule:    `env eq "pro"`,
			wantErr: false,
		},
		{
			name:    "missing closing parenthesis",
			rule:    `((env eq "pro") and (company eq "my-company")`,
			wantErr: true,
		},
		{
			name:    "invalid extra operator",
			rule:    `env eq "pro" eq "dev"`,
			wantErr: true,
		},
		{
			name:    "no operator",
			rule:    `invalid`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := evalSyntax(tt.rule)
			if tt.wantErr {
				assert.Error(t, err, tt.rule)
				return
			}
			assert.NoError(t, err, tt.rule)
		})
	}
}

func TestNewEvaluatorRejectsInvalidSyntax(t *testing.T) {
	invalid := []string{
		`((env eq "pro") and (company eq "my-company")`,
		`env eq "pro" eq "dev"`,
		`invalid`,
		`x eq 1 leftover`,
		`(env eq "pro"))`,
		``,
	}
	for _, rule := range invalid {
		t.Run(rule, func(t *testing.T) {
			ev, err := NewEvaluator(rule)
			assert.Error(t, err, rule)
			assert.Nil(t, ev)
		})
	}
}

func TestNewEvaluatorAllowsTrailingWhitespace(t *testing.T) {
	ev, err := NewEvaluator("env eq \"pro\"  \n")
	require.NoError(t, err)
	result, err := ev.Process(obj{"env": "pro"})
	require.NoError(t, err)
	assert.True(t, result)
}

func evalSyntax(rule string) error {
	ev, err := NewEvaluator(rule)
	if err != nil {
		return err
	}
	_, err = ev.Process(map[string]interface{}{})
	return err
}
