package graph

import (
	"algorithms/graph/internals/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCutEdges(t *testing.T) {
	var target *tests.TestGraph
	var ret []Edge

	target = tests.NewTestGraph(7,
		0, 1,
		0, 2,
		1, 3,
		2, 3,
		3, 4,
		4, 5,
		4, 6,
	)

	ret = FindCutEdges(target)
	assert.Equal(t, []Edge{{4, 5}, {4, 6}, {3, 4}}, ret)

	target.AddEdge(5, 6)

	ret = FindCutEdges(target)
	assert.Equal(t, []Edge{{3, 4}}, ret)

	target.AddEdge(1, 5)

	ret = FindCutEdges(target)
	assert.Equal(t, []Edge(nil), ret)
}
