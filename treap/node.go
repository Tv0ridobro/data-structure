package treap

// Node represents node of a treap
type Node[T any] struct {
	priority int
	value    T
	left     *Node[T]
	right    *Node[T]
	size     int
}

// contains returns true if given node contains given value
// False otherwise
func (n *Node[T]) contains(value T, comp func(T, T) int) bool {
	if n == nil {
		return false
	}
	if comp(n.value, value) == 0 {
		return true
	}
	if comp(value, n.value) < 0 {
		return n.left.contains(value, comp)
	}
	return n.right.contains(value, comp)
}

// tryRemoveMin tries to remove minimal element in given node if this element is the same as given one
func tryRemoveMin[T any](n *Node[T], expected T, comp func(T, T) int) *Node[T] {
	if n == nil {
		return nil
	}
	if comp(n.value, expected) == 0 {
		n = merge(n.left, n.right)
		return n
	}
	n.left = tryRemoveMin(n.left, expected, comp)
	n.recalculateSize()
	return n
}

// merge merges two nodes, all elements of left node should be less than any of right elements
func merge[T any](left *Node[T], right *Node[T]) *Node[T] {
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
func split[T any](n *Node[T], key T, comp func(T, T) int) (*Node[T], *Node[T]) {
	if n == nil {
		return nil, nil
	}
	if comp(key, n.value) > 0 {
		left, right := split(n.right, key, comp)
		n.right = left
		n.recalculateSize()
		return n, right
	}
	left, right := split(n.left, key, comp)
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
