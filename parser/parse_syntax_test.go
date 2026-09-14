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
			ev, err := NewEvaluator(tt.rule)
			if tt.wantErr {
				assert.Error(t, err, tt.rule)
				assert.Nil(t, ev)
				return
			}
			assert.NoError(t, err, tt.rule)
			assert.NotNil(t, ev)
		})
	}
}

func TestNewEvaluatorRejectsInvalidSyntax(t *testing.T) {
	tests := []struct {
		name string
		rule string
	}{
		{name: "missing closing parenthesis", rule: `((env eq "pro") and (company eq "my-company")`},
		{name: "extra operator", rule: `env eq "pro" eq "dev"`},
		{name: "no operator", rule: `invalid`},
		{name: "leftover tokens", rule: `x eq 1 leftover`},
		{name: "extra closing parenthesis", rule: `(env eq "pro"))`},
		{name: "empty rule", rule: ``},
		{name: "unclosed string", rule: `x eq "abc`},
		{name: "incomplete and", rule: `x eq 1 and`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev, err := NewEvaluator(tt.rule)
			assert.Error(t, err, tt.rule)
			assert.Nil(t, ev)
		})
	}
}

func TestNewEvaluatorWhitespacePolicy(t *testing.T) {
	tests := []struct {
		name  string
		rule  string
		input obj
		want  bool
	}{
		{
			name:  "leading space",
			rule:  ` env eq "pro"`,
			input: obj{"env": "pro"},
			want:  true,
		},
		{
			name:  "trailing CRLF",
			rule:  "env eq \"pro\"\r\n",
			input: obj{"env": "pro"},
			want:  true,
		},
		{
			name:  "CRLF only",
			rule:  "env eq \"pro\"\r",
			input: obj{"env": "pro"},
			want:  true,
		},
		{
			name:  "double space before and evaluates both sides",
			rule:  `x eq 1  and y eq 2`,
			input: obj{"x": 1, "y": 2},
			want:  true,
		},
		{
			name:  "double space before and is not a prefix of the left clause",
			rule:  `x eq 1  and y eq 2`,
			input: obj{"x": 1, "y": 99},
			want:  false,
		},
		{
			name:  "extra spaces around operator",
			rule:  `x eq  1`,
			input: obj{"x": 1},
			want:  true,
		},
		{
			name:  "spaces inside string are preserved",
			rule:  `x eq "a  b"`,
			input: obj{"x": "a  b"},
			want:  true,
		},
		{
			name:  "trailing spaces and newline",
			rule:  "env eq \"pro\"  \n",
			input: obj{"env": "pro"},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev, err := NewEvaluator(tt.rule)
			require.NoError(t, err, tt.rule)
			result, err := ev.Process(tt.input)
			require.NoError(t, err, tt.rule)
			assert.Equal(t, tt.want, result, tt.rule)
		})
	}
}

func TestNormalizeRuleInput(t *testing.T) {
	assert.Equal(t, `x eq 1 and y eq 2`, normalizeRuleInput("x eq 1  and y eq 2"))
	assert.Equal(t, `env eq "pro"`, normalizeRuleInput(" env eq \"pro\" \r\n"))
	assert.Equal(t, `x eq "a  b"`, normalizeRuleInput(`x eq "a  b"`))
}
