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
	size   int
}

func (n *Node[T]) isRight() bool {
	if n == nil || n.parent == nil {
		return false
	}
	parent := n.parent
	return parent.right == n
}

func (n *Node[T]) isLeft() bool {
	if n == nil || n.parent == nil {
		return false
	}
	parent := n.parent
	return parent.left == n
}

func (n *Node[T]) isRoot() bool {
	return n.parent == nil
}

func setRight[T Ordered](p, c *Node[T]) {
	if p == nil {
		return
	}
	p.right = c
	if c != nil {
		c.parent = p
	}
}

func setLeft[T Ordered](p, c *Node[T]) {
	if p == nil {
		return
	}
	p.left = c
	if c != nil {
		c.parent = p
	}
}

func (n *Node[T]) find(value T) *Node[T] {
	if n == nil {
		return nil
	}
	if n.value == value {
		n.splay()
		return n
	}
	if value < n.value {
		if n.left == nil {
			n.splay()
			return n
		}
		return n.left.find(value)
	} else {
		if n.right == nil {
			n.splay()
			return n
		}
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

func split[T Ordered](n *Node[T], key T) (*Node[T], *Node[T]) {
	if n == nil {
		return nil, nil
	}
	nn := n.find(key)
	if key <= n.value {
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

func merge[T Ordered](left *Node[T], right *Node[T]) *Node[T] {
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
