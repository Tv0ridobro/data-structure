package graphalgo

import (
	"github.com/Tv0ridobro/data-structure/graph"
	"github.com/Tv0ridobro/data-structure/queue"
	"github.com/Tv0ridobro/data-structure/slices"
)

// FindShortestPath returns vertices that shortest path contains
// returns nil if start == end or if there are no path
func FindShortestPath[T any](g *graph.Graph[T], start, end int) []int {
	if start < 0 || end < 0 || start >= g.Size() || end >= g.Size() {
		return nil
	}
	if start == end {
		return nil
	}
	q := queue.New[int]()
	q.Enqueue(start)
	from := make([]int, g.Size())
	for i := 0; i < g.Size(); i++ {
		from[i] = -1
	}
	from[start] = start
	for q.Size() != 0 {
		v := q.Dequeue()
		if bfsShortestPath(v, end, g, q, from) {
			ans := make([]int, 0)
			n := end
			for n != start {
				ans = append(ans, n)
				n = from[n]
			}
			ans = append(ans, start)
			return slices.Reverse(ans)
		}
	}
	return nil
}

// bfsShortestPath helper function to find shortest path
func bfsShortestPath[T any](now, end int, g *graph.Graph[T], q *queue.Queue[int], from []int) bool {
	for _, e := range g.Edges[now] {
		if from[e.To] == -1 {
			from[e.To] = now
			if e.To == end {
				return true
			}
			q.Enqueue(e.To)
		}
	}
	return false
}
