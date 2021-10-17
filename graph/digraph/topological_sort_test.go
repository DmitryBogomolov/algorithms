package digraph

import (
	"algorithms/graph/internals/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopologicalSort(t *testing.T) {
	var ret []int
	graph := tests.NewTestDigraph(16,
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

	ret = TopologicalSort(graph)
	assert.Equal(t, []int{13, 14, 15, 8, 7, 2, 3, 0, 6, 9, 11, 12, 10, 5, 4, 1}, ret)

	graph.AddDirectedEdge(15, 13)

	ret = TopologicalSort(graph)
	assert.Equal(t, []int(nil), ret)
}
