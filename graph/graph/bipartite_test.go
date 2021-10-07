package graph

import (
	"algorithms/graph/internals"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBipartite(t *testing.T) {
	var target *internals.TestGraph

	target = internals.NewTestGraph(6,
		0, 1,
		0, 3,
		1, 2,
		2, 3,
		1, 4,
		2, 5,
		4, 5,
	)
	assert.Equal(t, true, IsBipartite(target))

	target = internals.NewTestGraph(6,
		0, 1,
		0, 3,
		1, 2,
		2, 3,
		2, 4,
		1, 4,
		2, 5,
		4, 5,
	)
	assert.Equal(t, false, IsBipartite(target))
}
