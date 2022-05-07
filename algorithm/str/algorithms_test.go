package str

import (
	"testing"

	"github.com/Tv0ridobro/data-structure/slices"
)

func TestZfunction(t *testing.T) {
	t.Parallel()
	tests := []struct {
		line      string
		zfunction []int
	}{
		{"aaaAAA", []int{0, 2, 1, 0, 0, 0}},
		{"abacaba", []int{0, 0, 1, 0, 3, 0, 1}},
		{"babaabbb", []int{0, 0, 2, 0, 0, 1, 1, 1}},
	}
	for i := range tests {
		ans := ZFunction(tests[i].line)
		if !slices.Equal(ans, tests[i].zfunction) {
			t.Errorf("test %d wrong answer for line %s", i+1, tests[i].line)
			t.Errorf("%v %v", ans, tests[i].zfunction)
		}
	}
}

func TestPrefixFunction(t *testing.T) {
	t.Parallel()
	tests := []struct {
		line   string
		prefix []int
	}{
		{"abcdabcabcdabcdab", []int{0, 0, 0, 0, 1, 2, 3, 1, 2, 3, 4, 5, 6, 7, 4, 5, 6}},
		{"abcdabscabcdabia", []int{0, 0, 0, 0, 1, 2, 0, 0, 1, 2, 3, 4, 5, 6, 0, 1}},
	}
	for i := range tests {
		ans := PrefixFunction(tests[i].line)
		if !slices.Equal(ans, tests[i].prefix) {
			t.Errorf("test %d wrong answer for line %s", i+1, tests[i].line)
			t.Errorf("%v %v", ans, tests[i].prefix)
		}
	}
}
