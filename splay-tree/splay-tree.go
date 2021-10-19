// Package splay_tree implements a splay tree
// See https://en.wikipedia.org/wiki/Splay_tree for more details
package splay_tree

// SplayTree represents a splay tree
// Zero value of SplayTree is empty splay tree
type SplayTree[T Ordered] struct {
	root *Node[T]
}

// New returns an initialized splay tree
func New[T Ordered]() *SplayTree[T] {
	return &SplayTree[T]{}
}

// Insert inserts value in a tree
func (s *SplayTree[T]) Insert(value T) {
	n := &Node[T]{
		value: value,
		size:  1,
	}
	if s.root == nil {
		s.root = n
		return
	}
	nn := s.root.find(value)
	l, r := split(nn, value)
	setLeft(n, l)
	setRight(n, r)
	s.root = n
	s.root.recalculateSize()
}

// Contains returns true if tree contains given value, false otherwise
func (s *SplayTree[T]) Contains(value T) bool {
	if s.root == nil {
		return false
	}
	s.root = s.root.find(value)
	return s.root.value == value
}

// Size returns size of the tree
func (s *SplayTree[T]) Size() int {
	if s.root == nil {
		return 0
	}
	return s.root.size
}

// Remove removes value from tree
// returns true if tree contained given value, false otherwise
func (s *SplayTree[T]) Remove(value T) bool {
	if s.root == nil {
		return false
	}
	s.root = s.root.find(value)
	if s.root.value != value {
		return false
	}
	l, r := s.root.left, s.root.right
	s.root.left, s.root.right = nil, nil
	if l != nil {
		l.parent = nil
	}
	if r != nil {
		r.parent = nil
	}
	s.root = merge(l, r)
	return true
}

// Kth returns kth greatest element
func (s *SplayTree[T]) Kth(i int) T {
	if i > s.Size() {
		panic("i is greater than size")
	}
	n := s.root.kth(i)
	return n.value
}

//Sub returns elements [l, r) in ascending order
func (s *SplayTree[T]) Sub(l, r int) []T {
	sl := make([]T, r-l)
	n := s.root.kth(l)
	sl[0] = n.value
	for i := 1; i < r-l; i++ {
		n = n.next()
		sl[i] = n.value
	}
	return sl
}

// GetAll returns all elements from tree
// returned slice is sorted
func (s *SplayTree[T]) GetAll() []T {
	if s.root == nil {
		return nil
	}
	d := make([]T, s.Size())
	s.root.getAll(d)
	return d
}
