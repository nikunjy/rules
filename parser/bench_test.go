package parser

import (
	"fmt"
	"testing"
)

var conditions = []string{
	`x eq 1`,
	`y IN [1, 2, 3, 4, 5]`,
	`z in ["abcd", "efghijkl", "asdoihadoihasoihsaoasihasdoihdadoiahsdoihdasodha", "asdoisahdadoihasdoisdwoqeioheqcnac"]`,
	`(a > 123 or b <= 456)`,
	`(aa == 123123131.1232313131312312 or bb == 1231.1312313123123131231)`,
	`(jh gt "1.2.3" or kl le "1.4.56")`,
	`o like "^(https?://)?(www\.)?([a-zA-Z0-9-]{1,63}\.)*[a-zA-Z0-9][a-zA-Z0-9-]{0,61}\.[a-zA-Z]{2,}(:\d+)?(/.*)?$"`,
	`p like ".*世界\s.*"`,
	`q like "离职.*未关闭"`,
}

func BenchmarkLogicalExpression(b *testing.B) {
	rule := ""
	for i, condition := range conditions {
		if i == len(conditions)-1 {
			rule += (" " + condition)
			continue
		}
		rule += fmt.Sprintf("%s or ", condition)
	}
	input := map[string]interface{}{
		"x":  123,
		"y":  2345,
		"z":  "asdada",
		"a":  "1212121",
		"b":  123131,
		"aa": 123123.123131232113,
		"jh": "1.2.3",
		"o":  "https://subdomain.example.org",
		"p":  "Hello, 世界 means world in Chinese.",
		"q":  "员工离职后系统账未关闭",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ev, _ := NewEvaluator(rule)
		ev.Process(input)
	}
}

func BenchmarkLogicalExpressionBig(b *testing.B) {
	rule := ""
	for j := 0; j < 100; j++ {
		for i, condition := range conditions {
			if i == len(conditions)-1 {
				rule += fmt.Sprintf(" (((%s)))", condition)
				continue
			}
			rule += fmt.Sprintf("(((%s))) or ", condition)
		}
	}
	input := map[string]interface{}{
		"x":  123,
		"y":  2345,
		"z":  "asdada",
		"a":  "1212121",
		"b":  123131,
		"aa": 123123.123131232113,
		"jh": "1.2.3",
		"o":  "https://subdomain.example.org",
		"p":  "Hello, 世界 means world in Chinese.",
		"q":  "员工离职后系统账未关闭",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ev, _ := NewEvaluator(rule)
		ev.Process(input)
	}

}

func BenchmarkLogicalExpressionBigTreeCached(b *testing.B) {
	rule := ""
	for j := 0; j < 100; j++ {
		for i, condition := range conditions {
			if i == len(conditions)-1 {
				rule += fmt.Sprintf(" (((%s)))", condition)
				continue
			}
			rule += fmt.Sprintf("(((%s))) or ", condition)
		}
	}
	input := map[string]interface{}{
		"x":  123,
		"y":  2345,
		"z":  "asdada",
		"a":  "1212121",
		"b":  123131,
		"aa": 123123.123131232113,
		"jh": "1.2.3",
		"o":  "https://subdomain.example.org",
		"p":  "Hello, 世界 means world in Chinese.",
		"q":  "员工离职后系统账未关闭",
	}
	b.ResetTimer()
	ev, _ := NewEvaluator(rule)
	for i := 0; i < b.N; i++ {
		ev.Process(input)
		ev.Reset()
	}

}

func BenchmarkLogicalExpressionNested(b *testing.B) {
	rule := ""
	for i, condition := range conditions {
		if i == len(conditions)-1 {
			rule += fmt.Sprintf(" (((%s)))", condition)
			continue
		}
		rule += fmt.Sprintf("(((%s))) or ", condition)
	}
	input := map[string]interface{}{
		"x":  123,
		"y":  2345,
		"z":  "asdada",
		"a":  "1212121",
		"b":  123131,
		"aa": 123123.123131232113,
		"jh": "1.2.3",
		"o":  "https://subdomain.example.org",
		"p":  "Hello, 世界 means world in Chinese.",
		"q":  "员工离职后系统账未关闭",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ev, _ := NewEvaluator(rule)
		ev.Process(input)
	}
}
