package graph

import (
	"testing"

	"github.com/DmitryBogomolov/algorithms/graph/internals/tests"

	"github.com/stretchr/testify/assert"
)

func TestFindConnectedComponents(t *testing.T) {
	target := tests.NewTestGraph(8,
		0, 1,
		1, 4,
		4, 7,
		7, 2,
		2, 0,
		1, 2,
		2, 4,
		5, 6,
	)

	ret := FindConnectedComponents(target)

	assert.Equal(t, 3, ret.Count(), "count")

	components := make([]int, target.NumVertices())
	for v := 0; v < target.NumVertices(); v++ {
		components[v] = ret.ComponentID(v)
	}
	assert.Equal(t,
		[]int{0, 0, 0, 1, 0, 2, 2, 0},
		components,
		"components",
	)

	assert.Equal(t, []int{0, 1, 2, 4, 7}, ret.Component(0), "component 0")
	assert.Equal(t, []int{3}, ret.Component(1), "component 1")
	assert.Equal(t, []int{5, 6}, ret.Component(2), "component 2")

	assert.Equal(t, true, ret.Connected(0, 7), "0 - 7")
	assert.Equal(t, false, ret.Connected(2, 5), "2 - 5")
	assert.Equal(t, false, ret.Connected(4, 3), "4 - 3")
	assert.Equal(t, true, ret.Connected(3, 3), "3 - 3")
	assert.Equal(t, true, ret.Connected(5, 6), "5 - 6")
	assert.Equal(t, true, ret.Connected(7, 1), "7 - 1")
	assert.Equal(t, true, ret.Connected(2, 4), "2 - 4")
}
