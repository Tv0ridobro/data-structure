// Package list implements a doubly linked list
// See https://en.wikipedia.org/wiki/Linked_list for more details
package list

import (
	"fmt"
)

// List represents a doubly linked list
// Zero value of List is empty list
type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

// New returns an initialized list
func New[T any]() *List[T] {
	return &List[T]{}
}

// PushFront adds value to the beginning of the list
func (l *List[T]) PushFront(value T) {
	l.len++
	nn := &Node[T]{
		value: value,
	}
	if l.head == nil {
		l.head = nn
		return
	}
	if l.tail == nil { // 1 element in list
		l.tail = l.head
		l.head = nn
		l.tail.prev = l.head
		l.head.next = l.tail
		return
	}
	nn.next = l.head
	l.head.prev = nn
	l.head = nn
}

// PushBack adds value to the end of the list
func (l *List[T]) PushBack(value T) {
	l.len++
	nn := &Node[T]{
		value: value,
	}
	if l.head == nil {
		l.head = nn
		return
	}
	if l.tail == nil { // 1 element in list
		l.tail = nn
		nn.prev = l.head
		l.head.next = l.tail
		return
	}
	l.tail.next = nn
	nn.prev = l.tail
	l.tail = nn
}

// Len returns length of the list
func (l *List[T]) Len() int {
	return l.len
}

// GetAll returns all elements from the list
func (l *List[T]) GetAll() []T {
	data := make([]T, 0, l.len)
	it := l.head
	for it != nil {
		data = append(data, it.value)
		it = it.next
	}
	return data
}

// PopBack removes last element of the list
// Panics if list is empty
func (l *List[T]) PopBack() T {
	if l.len == 0 {
		panic("empty list")
	}
	l.len--
	if l.tail == nil {
		v := l.head.value
		l.head = nil
		return v
	}
	v := l.tail.value
	l.tail.prev.next = nil
	l.tail = l.tail.prev
	if l.tail == l.head { // if 1 element in list
		l.tail = nil
	}
	return v
}

// PopFront removes first element of the list
// Panics if list is empty
func (l *List[T]) PopFront() T {
	if l.len == 0 {
		panic("empty list")
	}
	l.len--
	v := l.head.value
	l.head = l.head.next
	if l.head == nil {
		return v
	}
	l.head.prev = nil
	if l.head == l.tail {
		l.tail = nil
	}
	return v
}

// ChangeAt changes value at given index
// panic if list is empty
func (l *List[T]) ChangeAt(i int, value T) {
	l.Node(i).value = value
}

// Peek returns element at the given index
// Panics if index is less or equal to len of the list
func (l *List[T]) Peek(i int) T {
	return l.Node(i).value
}

// Node returns Node at the given index
// Panics if index is less or equal to len of the list
func (l *List[T]) Node(i int) *Node[T] {
	if i >= l.len {
		panic(fmt.Sprintf("index higher than len %d %d", i, l.len))
	}
	if i <= l.len/2 {
		it := l.head
		counter := 0
		for counter != i {
			it = it.next
			counter++
		}
		return it
	}
	it := l.tail
	counter := 0
	for counter != l.len-1-i {
		it = it.prev
		counter++
	}
	return it
}

// Merge merges two lists
func (l *List[T]) Merge(other *List[T]) {
	if other.head == nil {
		return
	}
	if l.head == nil {
		l.len = other.len
		l.head = other.head
		l.tail = other.tail
		return
	}
	l.len += other.len
	if l.tail == nil {
		l.head.next = other.head
		other.head.prev = l.head
	} else {
		l.tail.next = other.head
		other.head.prev = l.tail
	}
	if other.tail == nil {
		l.tail = other.head
	} else {
		l.tail = other.tail
	}
}

// Clear clears list
func (l *List[T]) Clear() {
	l.len = 0
	l.head = nil
	l.tail = nil
}

// Reverse reverses given list
func (l *List[T]) Reverse() {
	if l.head == nil || l.tail == nil {
		return
	}
	var prev *Node[T]
	it, f := l.head, l.head
	next := it.next
	for next != nil {
		it.next = prev
		it.prev = next
		prev, it, next = it, next, next.next
	}
	it.next = prev
	l.head = it
	l.tail = f
}

// InsertAfter adds new Node with given value after given Node
func (l *List[T]) InsertAfter(node *Node[T], value T) {
	if node == l.tail || l.len == 1 {
		l.PushBack(value)
		return
	}
	l.len++
	nn := &Node[T]{
		value: value,
	}
	next := node.next
	node.next = nn
	nn.next = next
	nn.prev = node
	next.prev = nn
}

// RemoveAfter removes next Node of given Node from List
func (l *List[T]) RemoveAfter(n *Node[T]) T {
	if n == l.tail || n.next == nil {
		panic("")
	}
	if n.next == l.tail {
		return l.PopBack()
	}
	l.len--
	next := n.next
	gnext := next.next
	gnext.prev = n
	n.next = gnext
	return next.value
}

// RemoveBefore removes previous Node of given Node from list
func (l *List[T]) RemoveBefore(n *Node[T]) T {
	if n == l.head {
		panic("")
	}
	if n.prev == l.head {
		return l.PopFront()
	}
	l.len--
	prev := n.prev
	gprev := prev.prev
	gprev.next = n
	n.prev = gprev
	return prev.value
}

// InsertBefore adds new Node with given value before given Node
func (l *List[T]) InsertBefore(n *Node[T], value T) {
	if n == l.head {
		l.PushFront(value)
		return
	}
	l.len--
	nn := &Node[T]{value: value}
	prev := n.prev
	prev.next = nn
	nn.prev = prev
	nn.next = n
	n.prev = nn
}
