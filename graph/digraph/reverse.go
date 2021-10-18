package digraph

import "github.com/DmitryBogomolov/algorithms/graph/graph"

type _ReversedDigraph struct {
	graph.Graph
	adjacency [][]int
}

// ReversibleDigraph is a reversible digraph.
type ReversibleDigraph interface {
	graph.Graph
	Reverse() graph.Graph
}

func (dgr _ReversedDigraph) AdjacentVertices(vertexID int) []int {
	return dgr.adjacency[vertexID]
}

func (dgr _ReversedDigraph) Reverse() graph.Graph {
	return dgr.Graph
}

// ReverseDigraph builts reversed digraph.
func ReverseDigraph(dgr graph.Graph) graph.Graph {
	if reversible, ok := dgr.(ReversibleDigraph); ok {
		return reversible.Reverse()
	}

	adjacency := make([][]int, dgr.NumVertices())
	for vertexID := 0; vertexID < dgr.NumVertices(); vertexID++ {
		for _, adjacentVertexID := range dgr.AdjacentVertices(vertexID) {
			adjacency[adjacentVertexID] = append(adjacency[adjacentVertexID], vertexID)
		}
	}
	return _ReversedDigraph{
		Graph:     dgr,
		adjacency: adjacency,
	}
}
