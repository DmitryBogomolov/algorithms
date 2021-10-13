package digraph

import (
	"algorithms/graph/internals"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasDirectedCycle(t *testing.T) {
	var target *internals.TestGraph

	target = internals.NewTestDigraph(4,
		0, 1,
		0, 2,
		2, 3,
		3, 1,
	)
	assert.Equal(t, []int(nil), FindDirectedCycle(target))

	target.AddDirectedEdge(3, 0)
	assert.Equal(t, []int{3, 0, 2, 3}, FindDirectedCycle(target))
}
