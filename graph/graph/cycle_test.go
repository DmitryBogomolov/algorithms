package graph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/graph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func TestFindCycle_EmptyGraph(t *testing.T) {
	gr := tests.NewTestGraph(0)
	ret := FindCycle(gr)
	assert.Equal(t, []int(nil), ret)
}

func TestFindCycle_NoEdges(t *testing.T) {
	gr := tests.NewTestGraph(4)
	ret := FindCycle(gr)
	assert.Equal(t, []int(nil), ret)
}

func TestFindCycle(t *testing.T) {
	gr := tests.NewTestGraph(5,
		0, 1,
		0, 2,
		1, 3,
		1, 4,
	)
	assert.Equal(t, []int(nil), FindCycle(gr))

	gr.AddEdge(2, 3)
	assert.Equal(t, []int{2, 0, 1, 3, 2}, FindCycle(gr))

	gr.AddEdge(3, 4)
	assert.Equal(t, []int{2, 0, 1, 3, 2}, FindCycle(gr))
}
