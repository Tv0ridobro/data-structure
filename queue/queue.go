package queue

import "github.com/Tv0ridobro/data-structure/list"

type Queue[T any] struct {
	list list.List[T]
}

func (q *Queue[T]) Enqueue(value T) {
	q.list.PushBack(value)
}

func (q *Queue[T]) Dequeue() T {
	return q.list.PopFront()
}

func (q *Queue[T]) Size() int {
	return q.list.Len()
}
