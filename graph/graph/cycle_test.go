package graph

import (
	"testing"

	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func TestFindCycle(t *testing.T) {
	var target *tests.TestGraph

	target = tests.NewTestGraph(3,
		0, 1,
		1, 2,
	)
	assert.Equal(t, []int(nil), FindCycle(target))

	target.AddEdge(2, 0)
	assert.Equal(t, []int{2, 0, 1, 2}, FindCycle(target))

	target = tests.NewTestGraph(5,
		0, 1,
		0, 2,
		1, 3,
		1, 4,
	)
	assert.Equal(t, []int(nil), FindCycle(target))

	target.AddEdge(2, 3)
	assert.Equal(t, []int{2, 0, 1, 3, 2}, FindCycle(target))
}
