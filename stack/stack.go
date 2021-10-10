package stack

import "github.com/Tv0ridobro/data-structure/list"

type Stack[T any] struct {
	list list.List[T]
}

func (s *Stack[T]) Push(value T) {
	s.list.PushBack(value)
}

func (s *Stack[T]) Pop() T {
	return s.list.PopBack()
}

func (s *Stack[T]) Size() int {
	return s.list.Len()
}
