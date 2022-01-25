package graph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/graph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func TestFindCutVertices_EmptyGraph(t *testing.T) {
	gr := tests.NewTestGraph(0)
	assert.Equal(t, []int(nil), FindCutVertices(gr))
}

func TestFindCutVertices_NoEdges(t *testing.T) {
	gr := tests.NewTestGraph(5)
	assert.Equal(t, []int(nil), FindCutVertices(gr))
}

func TestFindCutVertices(t *testing.T) {
	var gr *tests.TestGraph
	var ret []int

	gr = tests.NewTestGraph(7,
		0, 1,
		0, 2,
		1, 3,
		2, 3,
		3, 4,
		4, 5,
		4, 6,
	)

	ret = FindCutVertices(gr)
	assert.Equal(t, []int{3, 4}, ret)

	gr.AddEdge(1, 5)

	ret = FindCutVertices(gr)
	assert.Equal(t, []int{4}, ret)

	gr.AddEdge(2, 6)

	ret = FindCutVertices(gr)
	assert.Equal(t, []int(nil), ret)
}
