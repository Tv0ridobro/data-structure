package graphalgo

import (
	"github.com/Tv0ridobro/data-structure/graph"
)

// FindCycle returns cycle if there is any
// otherwise empty slice returned
func FindCycle[T any](g *graph.Graph[T]) []int {
	visited := make([]byte, g.Size())
	for i := 0; i < g.Size(); i++ {
		order := []int{i}
		if visited[i] == 0 && dfsCycle(i, -1, visited, g, &order) {
			last := order[len(order)-1]
			for j := len(order) - 2; j >= 0; j-- {
				if order[j] == last {
					return order[j:]
				}
			}
		}
	}
	return nil
}

// dfsCycle helper function to find cycle using dfs
func dfsCycle[T any](vertex, from int, visited []byte, g *graph.Graph[T], order *[]int) bool {
	visited[vertex] = 1
	for _, e := range g.Edges[vertex] {
		if !g.IsDirected() && e.To == from {
			continue
		}
		*order = append(*order, e.To)
		if visited[e.To] == 1 {
			return true
		}
		if visited[e.To] == 0 && dfsCycle(e.To, vertex, visited, g, order) {
			return true
		}
	}
	visited[vertex] = 2
	*order = (*order)[:len(*order)-1]
	return false
}
