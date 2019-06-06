package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerticesPQ(t *testing.T) {
	pq := newVerticesPQ(10)

	pq.update(2, 1.2)
	pq.update(3, 0.3)
	pq.update(5, 3.1)
	pq.update(8, 2.2)
	pq.update(0, 6.1)
	pq.update(3, 5.4)
	pq.update(8, 0.9)
	pq.update(1, 2.5)

	assert.Equal(t, 6, pq.Len())
	var data []int
	for pq.Len() > 0 {
		data = append(data, pq.pop())
	}
	assert.Equal(t, []int{8, 2, 1, 5, 3, 0}, data)
}

func TestMinimumSpanningTreePrim(t *testing.T) {
	graph := newTestEdgeWeightedGraph(8, []testWeightedEdge{
		{4, 5, 0.35},
		{4, 7, 0.37},
		{5, 7, 0.28},
		{0, 7, 0.16},
		{1, 5, 0.32},
		{0, 4, 0.38},
		{2, 3, 0.17},
		{1, 7, 0.19},
		{0, 2, 0.26},
		{1, 2, 0.36},
		{1, 3, 0.29},
		{2, 7, 0.34},
		{6, 2, 0.40},
		{3, 6, 0.52},
		{6, 0, 0.58},
		{6, 4, 0.93},
	})

	ret := MinimumSpanningTreePrim(graph)

	assert.Equal(t, 8, ret.NumVertices(), "vertices")
	assert.Equal(t, 7, ret.NumEdges(), "edges")
	assert.Equal(t,
		[]Edge{{0, 2}, {0, 7}, {1, 7}, {2, 3}, {2, 6}, {4, 5}, {5, 7}},
		AllGraphEdges(ret),
		"all edges",
	)
	assert.Equal(t,
		[]float64{0.26, 0.16, 0.19, 0.17, 0.4, 0.35, 0.28},
		AllGraphWeights(ret),
		"all weights",
	)
	assert.InDelta(t, 1.81, TotalGraphWeight(ret), 0.0001, "total weight")
}
