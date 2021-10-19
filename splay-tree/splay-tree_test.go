package splay_tree

import (
	"github.com/Tv0ridobro/data-structure/util"
	"math/rand"
	"sort"
	"testing"
)

func TestAll(t *testing.T) {
	tr := New[int]()
	for i := 0; i < 100; i++ {
		tr.Insert(i)
		if tr.Size() != i+1 {
			t.Errorf("wrong size expected %d  got %d", i+1, i)
		}
	}
	for i := 0; i < 100; i++ {
		if !tr.Contains(i) {
			t.Errorf("dosnt contain %d", i)
		}
	}
	for i := 0; i < 100; i++ {
		if i%7 == 4 {
			tr.Remove(i)
		}
		tr.Remove(i + 200)
	}
	for i := 0; i < 100; i++ {
		if i%7 == 4 {
			if tr.Contains(i) {
				t.Errorf("contain %d", i)
			}
		}
	}
}

func TestAll2(t *testing.T) {
	tr := New[int]()
	permutation := rand.Perm(1000000)
	for i := range permutation {
		tr.Insert(i)
	}
	sort.Ints(permutation)
	if !util.Equal(tr.GetAll(), permutation) {
		t.Errorf("permutation doesn't equal getAll call")
	}
	if tr.Size() != 1000000 {
		t.Errorf("wrong size")
	}
	rand.Shuffle(len(permutation), func(i, j int) {
		permutation[i], permutation[j] = permutation[i], permutation[j]
	})
	for i := range permutation {
		tr.Remove(i)
	}
	if tr.root != nil {
		t.Errorf("root is not nil")
	}
}

func TestSplayTree_Kth(t *testing.T) {
	s := New[int]()
	for i := 0; i < 1000; i++ {
		s.Insert(i)
	}
	for i := 0; i < 1000; i++ {
		if v := s.Kth(i); v != i {
			t.Errorf("Kth != ans %d %d", v, i)
		}
	}
}
