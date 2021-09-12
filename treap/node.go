package treap

type Ordered interface{
~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 |
~float64
}
type Node[T Ordered] struct {
	Priority int
	Value    T
	Left     *Node[T]
	Right    *Node[T]
	size     int
}

func (n *Node[T]) contains(value T) bool {
	if n == nil {
		return false
	}
	if n.Value == value {
		return true
	}
	if value < n.Value {
		return n.Left.contains(value)
	}
	return n.Right.contains(value)
}

func tryRemoveMin[T Ordered](n *Node[T], expected T) *Node[T] {
	if n == nil {
		return nil
	}
	if n.Value == expected {
		n = merge(n.Left, n.Right)
		return n
	}
	n.Left = tryRemoveMin(n.Left, expected)
	n.recalculateSize()
	return n
}

func tryRemoveMax[T Ordered](n *Node[T], expected T) *Node[T] {
	if n == nil {
		return nil
	}
	if n.Value == expected {
		n = merge(n.Left, n.Right)
		return n
	}
	n.Right = tryRemoveMax(n.Right, expected)
	n.recalculateSize()
	return n
}

func merge[T Ordered](left *Node[T], right *Node[T]) *Node[T] {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.Priority < right.Priority {
		right.Left = merge(left, right.Left)
		right.recalculateSize()
		return right
	} else {
		left.Right = merge(left.Right, right)
		left.recalculateSize()
		return left
	}
}

func split[T Ordered](n *Node[T], key T) (*Node[T], *Node[T]) {
	if n == nil {
		return nil, nil
	}
	if key > n.Value {
		left, right := split(n.Right, key)
		n.Right = left
		n.recalculateSize()
		return n, right
	}
	left, right := split(n.Left, key)
	n.Left = right
	n.recalculateSize()
	return left, n
}

func (n *Node[T]) recalculateSize() {
	if n == nil {
		return
	}
	n.size = 0
	if n.Left != nil {
		n.size += n.Left.size
	}
	if n.Right != nil {
		n.size += n.Right.size
	}
	n.size += 1
	return
}
