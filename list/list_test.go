package list

import (
	"log"
	"math/rand"
	"reflect"
	"testing"
)

func TestList_GetAll(t *testing.T) {
	const length = 100
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
			t.Fatalf("data[i] != list.Peek(i) %d %d", data[i], list.Peek(i))
		}
	}
	if !reflect.DeepEqual(data, list.GetAll()) {
		log.Fatal(data,list.GetAll())
	}
}
