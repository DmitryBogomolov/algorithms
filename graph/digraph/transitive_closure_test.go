package digraph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/digraph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func TestTransitiveClosure_EmptyGraph(t *testing.T) {
	gr := tests.NewTestDigraph(0)

	ret := BuildTransitiveClosure(gr)
	assert.Equal(t, 0, ret.NumVertices())
}

func TestTransitiveClosure_NoEdges(t *testing.T) {
	gr := tests.NewTestDigraph(5)

	ret := BuildTransitiveClosure(gr)
	assert.Equal(t, 5, ret.NumVertices())
	assert.Equal(t, false, ret.Reachable(0, 1), "0 - 1")
	assert.Equal(t, false, ret.Reachable(1, 2), "1 - 2")
	assert.Equal(t, false, ret.Reachable(2, 3), "2 - 3")
}

func TestTransitiveClosure(t *testing.T) {
	gr := tests.NewTestDigraph(6,
		0, 1,
		1, 2,
		0, 3,
		3, 2,
		4, 3,
		4, 5,
		5, 4,
	)

	ret := BuildTransitiveClosure(gr)
	assert.Equal(t, gr.NumVertices(), ret.NumVertices())
	assert.Equal(t, true, ret.Reachable(0, 2), "0 - 2")
	assert.Equal(t, false, ret.Reachable(1, 0), "1 - 0")
	assert.Equal(t, true, ret.Reachable(3, 2), "3 - 2")
	assert.Equal(t, true, ret.Reachable(4, 3), "4 - 3")
	assert.Equal(t, true, ret.Reachable(4, 2), "4 - 2")
	assert.Equal(t, false, ret.Reachable(5, 0), "5 - 0")
	assert.Equal(t, true, ret.Reachable(5, 2), "5 - 2")
}
