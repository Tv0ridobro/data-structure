package list

type Node[T any] struct {
	Value T
	Next  *Node[T]
}
