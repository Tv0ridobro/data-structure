package search

import (
	"math/rand"
)

type Ordered interface{
~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 |
~float64
}

// OrderStatistics returns kth largest element in given slice
func OrderStatistics[T Ordered](elements []T, k int) T {
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

// partition picks random element as pivot and partitions slice in a way that elements al lower
// indexes are less or equal than pivot
// partition returns index of pivot
func partition[T Ordered](elements []T) int {
	size := len(elements)
	ind := rand.Intn(size)
	elements[ind], elements[size-1] = elements[size-1], elements[ind]
	i, j := 0, 0
	for j < size-1 {
		if elements[j] <= elements[size-1] {
			if elements[j] == elements[size - 1] && rand.Float64() < 0.5 { // to optimize work in case a lot of elements == elements[size - 1]
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
