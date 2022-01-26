package ewgraph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/ewgraph"
	"github.com/stretchr/testify/assert"
)

func TestNewImplEdgeWeightedGraph(t *testing.T) {
	gr := NewImplEdgeWeightedGraph(
		5, 3,
		[][]int{
			{1, 3},
			{0, 2},
			{1},
			{0},
			nil,
		},
		[][]float64{
			{1.1, 0.4},
			{1.1, 2.1},
			{2.1},
			{0.4},
			nil,
		},
	)

	assert.Equal(t, gr.NumVertices(), 5, "vertices")
	assert.Equal(t, gr.NumEdges(), 3, "edges")
	assert.Equal(t, gr.AdjacentVertices(0), []int{1, 3}, "vertex 0")
	assert.Equal(t, gr.AdjacentVertices(1), []int{0, 2}, "vertex 1")
	assert.Equal(t, gr.AdjacentVertices(2), []int{1}, "vertex 2")
	assert.Equal(t, gr.AdjacentVertices(3), []int{0}, "vertex 3")
	assert.Equal(t, gr.AdjacentVertices(4), []int(nil), "vertex 4")
	assert.Equal(t, gr.AdjacentWeights(0), []float64{1.1, 0.4}, "vertex 0")
	assert.Equal(t, gr.AdjacentWeights(1), []float64{1.1, 2.1}, "vertex 1")
	assert.Equal(t, gr.AdjacentWeights(2), []float64{2.1}, "vertex 2")
	assert.Equal(t, gr.AdjacentWeights(3), []float64{0.4}, "vertex 3")
	assert.Equal(t, gr.AdjacentWeights(4), []float64(nil), "vertex 4")
}
