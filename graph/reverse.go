package graph

type reversedDigraph struct {
	Graph
	adjacency [][]int
}

// ReversibleDigraph represents reversible digraph.
type ReversibleDigraph interface {
	Graph
	Reverse() Graph
}

func (rd reversedDigraph) AdjacentVertices(vertex int) []int {
	return rd.adjacency[vertex]
}

func (rd reversedDigraph) Reverse() Graph {
	return rd.Graph
}

func reverseDigraph(digraph Graph) ReversibleDigraph {
	numVertices := digraph.NumVertices()
	adjacency := make([][]int, numVertices)
	for v := 0; v < numVertices; v++ {
		for _, w := range digraph.AdjacentVertices(v) {
			adjacency[w] = append(adjacency[w], v)
		}
	}
	return reversedDigraph{
		Graph:     digraph,
		adjacency: adjacency,
	}
}

// ReverseDigraph builts reversed digraph for a digraph.
func ReverseDigraph(digraph Graph) Graph {
	reversible, ok := digraph.(ReversibleDigraph)
	if ok {
		return reversible.Reverse()
	}
	return reverseDigraph(digraph)
}
