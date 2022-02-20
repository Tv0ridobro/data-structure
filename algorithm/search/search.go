package search

import (
	"golang.org/x/exp/constraints"
	"math/rand"
)

// OrderStatistics returns kth largest element in given slice
// Changes order in elements
func OrderStatistics[T constraints.Ordered](elements []T, k int) T {
	l, r := 0, len(elements)
	for {
		p := partition(elements[l:r]) + l
		switch {
		case p == k:
			return elements[p]
		case k < p:
			r = p
		default:
			l = p + 1
		}
	}
}

// partition picks random element as pivot and partitions slice in a way
// that elements at lower indexes are less or equal than pivot
// partition returns index of pivot
func partition[T constraints.Ordered](elements []T) int {
	size := len(elements)
	ind := rand.Intn(size)
	elements[ind], elements[size-1] = elements[size-1], elements[ind]
	i, j := 0, 0
	for j < size-1 {
		if elements[j] <= elements[size-1] {
			if elements[j] == elements[size-1] && rand.Float64() < 0.5 { // to optimize work in case a lot of elements == elements[size - 1]
				j++
				continue
			}
			elements[i], elements[j] = elements[j], elements[i]
			i++
		}
		j++
	}
	elements[i], elements[size-1] = elements[size-1], elements[i]
	return i
}
