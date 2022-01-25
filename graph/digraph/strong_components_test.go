package digraph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/digraph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func checkStrongComponents(t *testing.T, sc StrongComponents, componentsData [][]int) {
	assert.Equal(t, len(componentsData), sc.Count())
	for i, componentVertices := range componentsData {
		assert.Equal(t, componentVertices, sc.Component(i))
		for _, vertexID := range componentVertices {
			assert.Equal(t, i, sc.ComponentID(vertexID))
		}
		for k := 0; k < len(componentVertices)-1; k++ {
			assert.Equal(t, true, sc.Connected(componentVertices[k], componentVertices[k+1]))
		}
	}
	for k := 0; k < len(componentsData)-1; k++ {
		assert.Equal(t, false, sc.Connected(componentsData[k][0], componentsData[k+1][0]))
	}
}

func TestFindStrongComponents_EmptyGraph(t *testing.T) {
	gr := tests.NewTestDigraph(0)

	ret := FindStrongComponents(gr)
	checkStrongComponents(t, ret, nil)
}

func TestFindStrongComponents_NoEdges(t *testing.T) {
	gr := tests.NewTestDigraph(4)

	ret := FindStrongComponents(gr)
	checkStrongComponents(t, ret, [][]int{
		{3},
		{2},
		{1},
		{0},
	})
}

func TestFindStrongComponents(t *testing.T) {
	gr := tests.NewTestDigraph(13,
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

	ret := FindStrongComponents(gr)
	checkStrongComponents(t, ret, [][]int{
		{1},
		{0, 2, 3, 4, 5},
		{9, 10, 11, 12},
		{6, 8},
		{7},
	})
}
