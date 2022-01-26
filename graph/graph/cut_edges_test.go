package graph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/graph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func TestFindCutEdges_EmptyGraph(t *testing.T) {
	gr := tests.NewTestGraph(0)
	assert.Equal(t, []Edge(nil), FindCutEdges(gr))
}

func TestFindCutEdges_NoEdges(t *testing.T) {
	gr := tests.NewTestGraph(5)
	assert.Equal(t, []Edge(nil), FindCutEdges(gr))
}

func TestFindCutEdges(t *testing.T) {
	var gr *tests.TestGraph
	var ret []Edge

	gr = tests.NewTestGraph(7,
		0, 1,
		0, 2,
		1, 3,
		2, 3,
		3, 4,
		4, 5,
		4, 6,
	)

	ret = FindCutEdges(gr)
	assert.Equal(t, []Edge{NewEdge(4, 5), NewEdge(4, 6), NewEdge(3, 4)}, ret)

	gr.AddEdge(5, 6)

	ret = FindCutEdges(gr)
	assert.Equal(t, []Edge{NewEdge(3, 4)}, ret)

	gr.AddEdge(1, 5)

	ret = FindCutEdges(gr)
	assert.Equal(t, []Edge(nil), ret)
}
