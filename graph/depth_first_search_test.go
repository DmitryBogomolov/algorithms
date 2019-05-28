package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepthFirstSearch(t *testing.T) {
	graph := &testGraph{
		numVertices: 7,
		numEdges:    8,
		adjacency:   make([][]int, 7),
	}
	graph.addEdge(0, 5)
	graph.addEdge(2, 4)
	graph.addEdge(2, 3)
	graph.addEdge(1, 2)
	graph.addEdge(0, 1)
	graph.addEdge(3, 4)
	graph.addEdge(3, 5)
	graph.addEdge(0, 2)

	t.Run("Count", func(t *testing.T) {
		var ret DepthFirstSearchResult

		ret = DepthFirstSearch(graph, 0)
		assert.Equal(t, ret.Count(), 6, "vertex 0")

		ret = DepthFirstSearch(graph, 5)
		assert.Equal(t, ret.Count(), 6, "vertex 5")

		ret = DepthFirstSearch(graph, 6)
		assert.Equal(t, ret.Count(), 1, "vertex 6")
	})

	t.Run("Marked", func(t *testing.T) {
		var ret DepthFirstSearchResult

		allMarked := func() []bool {
			all := make([]bool, graph.NumVertices())
			for i := 0; i < graph.NumVertices(); i++ {
				all[i] = ret.Marked(i)
			}
			return all
		}

		ret = DepthFirstSearch(graph, 0)
		assert.Equal(t,
			[]bool{true, true, true, true, true, true, false},
			allMarked(),
			"vertex 0",
		)

		ret = DepthFirstSearch(graph, 5)
		assert.Equal(t,
			[]bool{true, true, true, true, true, true, false},
			allMarked(),
			"vertex 5",
		)

		ret = DepthFirstSearch(graph, 6)
		assert.Equal(t,
			[]bool{false, false, false, false, false, false, true},
			allMarked(),
			"vertex 6",
		)
	})

	t.Run("PathTo", func(t *testing.T) {
		var ret DepthFirstSearchResult

		ret = DepthFirstSearch(graph, 0)
		assert.Equal(t, []int{0}, ret.PathTo(0), "0 -> 0")
		assert.Equal(t, []int{0, 5, 3, 2, 1}, ret.PathTo(1), "0 -> 1")
		assert.Equal(t, []int{0, 5, 3, 2, 4}, ret.PathTo(4), "0 -> 4")

		ret = DepthFirstSearch(graph, 2)
		assert.Equal(t, []int{2}, ret.PathTo(2), "2 -> 2")
		assert.Equal(t, []int{2, 4, 3, 5, 0, 1}, ret.PathTo(1), "2 -> 1")
		assert.Equal(t, []int(nil), ret.PathTo(6), "2 -> 6")

		ret = DepthFirstSearch(graph, 6)
		assert.Equal(t, []int{6}, ret.PathTo(6), "6 -> 6")
		assert.Equal(t, []int(nil), ret.PathTo(5), "6 -> 5")
	})
}
