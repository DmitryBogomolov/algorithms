package ewgraph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/ewgraph"
	"github.com/DmitryBogomolov/algorithms/graph/graph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func TestBuildMinimumSpanningTreeKruskal(t *testing.T) {
	gr := tests.NewTestEdgeWeightedGraph(8, []tests.TestWeightedEdge{
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

	ret := BuildMinimumSpanningTreeKruskal(gr)

	assert.Equal(t, 8, ret.NumVertices(), "vertices")
	assert.Equal(t, 7, ret.NumEdges(), "edges")
	assert.Equal(t,
		[]graph.Edge{
			graph.NewEdge(0, 7),
			graph.NewEdge(0, 2),
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
		[]float64{0.16, 0.26, 0.19, 0.17, 0.4, 0.35, 0.28},
		AllGraphWeights(ret),
		"all weights",
	)
	assert.InDelta(t, 1.81, TotalGraphWeight(ret), 0.0001, "total weight")
}
