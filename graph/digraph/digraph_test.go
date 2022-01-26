package digraph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/digraph"
	"github.com/DmitryBogomolov/algorithms/graph/graph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func TestAllDigraphEdges_EmptyGraph(t *testing.T) {
	gr := tests.NewTestDigraph(0)
	assert.Equal(t, []graph.Edge(nil), AllDigraphEdges(gr))
}

func TestAllDigraphEdges_NoEdges(t *testing.T) {
	gr := tests.NewTestDigraph(5)
	assert.Equal(t, []graph.Edge(nil), AllDigraphEdges(gr))
}

func TestAllDigraphEdges(t *testing.T) {
	gr := tests.NewTestDigraph(6,
		0, 1,
		1, 2,
		0, 3,
		3, 0,
		4, 3,
		4, 5,
		5, 0,
	)

	ret := AllDigraphEdges(gr)
	assert.Equal(t, []graph.Edge{
		graph.NewEdge(0, 1),
		graph.NewEdge(0, 3),
		graph.NewEdge(1, 2),
		graph.NewEdge(3, 0),
		graph.NewEdge(4, 3),
		graph.NewEdge(4, 5),
		graph.NewEdge(5, 0),
	}, ret)
}
