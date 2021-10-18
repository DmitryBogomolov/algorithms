package graph

import (
	"testing"

	"github.com/DmitryBogomolov/algorithms/graph/internals/tests"

	"github.com/stretchr/testify/assert"
)

func TestAllGraphEdges(t *testing.T) {
	target := tests.NewTestGraph(6,
		0, 1,
		1, 2,
		0, 3,
		3, 2,
		4, 3,
		4, 5,
		5, 0,
	)

	ret := AllGraphEdges(target)

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
