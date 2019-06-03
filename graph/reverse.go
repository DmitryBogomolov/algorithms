package graph

type reversedDigraph struct {
	original  Graph
	adjacency [][]int
}

func (rd reversedDigraph) NumVertices() int {
	return rd.original.NumVertices()
}

func (rd reversedDigraph) NumEdges() int {
	return rd.original.NumEdges()
}

func (rd reversedDigraph) AdjacentVertices(vertex int) []int {
	return rd.adjacency[vertex]
}

// Reverse builts reversed digraph for a digraph.
func Reverse(digraph Graph) Graph {
	numVertices := digraph.NumVertices()
	adjacency := make([][]int, numVertices)
	for v := 0; v < numVertices; v++ {
		for _, w := range digraph.AdjacentVertices(v) {
			adjacency[w] = append(adjacency[w], v)
		}
	}
	return reversedDigraph{
		original:  digraph,
		adjacency: adjacency,
	}
}
