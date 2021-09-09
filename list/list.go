package list

import (
	"fmt"
	"github.com/Tv0ridobro/dataStructure/node"
)

// List is simple linked list
type List[T any] struct {
	head *node.Node[T]
	len  int
}

//New returns new list
func New[T any]() *List[T] {
	return &List[T]{}
}

//PushFront adds data to the beginning of the list
func (l *List[T]) PushFront(value T) {
	l.len++
	nn := node.New[T](value)
	if l.head == nil {
		l.head = nn
		return
	}
	nn.Next = l.head
	l.head = nn
}

//PushBack adds data to the end of the list
func (l *List[T]) PushBack(value T) {
	l.len++
	newNode := node.New(value)
	if l.head == nil {
		l.head = newNode
		return
	}
	it := l.head
	for it.Next != nil {
		it = it.Next
	}
	it.Next = newNode
}

//Len returns length of the list
func (l *List[T]) Len() int {
	return l.len
}

//GetAll returns all data from list from beginning to end
func (l *List[T]) GetAll() []T {
	data := make([]T, 0, l.len)
	it := l.head
	for it != nil {
		data = append(data, it.Value)
		it = it.Next
	}
	return data
}

//PopBack removes data from the end of the list
//panic if list is empty
func (l *List[T]) PopBack() T {
	if l.len == 0 {
		panic("empty list")
	}
	l.len--
	prev := l.head
	it := prev.Next
	if it == nil {
		l.head = nil
		return prev.Value
	}
	for it.Next != nil {
		prev = it
		it = it.Next
	}
	prev.Next = nil
	return it.Value
}

//PopFront removes data from the beginning of the list
//panic if list is empty
func (l *List[T]) PopFront() T {
	if l.len == 0 {
		panic("empty list")
	}
	l.len--
	data := l.head.Value
	l.head = l.head.Next
	return data
}

// ChangeAt changes value at given index
//panic if list is empty
func (l *List[T]) ChangeAt(i int, value T) {
	if i >= l.len {
		panic(fmt.Sprintf("index higher than len %d %d", i, l.len))
	}
	it := l.head
	counter := 0
	for counter != i {
		it = it.Next
		counter++
	}
	it.Value = value
}

//Peek returns element at the given index
//panic if index is less or equal to len of the list
func (l *List[T]) Peek(i int) T {
	if i >= l.len {
		panic(fmt.Sprintf("index higher than len %d %d", i, l.len))
	}
	it := l.head
	counter := 0
	for counter != i {
		it = it.Next
		counter++
	}
	return it.Value
}
