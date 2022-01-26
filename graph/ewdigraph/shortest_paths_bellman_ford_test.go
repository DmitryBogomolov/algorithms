package ewdigraph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/ewdigraph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func TestFindShortedPathsBellmanFord(t *testing.T) {
	gr := tests.NewTestEdgeWeightedDigraph(8, []tests.TestWeightedEdge{
		{V1: 4, V2: 5, Weight: 0.35},
		{V1: 5, V2: 4, Weight: 0.35},
		{V1: 4, V2: 7, Weight: 0.37},
		{V1: 5, V2: 7, Weight: 0.28},
		{V1: 7, V2: 5, Weight: 0.28},
		{V1: 5, V2: 1, Weight: 0.32},
		{V1: 0, V2: 4, Weight: 0.38},
		{V1: 0, V2: 2, Weight: 0.26},
		{V1: 7, V2: 3, Weight: 0.39},
		{V1: 1, V2: 3, Weight: 0.29},
		{V1: 2, V2: 7, Weight: 0.34},
		{V1: 6, V2: 2, Weight: -1.20},
		{V1: 3, V2: 6, Weight: 0.52},
		{V1: 6, V2: 0, Weight: -1.40},
		{V1: 6, V2: 4, Weight: -1.25},
	})

	paths, _ := FindShortedPathsBellmanFord(gr, 0)

	check := func(vertexID int, expectedVertices []int, expectedWeight float64) {
		vertices, weight := paths.PathTo(vertexID)
		assert.Equal(t, expectedVertices, vertices)
		assert.InDelta(t, expectedWeight, weight, 0.0001)
	}

	check(0, []int{0}, 0.0)
	check(1, []int{0, 2, 7, 3, 6, 4, 5, 1}, 0.93)
	check(2, []int{0, 2}, 0.26)
	check(3, []int{0, 2, 7, 3}, 0.99)
	check(4, []int{0, 2, 7, 3, 6, 4}, 0.26)
	check(5, []int{0, 2, 7, 3, 6, 4, 5}, 0.61)
	check(6, []int{0, 2, 7, 3, 6}, 1.51)
	check(7, []int{0, 2, 7}, 0.6)
}
