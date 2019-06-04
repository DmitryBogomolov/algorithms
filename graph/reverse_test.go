package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	graph := newTestDigraph(4,
		0, 1,
		0, 2,
		2, 1,
		1, 3,
		2, 3,
	)

	ret := Reverse(graph)

	assert.Equal(t, 4, ret.NumVertices(), "vertices")
	assert.Equal(t, 5, ret.NumEdges(), "edges")
	assert.Equal(t, []int(nil), ret.AdjacentVertices(0), "vertex 0 adjacency")
	assert.Equal(t, []int{0, 2}, ret.AdjacentVertices(1), "vertex 1 adjacency")
	assert.Equal(t, []int{0}, ret.AdjacentVertices(2), "vertex 2 adjacency")
	assert.Equal(t, []int{1, 2}, ret.AdjacentVertices(3), "vertex 3 adjacency")
}
