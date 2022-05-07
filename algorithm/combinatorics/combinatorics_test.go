package combinatorics

import (
	"testing"

	"github.com/Tv0ridobro/data-structure/slices"
)

func TestNextPermutation(t *testing.T) {
	t.Parallel()
	tests := []struct {
		now  []int
		next []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}}, // no next permutation
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 6, 5}},
		{[]int{}, []int{}},
		{[]int{3, 9, 8, 7, 6, 5, 4, 2, 1}, []int{4, 1, 2, 3, 5, 6, 7, 8, 9}},
	}
	for i := range tests {
		if v := NextPermutation(tests[i].now); !slices.Equal(v, tests[i].next) {
			t.Errorf("wrong next perutation %v != %v", v, tests[i].next)
		}
	}
}
