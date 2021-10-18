package ewdigraph

import (
	"testing"

	"github.com/DmitryBogomolov/algorithms/graph/internals/tests"

	"github.com/stretchr/testify/assert"
)

func TestAllDigraphWeights(t *testing.T) {
	target := tests.NewTestEdgeWeightedDigraph(6, []tests.TestWeightedEdge{
		{V1: 0, V2: 1, Weight: 1.2},
		{V1: 1, V2: 2, Weight: 2.3},
		{V1: 0, V2: 3, Weight: 3.1},
		{V1: 3, V2: 0, Weight: 4.1},
		{V1: 4, V2: 3, Weight: 5.1},
		{V1: 4, V2: 5, Weight: 1.6},
		{V1: 5, V2: 0, Weight: 2.2},
	})

	ret := AllDigraphWeights(target)

	assert.Equal(t, []float64{1.2, 3.1, 2.3, 4.1, 5.1, 1.6, 2.2}, ret)
	assert.InDelta(t, 19.6, TotalDigraphWeight(target), 0.0001)
}
