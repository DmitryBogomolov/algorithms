package digraph

import (
	"algorithms/graph/internals/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasDirectedCycle(t *testing.T) {
	var target *tests.TestGraph

	target = tests.NewTestDigraph(4,
		0, 1,
		0, 2,
		2, 3,
		3, 1,
	)
	assert.Equal(t, []int(nil), FindDirectedCycle(target))

	target.AddDirectedEdge(3, 0)
	assert.Equal(t, []int{3, 0, 2, 3}, FindDirectedCycle(target))
}
