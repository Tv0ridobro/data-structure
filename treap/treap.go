// Package treap implements a treap
// See https://en.wikipedia.org/wiki/Treap for more details
package treap

import (
	"math/rand"
	"constraints"
)


// Treap represents a treap
// Zero value of Treap is invalid treap, should be used only with New() or NewWithSource()
type Treap[T constraints.Ordered] struct {
	rand *rand.Rand
	root *Node[T]
}

// New returns an initialized treap
// rand.Rand is used with zero seed
// For custom rand.Rand use NewWithSource
func New[T constraints.Ordered]() *Treap[T] {
	return &Treap[T]{
		rand: rand.New(rand.NewSource(0)),
	}
}

// NewWithSource returns an initialized treap with given source
func NewWithSource[T constraints.Ordered](s rand.Source) *Treap[T] {
	return &Treap[T]{
		rand: rand.New(s),
	}
}

// Insert inserts value in a tree
func (t *Treap[T]) Insert(value T) {
	n := &Node[T]{
		priority: t.rand.Int(),
		value:    value,
		size:     1,
	}
	if t.root == nil {
		t.root = n
		return
	}
	left, right := split(t.root, n.value)
	left1 := merge(left, n)
	right1 := merge(left1, right)
	t.root = right1
}

// Remove removes value from tree
// returns true if tree contained given value, false otherwise
func (t *Treap[T]) Remove(value T) bool {
	if t.root == nil {
		return false
	}
	oldSize := t.root.size
	left, right := split(t.root, value)
	if right != nil {
		right = tryRemoveMin(right, value)
	}
	t.root = merge(left, right)
	return oldSize != t.Size()
}

// Contains returns true if tree contains given value, false otherwise
func (t *Treap[T]) Contains(value T) bool {
	if t.root == nil {
		return false
	}
	return t.root.contains(value)
}

// Size returns size of the tree
func (t *Treap[T]) Size() int {
	if t.root == nil {
		return 0
	}
	return t.root.size
}

// GetAll returns all elements from tree
// returned slice is sorted
func (t *Treap[T]) GetAll() []T {
	if t.root == nil {
		return nil
	}
	d := make([]T, t.Size())
	t.root.getAll(d)
	return d
}
