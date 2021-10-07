package graph

import (
	"algorithms/graph/internals"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCutVertices(t *testing.T) {
	var target *internals.TestGraph
	var ret []int

	target = internals.NewTestGraph(7,
		0, 1,
		0, 2,
		1, 3,
		2, 3,
		3, 4,
		4, 5,
		4, 6,
	)

	ret = FindCutVertices(target)
	assert.Equal(t, []int{3, 4}, ret)

	target.AddEdge(1, 5)

	ret = FindCutVertices(target)
	assert.Equal(t, []int{4}, ret)

	target.AddEdge(2, 6)

	ret = FindCutVertices(target)
	assert.Equal(t, []int(nil), ret)
}
