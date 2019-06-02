package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	var graph *testGraph

	graph = newTestGraph(3,
		0, 1,
		1, 2,
	)
	assert.Equal(t, false, HasCycle(graph))

	graph.addEdge(2, 0)
	assert.Equal(t, true, HasCycle(graph))

	graph = newTestGraph(5,
		0, 1,
		0, 2,
		1, 3,
		1, 4,
	)
	assert.Equal(t, false, HasCycle(graph))

	graph.addEdge(2, 3)
	assert.Equal(t, true, HasCycle(graph))
}
