// Package stack implements a stack
// See https://en.wikipedia.org/wiki/Stack_(abstract_data_type) for more details
package stack

// Stack represents a stack
// Zero value of Stack is empty stack
type Stack[T any] struct {
	array []T
}

// New returns an initialized stack
func New[T any]() *Stack[T] {
	return &Stack[T]{[]T{}}
}

// Push adds element
func (s *Stack[T]) Push(value T) {
	s.array = append(s.array, value)
}

// Pop removes the most recently added element
func (s *Stack[T]) Pop() T {
	v := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return v
}

// Peek returns element on tom of stack
func (s *Stack[T]) Peek() T {
	return s.array[len(s.array)-1]
}

// Size returns size of the stack
func (s *Stack[T]) Size() int {
	return len(s.array)
}
