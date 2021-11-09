package splaytree

// Node represents node of a splay tree
type Node[T any] struct {
	parent *Node[T]
	right  *Node[T]
	left   *Node[T]
	value  T
	size   int
}

// isRight returns true if node is right child of his parent, false otherwise
func (n *Node[T]) isRight() bool {
	if n == nil || n.parent == nil {
		return false
	}
	parent := n.parent
	return parent.right == n
}

// isRight returns true if node is left child of his parent, false otherwise
func (n *Node[T]) isLeft() bool {
	if n == nil || n.parent == nil {
		return false
	}
	parent := n.parent
	return parent.left == n
}

// isRoot returns true if given node is root
func (n *Node[T]) isRoot() bool {
	return n.parent == nil
}

// setRight sets c as right child of p
func setRight[T any](p, c *Node[T]) {
	if p == nil {
		return
	}
	p.right = c
	if c != nil {
		c.parent = p
	}
}

// setRight sets c as left child of p
func setLeft[T any](p, c *Node[T]) {
	if p == nil {
		return
	}
	p.left = c
	if c != nil {
		c.parent = p
	}
}

// find returns node containing given value
// or last node reached
func (n *Node[T]) find(value T, comp func(T, T) int) *Node[T] {
	if n == nil {
		return nil
	}
	if comp(n.value, value) == 0 {
		n.splay()
		return n
	}
	if comp(value, n.value) < 0 {
		if n.left == nil {
			n.splay()
			return n
		}
		return n.left.find(value, comp)
	} else {
		if n.right == nil {
			n.splay()
			return n
		}
		return n.right.find(value, comp)
	}
}

// splay uses splay
func (n *Node[T]) splay() {
	for !n.isRoot() && !n.parent.isRoot() {
		if (n.isLeft() && n.parent.isLeft()) || (n.isRight() && n.parent.isRight()) {
			n.zigZig()
		} else {
			n.zigZag()
		}
	}
	if !n.isRoot() {
		n.zig()
	}
}

// zig uses left or right rotation
func (n *Node[T]) zig() {
	grandParent := n.parent.parent
	isLeft := n.parent.isLeft()
	if n.isLeft() {
		right := n.right
		setRight(n, n.parent)
		setLeft(n.right, right)
		n.right.recalculateSize()
	} else {
		left := n.left
		setLeft(n, n.parent)
		setRight(n.left, left)
		n.left.recalculateSize()
	}
	n.recalculateSize()
	n.parent = grandParent
	if isLeft {
		setLeft(grandParent, n)
	} else {
		setRight(grandParent, n)
	}
}

// zigZig uses zig zig rotation
func (n *Node[T]) zigZig() {
	n.parent.zig()
	n.zig()
}

// zigZag uses zig zag rotation
func (n *Node[T]) zigZag() {
	n.zig()
	n.zig()
}

// max returns node with max element
func (n *Node[T]) max() *Node[T] {
	for n.right != nil {
		n = n.right
	}
	return n
}

// max returns node with max element
func (n *Node[T]) min() *Node[T] {
	for n.left != nil {
		n = n.left
	}
	return n
}

// split splits given node by given key into two nodes
func split[T any](n *Node[T], key T, comp func(T, T) int) (*Node[T], *Node[T]) {
	if n == nil {
		return nil, nil
	}
	nn := n.find(key, comp)
	if comp(key, n.value) <= 0 {
		left := nn.left
		if left != nil {
			left.parent = nil
		}
		nn.left = nil
		nn.recalculateSize()
		return left, n
	} else {
		right := nn.right
		if right != nil {
			right.parent = nil
		}
		nn.right = nil
		nn.recalculateSize()
		return nn, right
	}
}

// merge merges two nodes, all elements of left node should be less than any of right elements
func merge[T any](left *Node[T], right *Node[T]) *Node[T] {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	max := left.max()
	max.splay()
	setRight(max, right)
	max.recalculateSize()
	return max
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

// recalculateSize recalculates size of given node
func (n *Node[T]) kth(i int) *Node[T] {
	if n.size-1 == i {
		return n
	}
	if n.size > i {
		return n.left.kth(i)
	}
	return n.right.kth(i - n.left.size - 1)
}

func (n *Node[T]) next() *Node[T] {
	if n.right != nil {
		return n.right.min()
	}
	for n.isRight() {
		n = n.parent
	}
	n = n.parent
	return n
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
