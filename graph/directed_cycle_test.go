package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasDirectedCycle(t *testing.T) {
	var graph *testGraph

	graph = newTestDigraph(4,
		0, 1,
		0, 2,
		2, 3,
		3, 1,
	)
	assert.Equal(t, false, HasDirectedCycle(graph))

	graph.addDirectedEdge(1, 0)
	assert.Equal(t, true, HasDirectedCycle(graph))
}
