package graph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/graph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func checkConnectedComponents(t *testing.T, cc ConnectedComponents, componentsData [][]int) {
	assert.Equal(t, len(componentsData), cc.Count())
	for i, componentVertices := range componentsData {
		assert.Equal(t, componentVertices, cc.Component(i))
		for _, vertexID := range componentVertices {
			assert.Equal(t, i, cc.ComponentID(vertexID))
		}
		for k := 0; k < len(componentVertices)-1; k++ {
			assert.Equal(t, true, cc.Connected(componentVertices[k], componentVertices[k+1]))
		}
	}
	for k := 0; k < len(componentsData)-1; k++ {
		assert.Equal(t, false, cc.Connected(componentsData[k][0], componentsData[k+1][0]))
	}
}

func TestFindConnectedComponents_EmptyGraph(t *testing.T) {
	gr := tests.NewTestGraph(0)

	ret := FindConnectedComponents(gr)
	checkConnectedComponents(t, ret, nil)
}

func TestFindConnectedComponents_NoEdges(t *testing.T) {
	gr := tests.NewTestGraph(4)

	ret := FindConnectedComponents(gr)
	checkConnectedComponents(t, ret, [][]int{
		{0},
		{1},
		{2},
		{3},
	})
}

func TestFindConnectedComponents(t *testing.T) {
	gr := tests.NewTestGraph(8,
		0, 1,
		1, 4,
		4, 7,
		7, 2,
		2, 0,
		1, 2,
		2, 4,
		5, 6,
	)

	ret := FindConnectedComponents(gr)
	checkConnectedComponents(t, ret, [][]int{
		{0, 1, 2, 4, 7},
		{3},
		{5, 6},
	})
}
