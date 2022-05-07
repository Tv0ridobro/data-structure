package graphalgo

import (
	"testing"

	"github.com/Tv0ridobro/data-structure/graph"
)

func TestFindCycle(t *testing.T) {
	t.Parallel()
	g := graph.New[int](4, false)
	g.AddEdge(0, 1, 0)
	g.AddEdge(1, 3, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(0, 2, 0)
}
