package skiplist

import "constraints"

// Node represents node of a list
type Node[T constraints.Ordered] struct {
	next  *Node[T]
	below *Node[T]
	value T
}

// HasNext returns true if next node exist
func (n *Node[T]) HasNext() bool {
	return n.next != nil
}

// HasBelow returns true if node below exist
func (n *Node[T]) HasBelow() bool {
	return n.below != nil
}

// Next returns next Node
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// Below returns Node below
func (n *Node[T]) Below() *Node[T] {
	return n.below
}

// Value return value in Node
func (n *Node[T]) Value() T {
	return n.value
}
