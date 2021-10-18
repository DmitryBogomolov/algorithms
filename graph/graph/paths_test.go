package graph

import (
	"testing"

	"github.com/DmitryBogomolov/algorithms/graph/internals/tests"

	"github.com/stretchr/testify/assert"
)

func TestFindPathsDepthFirst(t *testing.T) {
	target := tests.NewTestGraph(7,
		0, 5,
		2, 4,
		2, 3,
		1, 2,
		0, 1,
		3, 4,
		3, 5,
		0, 2,
	)

	t.Run("Count", func(t *testing.T) {
		var ret Paths

		ret = FindPathsDepthFirst(target, 0)
		assert.Equal(t, ret.VertexCount(), 5, "vertex 0")

		ret = FindPathsDepthFirst(target, 5)
		assert.Equal(t, ret.VertexCount(), 5, "vertex 5")

		ret = FindPathsDepthFirst(target, 6)
		assert.Equal(t, ret.VertexCount(), 0, "vertex 6")
	})

	t.Run("Marked", func(t *testing.T) {
		var ret Paths

		allMarked := func() []bool {
			all := make([]bool, target.NumVertices())
			for i := 0; i < target.NumVertices(); i++ {
				all[i] = ret.HasPathTo(i)
			}
			return all
		}

		ret = FindPathsDepthFirst(target, 0)
		assert.Equal(t,
			[]bool{true, true, true, true, true, true, false},
			allMarked(),
			"vertex 0",
		)

		ret = FindPathsDepthFirst(target, 5)
		assert.Equal(t,
			[]bool{true, true, true, true, true, true, false},
			allMarked(),
			"vertex 5",
		)

		ret = FindPathsDepthFirst(target, 6)
		assert.Equal(t,
			[]bool{false, false, false, false, false, false, true},
			allMarked(),
			"vertex 6",
		)
	})

	t.Run("PathTo", func(t *testing.T) {
		var ret Paths

		ret = FindPathsDepthFirst(target, 0)
		assert.Equal(t, []int{0}, ret.PathTo(0), "0 -> 0")
		assert.Equal(t, []int{0, 5, 3, 2, 1}, ret.PathTo(1), "0 -> 1")
		assert.Equal(t, []int{0, 5, 3, 2, 4}, ret.PathTo(4), "0 -> 4")

		ret = FindPathsDepthFirst(target, 2)
		assert.Equal(t, []int{2}, ret.PathTo(2), "2 -> 2")
		assert.Equal(t, []int{2, 4, 3, 5, 0, 1}, ret.PathTo(1), "2 -> 1")
		assert.Equal(t, []int(nil), ret.PathTo(6), "2 -> 6")

		ret = FindPathsDepthFirst(target, 6)
		assert.Equal(t, []int{6}, ret.PathTo(6), "6 -> 6")
		assert.Equal(t, []int(nil), ret.PathTo(5), "6 -> 5")
	})
}

func TestFindPathsBreadthFirst(t *testing.T) {
	target := tests.NewTestGraph(7,
		0, 5,
		2, 4,
		2, 3,
		1, 2,
		0, 1,
		3, 4,
		3, 5,
		0, 2,
	)

	t.Run("Count", func(t *testing.T) {
		var ret Paths

		ret = FindPathsBreadthFirst(target, 0)
		assert.Equal(t, ret.VertexCount(), 5, "vertex 0")

		ret = FindPathsBreadthFirst(target, 5)
		assert.Equal(t, ret.VertexCount(), 5, "vertex 5")

		ret = FindPathsBreadthFirst(target, 6)
		assert.Equal(t, ret.VertexCount(), 0, "vertex 6")
	})

	t.Run("Marked", func(t *testing.T) {
		var ret Paths

		allMarked := func() []bool {
			all := make([]bool, target.NumVertices())
			for v := 0; v < target.NumVertices(); v++ {
				all[v] = ret.HasPathTo(v)
			}
			return all
		}

		ret = FindPathsBreadthFirst(target, 0)
		assert.Equal(t,
			[]bool{true, true, true, true, true, true, false},
			allMarked(),
			"vertex 0",
		)

		ret = FindPathsBreadthFirst(target, 5)
		assert.Equal(t,
			[]bool{true, true, true, true, true, true, false},
			allMarked(),
			"vertex 5",
		)

		ret = FindPathsBreadthFirst(target, 6)
		assert.Equal(t,
			[]bool{false, false, false, false, false, false, true},
			allMarked(),
			"vertex 6",
		)
	})

	t.Run("PathTo", func(t *testing.T) {
		var ret Paths

		ret = FindPathsBreadthFirst(target, 0)
		assert.Equal(t, []int{0}, ret.PathTo(0), "0 -> 0")
		assert.Equal(t, []int{0, 1}, ret.PathTo(1), "0 -> 1")
		assert.Equal(t, []int{0, 2, 4}, ret.PathTo(4), "0 -> 4")

		ret = FindPathsBreadthFirst(target, 2)
		assert.Equal(t, []int{2}, ret.PathTo(2), "2 -> 2")
		assert.Equal(t, []int{2, 1}, ret.PathTo(1), "2 -> 1")
		assert.Equal(t, []int{2, 3, 5}, ret.PathTo(5), "2 -> 5")
		assert.Equal(t, []int(nil), ret.PathTo(6), "2 -> 6")

		ret = FindPathsBreadthFirst(target, 6)
		assert.Equal(t, []int{6}, ret.PathTo(6), "6 -> 6")
		assert.Equal(t, []int(nil), ret.PathTo(5), "6 -> 5")
	})
}
