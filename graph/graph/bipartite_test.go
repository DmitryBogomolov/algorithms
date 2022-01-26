package graph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/graph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func TestIsBipartite_EmptyGraph(t *testing.T) {
	gr := tests.NewTestGraph(0)
	assert.Equal(t, true, IsBipartite(gr))
}

func TestIsBipartite_NoEdges(t *testing.T) {
	gr := tests.NewTestGraph(4)
	assert.Equal(t, true, IsBipartite(gr))
}

func TestIsBipartite(t *testing.T) {
	var gr *tests.TestGraph

	gr = tests.NewTestGraph(3,
		0, 1,
		0, 2,
	)
	assert.Equal(t, true, IsBipartite(gr))
	gr.AddEdge(1, 2)
	assert.Equal(t, false, IsBipartite(gr))

	gr = tests.NewTestGraph(6,
		0, 1,
		0, 3,
		1, 2,
		2, 3,
		1, 4,
		2, 5,
		4, 5,
	)
	assert.Equal(t, true, IsBipartite(gr))

	gr = tests.NewTestGraph(6,
		0, 1,
		0, 3,
		1, 2,
		2, 3,
		2, 4,
		1, 4,
		2, 5,
		4, 5,
	)
	assert.Equal(t, false, IsBipartite(gr))
}
