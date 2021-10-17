package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewImplGraph(t *testing.T) {
	target := NewImplGraph(5, 3, [][]int{
		{1, 3},
		{0, 2},
		{1},
		{0},
		nil,
	})

	assert.Equal(t, target.NumVertices(), 5, "vertices")
	assert.Equal(t, target.NumEdges(), 3, "edges")
	assert.Equal(t, target.AdjacentVertices(0), []int{1, 3}, "vertex 0")
	assert.Equal(t, target.AdjacentVertices(1), []int{0, 2}, "vertex 1")
	assert.Equal(t, target.AdjacentVertices(2), []int{1}, "vertex 2")
	assert.Equal(t, target.AdjacentVertices(3), []int{0}, "vertex 3")
	assert.Equal(t, target.AdjacentVertices(4), []int(nil), "vertex 4")
}
