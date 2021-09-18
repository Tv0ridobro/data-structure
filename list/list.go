package list

import (
	"fmt"
)

//List is simple linked list
type List[T any] struct {
	head *Node[T]
	len  int
}

//New returns new list
func New[T any]() *List[T] {
	return &List[T]{}
}

//PushFront adds data to the beginning of the list
func (l *List[T]) PushFront(value T) {
	l.len++
	nn := &Node[T]{
		value: value,
		next:  nil,
	}
	if l.head == nil {
		l.head = nn
		return
	}
	nn.next = l.head
	l.head = nn
}

//PushBack adds data to the end of the list
func (l *List[T]) PushBack(value T) {
	l.len++
	newNode := &Node[T]{
		value: value,
		next:  nil,
	}
	if l.head == nil {
		l.head = newNode
		return
	}
	it := l.head
	for it.next != nil {
		it = it.next
	}
	it.next = newNode
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
		data = append(data, it.value)
		it = it.next
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
	it := prev.next
	if it == nil {
		l.head = nil
		return prev.value
	}
	for it.next != nil {
		prev = it
		it = it.next
	}
	prev.next = nil
	return it.value
}

//PopFront removes data from the beginning of the list
//panic if list is empty
func (l *List[T]) PopFront() T {
	if l.len == 0 {
		panic("empty list")
	}
	l.len--
	data := l.head.value
	l.head = l.head.next
	return data
}

//ChangeAt changes value at given index
//panic if list is empty
func (l *List[T]) ChangeAt(i int, value T) {
	if i >= l.len {
		panic(fmt.Sprintf("index higher than len %d %d", i, l.len))
	}
	it := l.head
	counter := 0
	for counter != i {
		it = it.next
		counter++
	}
	it.value = value
}

//Peek returns element at the given index
//panic if index is less or equal to len of the list
func (l *List[T]) Peek(i int) T {
	return l.Node(i).value
}

//Node returns Node at the given index
//panic if index is less or equal to len of the list
func (l *List[T]) Node(i int) *Node[T] {
	if i >= l.len {
		panic(fmt.Sprintf("index higher than len %d %d", i, l.len))
	}
	it := l.head
	counter := 0
	for counter != i {
		it = it.next
		counter++
	}
	return it
}

//InsertAfter adds value after given node
func (l *List[T]) InsertAfter(node *Node[T], value T) {
	newNode := &Node[T]{
		value: value,
		next:  nil,
	}
	next := node.next
	node.next = newNode
	newNode.next = next
	l.len++
}

//Merge adds given list at the end of initial
//don't use second list after merging
func (l *List[T]) Merge(other *List[T]) {
	if l.head == nil {
		l.len = other.len
		l.head = other.head
		return
	}
	it := l.head
	for it.next != nil {
		it = it.next
	}
	it.next = other.head
	l.len += other.len
}

//Clear set len 0 and head to nil
func (l *List[T]) Clear() {
	l.len = 0
	l.head = nil
}

//Reverse reverse given list
func (l *List[T]) Reverse() {
	if l.head == nil {
		return
	}
	var prev *Node[T]
	it := l.head
	next := it.next
	for next != nil {
		it.next = prev
		prev, it, next = it, next, next.next
	}
	it.next = prev
	l.head = it
}

//RemoveAfter remove next Node of given Node from List
func (l *List[T]) RemoveAfter(n *Node[T]) {
	if n.next == nil {
		return
	}
	n.next = n.next.next
	l.len--
}
