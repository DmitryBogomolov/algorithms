package digraph

import (
	"testing"

	"github.com/DmitryBogomolov/algorithms/graph/internals/tests"

	"github.com/stretchr/testify/assert"
)

func TestReverseDigraph(t *testing.T) {
	graph := tests.NewTestDigraph(4,
		0, 1,
		0, 2,
		2, 1,
		1, 3,
		2, 3,
	)

	ret := ReverseDigraph(graph)

	assert.Equal(t, 4, ret.NumVertices(), "vertices")
	assert.Equal(t, 5, ret.NumEdges(), "edges")
	assert.Equal(t, []int(nil), ret.AdjacentVertices(0), "vertex 0 adjacency")
	assert.Equal(t, []int{0, 2}, ret.AdjacentVertices(1), "vertex 1 adjacency")
	assert.Equal(t, []int{0}, ret.AdjacentVertices(2), "vertex 2 adjacency")
	assert.Equal(t, []int{1, 2}, ret.AdjacentVertices(3), "vertex 3 adjacency")
}

func TestReverseReversedDigraph(t *testing.T) {
	graph := tests.NewTestDigraph(4,
		0, 1,
		0, 2,
		2, 1,
		1, 3,
		2, 3,
	)

	ret, ok := ReverseDigraph(graph).(ReversibleDigraph)

	assert.Equal(t, true, ok, "reversible")
	assert.Equal(t, graph, ret.Reverse(), "reverse to original")
	assert.Equal(t, graph, ReverseDigraph(ret), "reverse to original")
}
