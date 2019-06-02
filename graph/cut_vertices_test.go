package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCutVertices(t *testing.T) {
	var graph *testGraph
	var ret []int

	graph = newTestGraph(7,
		0, 1,
		0, 2,
		1, 3,
		2, 3,
		3, 4,
		4, 5,
		4, 6,
	)

	ret = FindCutVertices(graph)
	assert.Equal(t, []int{3, 4}, ret)

	graph.addEdge(1, 5)

	ret = FindCutVertices(graph)
	assert.Equal(t, []int{4}, ret)

	graph.addEdge(2, 6)

	ret = FindCutVertices(graph)
	assert.Equal(t, []int(nil), ret)
}
