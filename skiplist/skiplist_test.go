package skiplist

import (
	"math/rand"
	"testing"
)

func TestSkipList_Insert(t *testing.T) {
	t.Parallel()
	sl := NewWithProbability[int](0.9)
	rand.Seed(10)
	for i := 0; i < 100; i++ {
		sl.Insert(rand.Intn(1000))
	}
}
