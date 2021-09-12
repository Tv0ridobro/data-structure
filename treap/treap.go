package treap

import "math/rand"

//https://en.wikipedia.org/wiki/Treap
type Treap[T Ordered] struct {
	rand *rand.Rand
	root *Node[T]
}

func New[T Ordered]() *Treap[T] {
	return &Treap[T]{
		rand: rand.New(rand.NewSource(0)),
		root: nil,
	}
}

func NewWithSource[T Ordered](s rand.Source) *Treap[T] {
	return &Treap[T]{
		rand: rand.New(s),
		root: nil,
	}
}

func (t *Treap[T]) Insert(value T) {
	n := &Node[T]{
		Priority: t.rand.Int(),
		Value:    value,
		Left:     nil,
		Right:    nil,
		size:     1,
	}
	if t.root == nil {
		t.root = n
		return
	}
	left, right := split(t.root, n.Value)
	left1 := merge(left, n)
	right1 := merge(left1, right)
	t.root = right1
}

func (t *Treap[T]) Remove(value T) {
	if t.root == nil {
		return
	}
	left, right := split(t.root, value)
	if right == nil {
		return
	}
	right = tryRemoveMin(right, value)
	t.root = merge(left, right)
}

func (t *Treap[T]) Contains(value T) bool {
	if t.root == nil {
		return false
	}
	return t.root.contains(value)
}

func (t *Treap[T]) Size() int {
	if t.root == nil {
		return 0
	}
	return t.root.size
}
