package graphalgo

import (
	"testing"

	"github.com/Tv0ridobro/data-structure/graph"
)

func TestTopologicalSort(t *testing.T) {
	t.Parallel()
	g := graph.New[int](12, true)
	g.AddEdge(7, 11, 0)
	g.AddEdge(7, 8, 0)
	g.AddEdge(5, 11, 0)
	g.AddEdge(3, 8, 0)
	g.AddEdge(11, 2, 0)
	g.AddEdge(11, 9, 0)
	g.AddEdge(11, 10, 0)
	g.AddEdge(8, 9, 0)
	g.AddEdge(3, 10, 0)
	g.AddEdge(5, 7, 0)
}
