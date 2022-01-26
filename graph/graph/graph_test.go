package graph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/graph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func TestAllGraphEdges_EmptyGraph(t *testing.T) {
	gr := tests.NewTestGraph(0)

	ret := AllGraphEdges(gr)
	assert.Equal(t, []Edge(nil), ret)
}

func TestAllGraphEdges_NoEdges(t *testing.T) {
	gr := tests.NewTestGraph(5)

	ret := AllGraphEdges(gr)
	assert.Equal(t, []Edge(nil), ret)
}

func TestAllGraphEdges(t *testing.T) {
	gr := tests.NewTestGraph(6,
		0, 1,
		1, 2,
		0, 3,
		3, 2,
		4, 3,
		4, 5,
		5, 0,
	)

	ret := AllGraphEdges(gr)
	assert.Equal(t, []Edge{
		NewEdge(0, 1),
		NewEdge(0, 3),
		NewEdge(0, 5),
		NewEdge(1, 2),
		NewEdge(2, 3),
		NewEdge(3, 4),
		NewEdge(4, 5),
	}, ret)
}
