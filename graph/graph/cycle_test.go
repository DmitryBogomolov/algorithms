package graph

import (
	"algorithms/graph/internals"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	var target *internals.TestGraph

	target = internals.NewTestGraph(3,
		0, 1,
		1, 2,
	)
	assert.Equal(t, false, HasCycle(target))

	target.AddEdge(2, 0)
	assert.Equal(t, true, HasCycle(target))

	target = internals.NewTestGraph(5,
		0, 1,
		0, 2,
		1, 3,
		1, 4,
	)
	assert.Equal(t, false, HasCycle(target))

	target.AddEdge(2, 3)
	assert.Equal(t, true, HasCycle(target))
}
