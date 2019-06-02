package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPathsDepthFirst(t *testing.T) {
	graph := newTestGraph(7,
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
		var ret VertexPaths

		ret = FindPathsDepthFirst(graph, 0)
		assert.Equal(t, ret.Count(), 6, "vertex 0")

		ret = FindPathsDepthFirst(graph, 5)
		assert.Equal(t, ret.Count(), 6, "vertex 5")

		ret = FindPathsDepthFirst(graph, 6)
		assert.Equal(t, ret.Count(), 1, "vertex 6")
	})

	t.Run("Marked", func(t *testing.T) {
		var ret VertexPaths

		allMarked := func() []bool {
			all := make([]bool, graph.NumVertices())
			for i := 0; i < graph.NumVertices(); i++ {
				all[i] = ret.HasPathTo(i)
			}
			return all
		}

		ret = FindPathsDepthFirst(graph, 0)
		assert.Equal(t,
			[]bool{true, true, true, true, true, true, false},
			allMarked(),
			"vertex 0",
		)

		ret = FindPathsDepthFirst(graph, 5)
		assert.Equal(t,
			[]bool{true, true, true, true, true, true, false},
			allMarked(),
			"vertex 5",
		)

		ret = FindPathsDepthFirst(graph, 6)
		assert.Equal(t,
			[]bool{false, false, false, false, false, false, true},
			allMarked(),
			"vertex 6",
		)
	})

	t.Run("PathTo", func(t *testing.T) {
		var ret VertexPaths

		ret = FindPathsDepthFirst(graph, 0)
		assert.Equal(t, []int{0}, ret.PathTo(0), "0 -> 0")
		assert.Equal(t, []int{0, 5, 3, 2, 1}, ret.PathTo(1), "0 -> 1")
		assert.Equal(t, []int{0, 5, 3, 2, 4}, ret.PathTo(4), "0 -> 4")

		ret = FindPathsDepthFirst(graph, 2)
		assert.Equal(t, []int{2}, ret.PathTo(2), "2 -> 2")
		assert.Equal(t, []int{2, 4, 3, 5, 0, 1}, ret.PathTo(1), "2 -> 1")
		assert.Equal(t, []int(nil), ret.PathTo(6), "2 -> 6")

		ret = FindPathsDepthFirst(graph, 6)
		assert.Equal(t, []int{6}, ret.PathTo(6), "6 -> 6")
		assert.Equal(t, []int(nil), ret.PathTo(5), "6 -> 5")
	})
}

func TestFindPathsBreadthFirst(t *testing.T) {
	graph := newTestGraph(7,
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
		var ret VertexPaths

		ret = FindPathsBreadthFirst(graph, 0)
		assert.Equal(t, ret.Count(), 6, "vertex 0")

		ret = FindPathsBreadthFirst(graph, 5)
		assert.Equal(t, ret.Count(), 6, "vertex 5")

		ret = FindPathsBreadthFirst(graph, 6)
		assert.Equal(t, ret.Count(), 1, "vertex 6")
	})

	t.Run("Marked", func(t *testing.T) {
		var ret VertexPaths

		allMarked := func() []bool {
			all := make([]bool, graph.NumVertices())
			for v := 0; v < graph.NumVertices(); v++ {
				all[v] = ret.HasPathTo(v)
			}
			return all
		}

		ret = FindPathsBreadthFirst(graph, 0)
		assert.Equal(t,
			[]bool{true, true, true, true, true, true, false},
			allMarked(),
			"vertex 0",
		)

		ret = FindPathsBreadthFirst(graph, 5)
		assert.Equal(t,
			[]bool{true, true, true, true, true, true, false},
			allMarked(),
			"vertex 5",
		)

		ret = FindPathsBreadthFirst(graph, 6)
		assert.Equal(t,
			[]bool{false, false, false, false, false, false, true},
			allMarked(),
			"vertex 6",
		)
	})

	t.Run("PathTo", func(t *testing.T) {
		var ret VertexPaths

		ret = FindPathsBreadthFirst(graph, 0)
		assert.Equal(t, []int{0}, ret.PathTo(0), "0 -> 0")
		assert.Equal(t, []int{0, 1}, ret.PathTo(1), "0 -> 1")
		assert.Equal(t, []int{0, 2, 4}, ret.PathTo(4), "0 -> 4")

		ret = FindPathsBreadthFirst(graph, 2)
		assert.Equal(t, []int{2}, ret.PathTo(2), "2 -> 2")
		assert.Equal(t, []int{2, 1}, ret.PathTo(1), "2 -> 1")
		assert.Equal(t, []int{2, 3, 5}, ret.PathTo(5), "2 -> 5")
		assert.Equal(t, []int(nil), ret.PathTo(6), "2 -> 6")

		ret = FindPathsBreadthFirst(graph, 6)
		assert.Equal(t, []int{6}, ret.PathTo(6), "6 -> 6")
		assert.Equal(t, []int(nil), ret.PathTo(5), "6 -> 5")
	})
}

func TestFindConnectedComponents(t *testing.T) {
	graph := newTestGraph(8,
		0, 1,
		1, 4,
		4, 7,
		7, 2,
		2, 0,
		1, 2,
		2, 4,
		5, 6,
	)

	ret := FindConnectedComponents(graph)

	assert.Equal(t, 3, ret.Count(), "count)")

	components := make([]int, graph.NumVertices())
	for v := 0; v < graph.NumVertices(); v++ {
		components[v] = ret.ComponentID(v)
	}
	assert.Equal(t,
		[]int{0, 0, 0, 1, 0, 2, 2, 0},
		components,
		"components",
	)

	assert.Equal(t, true, ret.Connected(0, 7), "0 - 7")
	assert.Equal(t, false, ret.Connected(2, 5), "2 - 5")
	assert.Equal(t, false, ret.Connected(4, 3), "4 - 3")
	assert.Equal(t, true, ret.Connected(3, 3), "3 - 3")
	assert.Equal(t, true, ret.Connected(5, 6), "5 - 6")
	assert.Equal(t, true, ret.Connected(7, 1), "7 - 1")
	assert.Equal(t, true, ret.Connected(2, 4), "2 - 4")
}

func TestHasCycle(t *testing.T) {
	var graph *testGraph

	graph = newTestGraph(3,
		0, 1,
		1, 2,
		2, 0,
	)
	assert.Equal(t, true, HasCycle(graph))

	graph = newTestGraph(3,
		0, 1,
		1, 2,
	)
	assert.Equal(t, false, HasCycle(graph))

	graph = newTestGraph(5,
		0, 1,
		0, 2,
		1, 3,
		1, 4,
	)
	assert.Equal(t, false, HasCycle(graph))

	graph = newTestGraph(5,
		0, 1,
		0, 2,
		1, 3,
		1, 4,
		2, 3,
		3, 4,
	)
	assert.Equal(t, true, HasCycle(graph))
}

func TestIsBipartite(t *testing.T) {
	var graph *testGraph

	graph = newTestGraph(6,
		0, 1,
		0, 3,
		1, 2,
		2, 3,
		1, 4,
		2, 5,
		4, 5,
	)
	assert.Equal(t, true, IsBipartite(graph))

	graph = newTestGraph(6,
		0, 1,
		0, 3,
		1, 2,
		2, 3,
		2, 4,
		1, 4,
		2, 5,
		4, 5,
	)
	assert.Equal(t, false, IsBipartite(graph))
}

func TestFindCutEdges(t *testing.T) {
	var graph *testGraph
	var ret []CutEdge

	graph = newTestGraph(7,
		0, 1,
		0, 2,
		1, 3,
		2, 3,
		3, 4,
		4, 5,
		4, 6,
	)

	ret = FindCutEdges(graph)
	assert.Equal(t, []CutEdge{CutEdge{4, 5}, CutEdge{4, 6}, CutEdge{3, 4}}, ret)

	graph.addEdge(5, 6)

	ret = FindCutEdges(graph)
	assert.Equal(t, []CutEdge{CutEdge{3, 4}}, ret)

	graph.addEdge(1, 5)

	ret = FindCutEdges(graph)
	assert.Equal(t, []CutEdge(nil), ret)
}

func TestFindCutVertices(t *testing.T) {
	var graph *testGraph
	var ret []int

	graph = newTestGraph(7,
		0, 1,
		0, 2,
		1, 3,
		2, 3,
		3, 4,
		4, 5,
		4, 6,
	)

	ret = FindCutVertices(graph)
	assert.Equal(t, []int{3, 4}, ret)

	graph.addEdge(1, 5)

	ret = FindCutVertices(graph)
	assert.Equal(t, []int{4}, ret)

	graph.addEdge(2, 6)

	ret = FindCutVertices(graph)
	assert.Equal(t, []int(nil), ret)
}
