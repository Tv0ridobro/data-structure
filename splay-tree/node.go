package splay_tree

type Ordered interface{
~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 |
~float64
}

type Node[T Ordered] struct {
	parent *Node[T]
	right  *Node[T]
	left   *Node[T]
	value  T
}

func (n *Node[T]) isRight() bool {
	parent := n.parent
	return parent.right == n
}

func (n *Node[T]) isLeft() bool {
	parent := n.parent
	return parent.left == n
}

func (n *Node[T]) isRoot() bool {
	return n.parent == nil
}

func setRight(p, c *Node[T]) {
	if p == nil || c == nil {
		return
	}
	p.right = c
	c.parent = p
}

func setLeft(p, c *Node[T]) {
	if p == nil || c == nil {
		return
	}
	p.left = c
	c.parent = p
}

func (n *Node[T]) find(value T) *Node[T] {
	if n == nil {
		return nil
	}
	if n.value == value {
		return n
	}
	if value < n.value {
		return n.left.find(value)
	} else {
		return n.right.find(value)
	}
}

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

func (n *Node[T]) zig() {
	grandParent := n.parent.parent
	if n.isLeft() {
		right := n.right
		setRight(n, n.parent)
		setLeft(n.right, right)
	} else {
		left := n.left
		setLeft(n, n.parent)
		setRight(n.left, left)
	}
	n.parent = grandParent
}

func (n *Node[T]) zigZig() {
	n.parent.zig()
	n.zig()
}

func (n *Node[T]) zigZag() {
	n.zig()
	n.zig()
}

func (n *Node[T]) max() *Node[T] {
	for n.right != nil {
		n = n.right
	}
	return n
}
