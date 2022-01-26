package ewdigraph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/ewdigraph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func TestAllDigraphWeights_EmptyGraph(t *testing.T) {
	gr := tests.NewTestEdgeWeightedDigraph(0, nil)
	assert.Equal(t, []float64(nil), AllDigraphWeights(gr))
	assert.Equal(t, 0.0, TotalDigraphWeight(gr))
}

func TestAllDigraphWeights_NoEdges(t *testing.T) {
	gr := tests.NewTestEdgeWeightedDigraph(4, nil)
	assert.Equal(t, []float64(nil), AllDigraphWeights(gr))
	assert.Equal(t, 0.0, TotalDigraphWeight(gr))
}

func TestAllDigraphWeights(t *testing.T) {
	gr := tests.NewTestEdgeWeightedDigraph(6, []tests.TestWeightedEdge{
		{V1: 0, V2: 1, Weight: 1.2},
		{V1: 1, V2: 2, Weight: 2.3},
		{V1: 0, V2: 3, Weight: 3.1},
		{V1: 3, V2: 0, Weight: 4.1},
		{V1: 4, V2: 3, Weight: 5.1},
		{V1: 4, V2: 5, Weight: 1.6},
		{V1: 5, V2: 0, Weight: 2.2},
	})

	assert.Equal(t, []float64{1.2, 3.1, 2.3, 4.1, 5.1, 1.6, 2.2}, AllDigraphWeights(gr))
	assert.InDelta(t, 19.6, TotalDigraphWeight(gr), 0.0001)
}
