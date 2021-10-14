package ewgraph

import (
	"algorithms/graph/graph"
	"algorithms/graph/internals"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerticesIndexPQ(t *testing.T) {
	pq := newVerticesIndexPQ(10)

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
	target := internals.NewTestEdgeWeightedGraph(8, []internals.TestWeightedEdge{
		{V1: 4, V2: 5, Weight: 0.35},
		{V1: 4, V2: 7, Weight: 0.37},
		{V1: 5, V2: 7, Weight: 0.28},
		{V1: 0, V2: 7, Weight: 0.16},
		{V1: 1, V2: 5, Weight: 0.32},
		{V1: 0, V2: 4, Weight: 0.38},
		{V1: 2, V2: 3, Weight: 0.17},
		{V1: 1, V2: 7, Weight: 0.19},
		{V1: 0, V2: 2, Weight: 0.26},
		{V1: 1, V2: 2, Weight: 0.36},
		{V1: 1, V2: 3, Weight: 0.29},
		{V1: 2, V2: 7, Weight: 0.34},
		{V1: 6, V2: 2, Weight: 0.40},
		{V1: 3, V2: 6, Weight: 0.52},
		{V1: 6, V2: 0, Weight: 0.58},
		{V1: 6, V2: 4, Weight: 0.93},
	})

	ret := MinimumSpanningTreePrim(target)

	assert.Equal(t, 8, ret.NumVertices(), "vertices")
	assert.Equal(t, 7, ret.NumEdges(), "edges")
	assert.Equal(t,
		[]graph.Edge{
			graph.NewEdge(0, 2),
			graph.NewEdge(0, 7),
			graph.NewEdge(1, 7),
			graph.NewEdge(2, 3),
			graph.NewEdge(2, 6),
			graph.NewEdge(4, 5),
			graph.NewEdge(5, 7),
		},
		graph.AllGraphEdges(ret),
		"all edges",
	)
	assert.Equal(t,
		[]float64{0.26, 0.16, 0.19, 0.17, 0.4, 0.35, 0.28},
		AllGraphWeights(ret),
		"all weights",
	)
	assert.InDelta(t, 1.81, TotalGraphWeight(ret), 0.0001, "total weight")
}
