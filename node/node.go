package node

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func New[T any](value T) *Node[T] {
	return &Node[T]{
		Value: value,
		Next:  nil,
	}
}
