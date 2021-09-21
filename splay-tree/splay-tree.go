package splay_tree

type SplayTree[T Ordered] struct {
	root *Node[T]
}

func New[T Ordered]() *SplayTree[T] {
	return &SplayTree[T]{}
}

//Insert...
func (s *SplayTree[T]) Insert(value T) {
	n := &Node[T]{
		value: value,
		size:  1,
	}
	if s.root == nil {
		s.root = n
		return
	}
	nn := s.root.find(value)
	l, r := split(nn, value)
	setLeft(n, l)
	setRight(n, r)
	s.root = n
	s.root.recalculateSize()
}

//Contains
func (s *SplayTree[T]) Contains(value T) bool {
	if s.root == nil {
		return false
	}
	s.root = s.root.find(value)
	return s.root.value == value
}

//Size...
func (s *SplayTree[T]) Size() int {
	if s.root == nil {
		return 0
	}
	return s.root.size
}

func (s *SplayTree[T]) Remove(value T) {
	if s.root == nil {
		return
	}
	s.root = s.root.find(value)
	if s.root.value != value {
		return
	}
	l, r := s.root.left, s.root.right
	s.root.left, s.root.right = nil, nil
	if l != nil {
		l.parent = nil
	}
	if r != nil {
		r.parent = nil
	}
	s.root = merge(l, r)
	s.root.recalculateSize()
}

func (s *SplayTree[T]) GetAll() []T {
	if s.root == nil {
		return nil
	}
	d := make([]T, s.Size())
	s.root.getAll(d)
	return d
}
