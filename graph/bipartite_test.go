package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBipartite(t *testing.T) {
	var graph *testGraph

	graph = newTestGraph(6,
		0, 1,
		0, 3,
		1, 2,
		2, 3,
		1, 4,
		2, 5,
		4, 5,
	)
	assert.Equal(t, true, IsBipartite(graph))

	graph = newTestGraph(6,
		0, 1,
		0, 3,
		1, 2,
		2, 3,
		2, 4,
		1, 4,
		2, 5,
		4, 5,
	)
	assert.Equal(t, false, IsBipartite(graph))
}
