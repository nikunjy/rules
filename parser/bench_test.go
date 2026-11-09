package parser

import (
	"strings"
	"testing"
)

var conditions = []string{
	`x eq 1`,
	`y IN [1, 2, 3, 4, 5]`,
	`z in ["abcd", "efghijkl", "asdoihadoihasoihsaoasihasdoihdadoiahsdoihdasodha", "asdoisahdadoihasdoisdwoqeioheqcnac"]`,
	`(a > 123 or b <= 456)`,
	`(aa == 123123131.1232313131312312 or bb == 1231.1312313123123131231)`,
	`(jh gt "1.2.3" or kl le "1.4.56")`,
}

// logicalRule joins conditions with "or", wrapping each one in parentheses
// depth times.
func logicalRule(conds []string, depth int) string {
	parts := make([]string, len(conds))
	for i, c := range conds {
		parts[i] = strings.Repeat("(", depth) + c + strings.Repeat(")", depth)
	}
	return strings.Join(parts, " or ")
}

// bigRule repeats conditions 100 times, each wrapped in three parentheses.
// Every repetition must be joined with "or": since rules are validated
// strictly, a missing join makes NewEvaluator reject the rule.
func bigRule() string {
	conds := make([]string, 0, 100*len(conditions))
	for j := 0; j < 100; j++ {
		conds = append(conds, conditions...)
	}
	return logicalRule(conds, 3)
}

func newBenchEvaluator(b *testing.B, rule string) *Evaluator {
	b.Helper()
	ev, err := NewEvaluator(rule)
	if err != nil {
		b.Fatal(err)
	}
	return ev
}

// checkBenchProcess fails the benchmark unless rule evaluates to want on
// input with no error and no debug error, so an input that does not match
// the rule cannot silently turn a benchmark into an error-path benchmark.
// It returns the evaluator, which has now had its first Process call.
func checkBenchProcess(b *testing.B, rule string, input map[string]interface{}, want bool) *Evaluator {
	b.Helper()
	ev := newBenchEvaluator(b, rule)
	got, err := ev.Process(input)
	if err != nil || got != want || ev.LastDebugErr() != nil {
		b.Fatalf("got (%v, %v, debug %v), want (%v, <nil>, debug <nil>)", got, err, ev.LastDebugErr(), want)
	}
	return ev
}

func benchNewAndProcess(b *testing.B, rule string) {
	b.Helper()
	checkBenchProcess(b, rule, benchMixedInput, false)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ev := newBenchEvaluator(b, rule)
		ev.Process(benchMixedInput)
	}
}

func BenchmarkLogicalExpression(b *testing.B) {
	benchNewAndProcess(b, logicalRule(conditions, 0))
}

func BenchmarkLogicalExpressionBig(b *testing.B) {
	benchNewAndProcess(b, bigRule())
}

func BenchmarkLogicalExpressionBigTreeCached(b *testing.B) {
	ev := checkBenchProcess(b, bigRule(), benchMixedInput, false)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ev.Process(benchMixedInput)
		ev.Reset()
	}
}

func BenchmarkLogicalExpressionNested(b *testing.B) {
	benchNewAndProcess(b, logicalRule(conditions, 3))
}

// BenchmarkProcess measures evaluation of an already-parsed rule, which is the
// hot path for callers that cache the *Evaluator. The DebugErr cases measure
// the cold path that builds a debug error, on purpose.
func BenchmarkProcess(b *testing.B) {
	rules := []struct {
		name      string
		rule      string
		want      bool
		wantDebug bool
	}{
		{name: "IntEq", rule: `x eq 1`},
		{name: "IntIn", rule: `y IN [1, 2, 3, 4, 5]`},
		{name: "StringEq", rule: `z eq "asdada"`, want: true},
		{name: "StringIn", rule: `z in ["abcd", "efghijkl", "asdada", "asdoisahdadoihasdoisdwoqeioheqcnac"]`, want: true},
		{name: "FloatCmp", rule: `aa == 123123131.1232313131312312`},
		{name: "Version", rule: `jh gt 1.2.0`, want: true},
		{name: "Nested", rule: `u.plan eq "pro" and u.age gt 18`, want: true},
		{name: "Logical", rule: `(x eq 1 or y eq 2) and (z eq "nope" or b > 100)`},
		{name: "Mixed", rule: strings.Join(conditions, " or ")},
		{name: "DebugErr/TypeMismatch", rule: `z gt 123`, wantDebug: true},
		{name: "DebugErr/Missing", rule: `missing eq "x"`, wantDebug: true},
	}
	for _, tc := range rules {
		tc := tc
		b.Run(tc.name, func(b *testing.B) {
			ev := newBenchEvaluator(b, tc.rule)
			got, err := ev.Process(benchMixedInput)
			if err != nil || got != tc.want || (ev.LastDebugErr() != nil) != tc.wantDebug {
				b.Fatalf("got (%v, %v, debug %v), want (%v, <nil>, debug %v)", got, err, ev.LastDebugErr(), tc.want, tc.wantDebug)
			}
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ev.Process(benchMixedInput)
			}
		})
	}
}

// BenchmarkProcessParallel measures Process on one evaluator shared by
// concurrent callers; run it with -cpu 1,8 to see contention.
func BenchmarkProcessParallel(b *testing.B) {
	const rule = `x eq 1 and y in [1, 2, 3]`
	items := map[string]interface{}{"x": 1, "y": 2}
	ev := checkBenchProcess(b, rule, items, true)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ev.Process(items)
		}
	})
}

// benchMixedInput holds a correctly typed value for every attribute the
// benchmark rules read, so they evaluate without errors or debug errors.
// Each condition of conditions is false for it, so an or-chain of them is
// walked to the end:
//
//	x eq 1                       x is 123
//	y IN [1, 2, 3, 4, 5]          y is 2345
//	z in ["abcd", ...]            z is "asdada"
//	a > 123 or b <= 456           a is 100, b is 123131
//	aa == ... or bb == ...        aa is 123123.123131232113, bb is 1.5
//	jh gt "1.2.3" or kl le "1.4.56"
//	                              (string comparisons) jh is "1.2.3",
//	                              kl is "1.5" > "1.4.56"
var benchMixedInput = map[string]interface{}{
	"x":  123,
	"y":  2345,
	"z":  "asdada",
	"a":  100,
	"b":  123131,
	"aa": 123123.123131232113,
	"bb": 1.5,
	"jh": "1.2.3",
	"kl": "1.5",
	"u":  map[string]interface{}{"plan": "pro", "age": 31},
}

// BenchmarkNewEvaluator measures construction only (parse and validate).
func BenchmarkNewEvaluator(b *testing.B) {
	rule := strings.Join(conditions, " or ")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		newBenchEvaluator(b, rule)
	}
}

// BenchmarkEvaluate measures the one-shot Evaluate helper.
func BenchmarkEvaluate(b *testing.B) {
	rule := strings.Join(conditions, " or ")
	checkBenchProcess(b, rule, benchMixedInput, false)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Evaluate(rule, benchMixedInput)
	}
}

// BenchmarkNewEvaluatorFirstProcess measures a new evaluator plus its first
// Process call, i.e. what any one-use evaluator pays.
func BenchmarkNewEvaluatorFirstProcess(b *testing.B) {
	rule := strings.Join(conditions, " or ")
	checkBenchProcess(b, rule, benchMixedInput, false)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ev := newBenchEvaluator(b, rule)
		ev.Process(benchMixedInput)
	}
}

// TestBenchmarkVersionRule checks that the bare semver constant used by
// BenchmarkProcess/Version goes through the semver comparison.
func TestBenchmarkVersionRule(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		jh        string
		want      bool
		wantDebug bool
	}{
		{name: "semver ordering", jh: "1.10.0", want: true},
		{name: "not a semver", jh: "abc", wantDebug: true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ev, err := NewEvaluator(`jh gt 1.2.0`)
			if err != nil {
				t.Fatal(err)
			}
			got, err := ev.Process(map[string]interface{}{"jh": tt.jh})
			if err != nil || got != tt.want || (ev.LastDebugErr() != nil) != tt.wantDebug {
				t.Errorf("got (%v, %v, debug %v)", got, err, ev.LastDebugErr())
			}
		})
	}
}

// TestBenchmarkRules checks that the rules built for the logical expression
// benchmarks are accepted by the strict parser and evaluate benchMixedInput
// to false (every condition is walked) without an error or debug error.
func TestBenchmarkRules(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		rule string
	}{
		{name: "logical", rule: logicalRule(conditions, 0)},
		{name: "nested", rule: logicalRule(conditions, 3)},
		{name: "big", rule: bigRule()},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ev, err := NewEvaluator(tt.rule)
			if err != nil {
				t.Fatal(err)
			}
			got, err := ev.Process(benchMixedInput)
			if err != nil || got || ev.LastDebugErr() != nil {
				t.Fatalf("got (%v, %v, debug %v), want (false, <nil>, debug <nil>)", got, err, ev.LastDebugErr())
			}
		})
	}
}
