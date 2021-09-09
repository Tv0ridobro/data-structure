package list

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestList_PeekPanic(t *testing.T) {
	list := New[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()
	list.Peek(0)
}

func TestList_PeekPanic2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()
	list := New[int]()
	for i := 0; i < 20; i++ {
		list.PushBack(i)
	}
	list.Peek(30)
}

func TestList_PopBack(t *testing.T) {
	list := New[int]()
	for i := 0; i < 20; i++ {
		list.PushBack(i)
	}
	for i := 19; i >= 0; i-- {
		if v := list.PopBack(); v != i {
			t.Errorf("Wrong value got %d expected %d", v, i)
		}
	}
}

func TestList_PopFront(t *testing.T) {
	list := New[int]()
	for i := 0; i < 20; i++ {
		list.PushFront(i)
	}
	for i := 19; i >= 0; i-- {
		if v := list.PopFront(); v != i {
			t.Errorf("Wrong value got %d expected %d", v, i)
		}
	}
}

func TestList_PopBackPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	list := New[int]()
	list.PopBack()
}

func TestList_PopFrontPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()
	list := New[int]()
	list.PopFront()
}

func TestList_Len(t *testing.T) {
	list := New[int]()
	if list.Len() != 0 {
		t.Error("wrong length size")
	}
	list.PushFront(7)
	if list.Len() != 1 {
		t.Error("wrong length size")
	}
	list.PopBack()
	if list.Len() != 0 {
		t.Error("wrong length size")
	}
	for i := 0; i < 15; i++ {
		list.PushBack(0)
	}
	if list.Len() != 15 {
		t.Error("wrong length size")
	}
}

func TestList_GetAllAndPeek(t *testing.T) {
	length := 100
	data := make([]int, 0)
	list := New[int]()
	for i := 0; i < length; i++ {
		value := rand.Int()
		command := rand.Intn(2)
		switch command {
		case 0:
			data = append(data, value)
			list.PushBack(value)
		case 1:
			data = append([]int{value}, data...)
			list.PushFront(value)
		}
	}
	for i := 0; i < length; i++ {
		if data[i] != list.Peek(i) {
			t.Errorf("data[i] != list.Peek(i) %d %d", data[i], list.Peek(i))
		}
	}
	if !reflect.DeepEqual(data, list.GetAll()) {
		t.Error(data, list.GetAll())
	}
}
