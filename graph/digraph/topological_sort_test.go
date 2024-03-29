package digraph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/digraph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func TestTopologicalSort_EmptyGraph(t *testing.T) {
	gr := tests.NewTestDigraph(0)
	assert.Equal(t, []int(nil), TopologicalSort(gr))
}

func TestTopologicalSort_NoEdges(t *testing.T) {
	gr := tests.NewTestDigraph(5)
	assert.Equal(t, []int{4, 3, 2, 1, 0}, TopologicalSort(gr))
}

func TestTopologicalSort(t *testing.T) {
	gr := tests.NewTestDigraph(16,
		0, 1,
		0, 5,
		0, 6,
		2, 0,
		2, 3,
		3, 5,
		5, 4,
		6, 4,
		6, 9,
		7, 6,
		8, 7,
		9, 10,
		9, 11,
		9, 12,
		11, 12,
		13, 14,
		13, 15,
		14, 15,
	)

	ret := TopologicalSort(gr)
	assert.Equal(t, []int{13, 14, 15, 8, 7, 2, 3, 0, 6, 9, 11, 12, 10, 5, 4, 1}, ret)

	gr.AddDirectedEdge(15, 13)

	ret = TopologicalSort(gr)
	assert.Equal(t, []int(nil), ret)
}
