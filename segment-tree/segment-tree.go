package segment_tree

import (
	"github.com/Tv0ridobro/data-structure/util"
)

//https://en.wikipedia.org/wiki/Segment_tree
type SegmentTree[T any] struct {
	op       func(T, T) T
	neutral  T
	elements []T
}

//New returns new SegmentTree
//T should be monoid
//op(op(a, b), c) = op(a, op(b, c)) and op(neutral, a) = op(a, neutral) = a)
func New[T any](elements []T, op func(T, T) T, neutral T, ) *SegmentTree[T] {
	d := util.NearestPowerOf2(len(elements))
	c := make([]T, d*2)
	copied := copy(c[d:], elements)
	for i := d + copied; i < 2*d; i++ {
		c[i] = neutral
	}
	st := SegmentTree[T]{
		op:       op,
		neutral:  neutral,
		elements: c,
	}
	st.build(1, 0, d-1)
	return &st
}

func (s *SegmentTree[T]) build(i, left, right int) {
	if right == left {
		return
	}
	middle := (left + right) / 2
	s.build(2*i, left, middle)
	s.build(2*i+1, middle+1, right)
	s.elements[i] = s.op(s.elements[2*i], s.elements[2*i+1])
}

func (s *SegmentTree[T]) Query(l, r int) T {
	return s.query(1, 0, len(s.elements)/2-1, l, r)
}

func (s *SegmentTree[T]) query(i, left, right, l, r int) T {
	if l > r {
		return s.neutral
	}
	if l == left && r == right {
		return s.elements[i]
	}
	middle := (left + right) / 2
	return s.op(
		s.query(i*2, left, middle, l, util.Min(r, middle)),
		s.query(i*2+1, middle+1, right, util.Max(l, middle+1), r),
	)
}

//Modify allows you to modify element at given index
func (s *SegmentTree[T]) Modify(i int, data T) {
	s.modify(1, i, 0, len(s.elements)/2-1, data)
	return
}

//Modify allows you to modify element at given index
func (s *SegmentTree[T]) modify(pos, i, left, right int, data T) {
	if right == left {
		s.elements[pos] = data
		return
	}
	middle := (left + right) / 2
	if i <= middle {
		s.modify(pos*2, i, left, middle, data)
	} else {
		s.modify(pos*2+1, i, middle+1, right, data)
	}
	s.elements[pos] = s.op(s.elements[pos*2], s.elements[pos*2+1])
	return
}
