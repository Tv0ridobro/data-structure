package list

// Node represents node of a list
type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

// Next returns next Node
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// Prev returns previous Node
func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

// Value return value in Node
func (n *Node[T]) Value() T {
	return n.value
}

// Replace replaces old value to new one and return old one
func (n *Node[T]) Replace(value T) T {
	old := n.value
	n.value = value
	return old
}

// HasNext return true if next node exists
func (n *Node[T]) HasNext() bool {
	return n.next != nil
}

// HasPrev return true if previous node exists
func (n *Node[T]) HasPrev() bool {
	return n.prev != nil
}
