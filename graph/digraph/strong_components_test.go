package digraph

import (
	"testing"

	"github.com/DmitryBogomolov/algorithms/graph/internals/tests"

	"github.com/stretchr/testify/assert"
)

func TestFindStrongComponents(t *testing.T) {
	graph := tests.NewTestDigraph(13,
		4, 2,
		2, 3,
		3, 2,
		6, 0,
		0, 1,
		2, 0,
		11, 12,
		12, 9,
		9, 10,
		9, 11,
		7, 9,
		10, 12,
		11, 4,
		4, 3,
		3, 5,
		6, 8,
		8, 6,
		5, 4,
		0, 5,
		6, 4,
		6, 9,
		7, 6,
	)

	ret := FindStrongComponents(graph)

	assert.Equal(t, 5, ret.Count(), "count")

	components := make([]int, graph.NumVertices())
	for v := 0; v < graph.NumVertices(); v++ {
		components[v] = ret.ComponentID(v)
	}
	assert.Equal(t,
		[]int{1, 0, 1, 1, 1, 1, 3, 4, 3, 2, 2, 2, 2},
		components,
		"components",
	)

	assert.Equal(t, []int{1}, ret.Component(0), "component 0")
	assert.Equal(t, []int{0, 2, 3, 4, 5}, ret.Component(1), "component 1")
	assert.Equal(t, []int{9, 10, 11, 12}, ret.Component(2), "component 2")
	assert.Equal(t, []int{6, 8}, ret.Component(3), "component 3")
	assert.Equal(t, []int{7}, ret.Component(4), "component 4")

	assert.Equal(t, true, ret.Connected(1, 1), "1 - 1")
	assert.Equal(t, false, ret.Connected(2, 1), "2 - 1")
	assert.Equal(t, false, ret.Connected(9, 7), "9 - 7")
	assert.Equal(t, true, ret.Connected(2, 3), "2 - 3")
	assert.Equal(t, true, ret.Connected(8, 6), "8 - 6")
	assert.Equal(t, true, ret.Connected(10, 11), "10 - 11")
	assert.Equal(t, true, ret.Connected(0, 5), "0 - 5")
}
