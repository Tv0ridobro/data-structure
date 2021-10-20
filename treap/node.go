package treap

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 |
		~float64
}

// Node represents node of a treap
type Node[T Ordered] struct {
	priority int
	value    T
	left     *Node[T]
	right    *Node[T]
	size     int
}

// contains returns true if given node contains given value
// False otherwise
func (n *Node[T]) contains(value T) bool {
	if n == nil {
		return false
	}
	if n.value == value {
		return true
	}
	if value < n.value {
		return n.left.contains(value)
	}
	return n.right.contains(value)
}

// tryRemoveMin tries to remove minimal element in given node if this element is the same as given one
func tryRemoveMin[T Ordered](n *Node[T], expected T) *Node[T] {
	if n == nil {
		return nil
	}
	if n.value == expected {
		n = merge(n.left, n.right)
		return n
	}
	n.left = tryRemoveMin(n.left, expected)
	n.recalculateSize()
	return n
}

// merge merges two nodes, all elements of left node should be less than any of right elements
func merge[T Ordered](left *Node[T], right *Node[T]) *Node[T] {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.priority < right.priority {
		right.left = merge(left, right.left)
		right.recalculateSize()
		return right
	} else {
		left.right = merge(left.right, right)
		left.recalculateSize()
		return left
	}
}

// split splits given node by given key into two nodes
func split[T Ordered](n *Node[T], key T) (*Node[T], *Node[T]) {
	if n == nil {
		return nil, nil
	}
	if key > n.value {
		left, right := split(n.right, key)
		n.right = left
		n.recalculateSize()
		return n, right
	}
	left, right := split(n.left, key)
	n.left = right
	n.recalculateSize()
	return left, n
}

// recalculateSize recalculates size of given node
func (n *Node[T]) recalculateSize() {
	if n == nil {
		return
	}
	n.size = 0
	if n.left != nil {
		n.size += n.left.size
	}
	if n.right != nil {
		n.size += n.right.size
	}
	n.size += 1
	return
}

// getAll returns all elements in node
// len of elements should be same as size of node
func (n *Node[T]) getAll(elements []T) {
	lSize := 0
	if n.left != nil {
		lSize = n.left.size
		n.left.getAll(elements[:lSize])
	}
	elements[lSize] = n.value
	if n.right != nil {
		n.right.getAll(elements[lSize+1:])
	}
}
