package list

type Node[T any] struct {
	value T
	next  *Node[T]
}

//Next return next Node
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

//Value return value in Node
func (n *Node[T]) Value() T {
	return n.value
}

//Replace old value to new one and return old one
func (n *Node[T]) Replace(value T) T {
	old := n.value
	n.value = value
	return old
}

//HasNext return true if next node exists
func (n *Node[T]) HasNext() bool {
	return n.next != nil
}
