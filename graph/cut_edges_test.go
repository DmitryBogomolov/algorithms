package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCutEdges(t *testing.T) {
	var graph *testGraph
	var ret []CutEdge

	graph = newTestGraph(7,
		0, 1,
		0, 2,
		1, 3,
		2, 3,
		3, 4,
		4, 5,
		4, 6,
	)

	ret = FindCutEdges(graph)
	assert.Equal(t, []CutEdge{CutEdge{4, 5}, CutEdge{4, 6}, CutEdge{3, 4}}, ret)

	graph.addEdge(5, 6)

	ret = FindCutEdges(graph)
	assert.Equal(t, []CutEdge{CutEdge{3, 4}}, ret)

	graph.addEdge(1, 5)

	ret = FindCutEdges(graph)
	assert.Equal(t, []CutEdge(nil), ret)
}
