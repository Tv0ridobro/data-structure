package queue

import "testing"

func TestQueue_Front(t *testing.T) {
	t.Parallel()
	q := New[int]()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	if v := q.Front(); v != 0 {
		t.Errorf("front wrong value %d != %d", v, 0)
	}
	q.Dequeue()
	if v := q.Front(); v != 1 {
		t.Errorf("front wrong value %d != %d", v, 1)
	}
	q.Dequeue()
	if v := q.Front(); v != 2 {
		t.Errorf("front wrong value %d != %d", v, 2)
	}
	q.Dequeue()
	if v := q.Front(); v != 3 {
		t.Errorf("front wrong value %d != %d", v, 3)
	}
	q.Dequeue()
	if v := q.Front(); v != 4 {
		t.Errorf("front wrong value %d != %d", v, 4)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	t.Parallel()
	q := New[int]()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	if v := q.Dequeue(); v != 0 {
		t.Errorf("front wrong value %d != %d", v, 0)
	}
	if v := q.Dequeue(); v != 1 {
		t.Errorf("front wrong value %d != %d", v, 1)
	}
	if v := q.Dequeue(); v != 2 {
		t.Errorf("front wrong value %d != %d", v, 2)
	}
	if v := q.Dequeue(); v != 3 {
		t.Errorf("front wrong value %d != %d", v, 3)
	}
	if v := q.Dequeue(); v != 4 {
		t.Errorf("front wrong value %d != %d", v, 4)
	}
}

func TestQueue_Back(t *testing.T) {
	t.Parallel()
	q := New[int]()
	q.Enqueue(0)
	if v := q.Back(); v != 0 {
		t.Errorf("front wrong value %d != %d", v, 0)
	}
	q.Enqueue(1)
	if v := q.Back(); v != 1 {
		t.Errorf("front wrong value %d != %d", v, 1)
	}
	q.Enqueue(2)
	if v := q.Back(); v != 2 {
		t.Errorf("front wrong value %d != %d", v, 2)
	}
	q.Enqueue(3)
	if v := q.Back(); v != 3 {
		t.Errorf("front wrong value %d != %d", v, 3)
	}
	q.Enqueue(4)
	if v := q.Back(); v != 4 {
		t.Errorf("front wrong value %d != %d", v, 4)
	}
}

func TestQueue_Size(t *testing.T) {
	t.Parallel()
	q := New[int]()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Dequeue()
	q.Enqueue(4)
	q.Dequeue()
	q.Dequeue()
	q.Enqueue(1)
	if v := q.Size(); v != 3 {
		t.Errorf("wrong size %d != %d", v, 3)
	}
}
