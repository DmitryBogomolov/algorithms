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
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		for _, adjacentVertexID := range digraph.AdjacentVertices(vertexID) {
			adjacency[adjacentVertexID] = append(adjacency[adjacentVertexID], vertexID)
		}
	}
	return reversedDigraph{
		Graph:     digraph,
		adjacency: adjacency,
	}
}

// ReverseDigraph builts reversed digraph.
func ReverseDigraph(digraph Graph) Graph {
	reversible, ok := digraph.(ReversibleDigraph)
	if ok {
		return reversible.Reverse()
	}
	return reverseDigraph(digraph)
}
