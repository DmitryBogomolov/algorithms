package digraph

import (
	"algorithms/graph/graph"
	"algorithms/graph/internals"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllDigraphEdges(t *testing.T) {
	target := internals.NewTestDigraph(6,
		0, 1,
		1, 2,
		0, 3,
		3, 0,
		4, 3,
		4, 5,
		5, 0,
	)

	ret := AllDigraphEdges(target)

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
