package graph

import (
	"testing"

	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func TestIsBipartite(t *testing.T) {
	var target *tests.TestGraph

	target = tests.NewTestGraph(6,
		0, 1,
		0, 3,
		1, 2,
		2, 3,
		1, 4,
		2, 5,
		4, 5,
	)
	assert.Equal(t, true, IsBipartite(target))

	target = tests.NewTestGraph(6,
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
