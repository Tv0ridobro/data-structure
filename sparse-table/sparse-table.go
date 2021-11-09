// Package sparsetable implements a sparse table
package sparsetable

import (
	"fmt"
	"github.com/Tv0ridobro/data-structure/math"
)

// SparseTable represents a sparse table
// Zero value of SparseTable is invalid sparse table, should be used only with New()
type SparseTable[T any] struct {
	op       func(T, T) T
	elements [][]T
}

// New returns an initialized sparse table
// op should be idempotent, commutative and associative
func New[T any](op func(T, T) T, elements []T) *SparseTable[T] {
	ts := make([][]T, len(elements))
	lg := math.Log2(len(ts))
	for i := 0; i < len(ts); i++ {
		ts[i] = make([]T, lg+1)
		ts[i][0] = elements[i]
	}
	for j := 1; j <= math.Log2(len(elements)); j++ {
		for i := 0; i < len(elements); i++ {
			if (i + (1 << (j - 1))) < len(elements) {
				ts[i][j] = op(ts[i][j-1], ts[i+(1<<(j-1))][j-1])
			}
		}
	}
	return &SparseTable[T]{
		op:       op,
		elements: ts,
	}
}

// Query returns result of operation on elements[l:r+1]
func (s *SparseTable[T]) Query(l, r int) T {
	if l > r || l < 0 || r >= len(s.elements) {
		panic(fmt.Sprintf("wrong indices %d %d", l, r))
	}
	log := math.Log2(r - l + 1)
	return s.op(s.elements[l][log], s.elements[r-(1<<log)+1][log])
}
