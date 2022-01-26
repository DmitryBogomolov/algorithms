package digraph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/digraph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func TestHasDirectedCycle_EmptyGraph(t *testing.T) {
	gr := tests.NewTestDigraph(0)
	ret := FindDirectedCycle(gr)
	assert.Equal(t, []int(nil), ret)
}

func TestHasDirectedCycle_NoEdges(t *testing.T) {
	gr := tests.NewTestDigraph(4)
	ret := FindDirectedCycle(gr)
	assert.Equal(t, []int(nil), ret)
}

func TestHasDirectedCycle(t *testing.T) {
	gr := tests.NewTestDigraph(4,
		0, 1,
		0, 2,
		2, 3,
		3, 1,
	)
	assert.Equal(t, []int(nil), FindDirectedCycle(gr))

	gr.AddDirectedEdge(3, 0)
	assert.Equal(t, []int{3, 0, 2, 3}, FindDirectedCycle(gr))
}
