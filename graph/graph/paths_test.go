package graph_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/graph/graph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/tests"

	"github.com/stretchr/testify/assert"
)

func checkPaths(t *testing.T, paths Paths, source int, count int, pathsTo map[int][]int) {
	assert.Equal(t, source, paths.SourceVertex())
	assert.Equal(t, count, paths.VertexCount())
	assert.Equal(t, true, paths.HasPathTo(source))
	assert.Equal(t, []int{source}, paths.PathTo(source))
	for i, pathTo := range pathsTo {
		assert.Equal(t, pathTo != nil, paths.HasPathTo(i))
		assert.Equal(t, pathTo, paths.PathTo(i))
	}
}

func TestFindPathsDepthFirst_OneVertex(t *testing.T) {
	gr := tests.NewTestGraph(1)

	ret := FindPathsDepthFirst(gr, 0)
	checkPaths(t, ret, 0, 0, nil)
}

func TestFindPathsDepthFirst_NoEdges(t *testing.T) {
	gr := tests.NewTestGraph(3)

	ret := FindPathsDepthFirst(gr, 1)
	checkPaths(t, ret, 1, 0, map[int][]int{
		0: nil,
		2: nil,
	})
}

func TestFindPathsDepthFirst(t *testing.T) {
	gr := tests.NewTestGraph(7,
		0, 5,
		2, 4,
		2, 3,
		1, 2,
		0, 1,
		3, 4,
		3, 5,
		0, 2,
	)

	checkPaths(t, FindPathsDepthFirst(gr, 0), 0, 5, map[int][]int{
		1: {0, 5, 3, 2, 1},
		2: {0, 5, 3, 2},
		3: {0, 5, 3},
		4: {0, 5, 3, 2, 4},
		5: {0, 5},
		6: nil,
	})

	checkPaths(t, FindPathsDepthFirst(gr, 3), 3, 5, map[int][]int{
		1: {3, 2, 1},
		2: {3, 2},
		4: {3, 2, 4},
		5: {3, 2, 1, 0, 5},
		6: nil,
	})

	checkPaths(t, FindPathsDepthFirst(gr, 6), 6, 0, map[int][]int{})
}

func TestFindPathsBreadthFirst_OneVertex(t *testing.T) {
	gr := tests.NewTestGraph(1)

	ret := FindPathsBreadthFirst(gr, 0)
	checkPaths(t, ret, 0, 0, nil)
}

func TestFindPathsBreadthFirst_NoEdges(t *testing.T) {
	gr := tests.NewTestGraph(3)

	ret := FindPathsBreadthFirst(gr, 1)
	checkPaths(t, ret, 1, 0, map[int][]int{
		0: nil,
		2: nil,
	})
}

func TestFindPathsBreadthFirst(t *testing.T) {
	gr := tests.NewTestGraph(7,
		0, 5,
		2, 4,
		2, 3,
		1, 2,
		0, 1,
		3, 4,
		3, 5,
		0, 2,
	)

	checkPaths(t, FindPathsBreadthFirst(gr, 0), 0, 5, map[int][]int{
		1: {0, 1},
		2: {0, 2},
		3: {0, 5, 3},
		4: {0, 2, 4},
		5: {0, 5},
		6: nil,
	})

	checkPaths(t, FindPathsBreadthFirst(gr, 3), 3, 5, map[int][]int{
		1: {3, 2, 1},
		2: {3, 2},
		4: {3, 4},
		5: {3, 5},
		6: nil,
	})

	checkPaths(t, FindPathsBreadthFirst(gr, 6), 6, 0, map[int][]int{})
}
