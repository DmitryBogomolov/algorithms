package digraph

import "algorithms/graph/graph"

type _ReversedDigraph struct {
	graph.Graph
	adjacency [][]int
}

// ReversibleDigraph is a reversible digraph.
type ReversibleDigraph interface {
	graph.Graph
	Reverse() graph.Graph
}

func (obj _ReversedDigraph) AdjacentVertices(vertexID int) []int {
	return obj.adjacency[vertexID]
}

func (obj _ReversedDigraph) Reverse() graph.Graph {
	return obj.Graph
}

// ReverseDigraph builts reversed digraph.
func ReverseDigraph(digraph graph.Graph) graph.Graph {
	if reversible, ok := digraph.(ReversibleDigraph); ok {
		return reversible.Reverse()
	}

	adjacency := make([][]int, digraph.NumVertices())
	for vertexID := 0; vertexID < digraph.NumVertices(); vertexID++ {
		for _, adjacentVertexID := range digraph.AdjacentVertices(vertexID) {
			adjacency[adjacentVertexID] = append(adjacency[adjacentVertexID], vertexID)
		}
	}
	return _ReversedDigraph{
		Graph:     digraph,
		adjacency: adjacency,
	}
}
