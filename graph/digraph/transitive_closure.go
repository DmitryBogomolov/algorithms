package digraph

import (
	"algorithms/graph/graph"
	"sort"
)

// TransitiveClosure is a transitive closure of a digraph.
// Transitive closure of a digraph is another digraph with the same set of vertices
// but with an edge from "v" to "w" iif "w" is reachable from "v".
type TransitiveClosure struct {
	numVertices int
	numEdges    int
	adjacency   [][]int
}

// NumVertices gets number of vertices.
func (transitiveClosure TransitiveClosure) NumVertices() int {
	return transitiveClosure.numVertices
}

// NumEdges gets number of edges.
func (transitiveClosure TransitiveClosure) NumEdges() int {
	return transitiveClosure.numEdges
}

// AdjacentVertices returns vertices adjacent to the vertex.
func (transitiveClosure TransitiveClosure) AdjacentVertices(vertexID int) []int {
	return transitiveClosure.adjacency[vertexID]
}

// Reachable tells if there is a directed path between vertices.
func (transitiveClosure TransitiveClosure) Reachable(fromVertexID int, toVertexID int) bool {
	vertexAdjacency := transitiveClosure.adjacency[fromVertexID]
	idx := sort.SearchInts(vertexAdjacency, toVertexID)
	return idx < len(vertexAdjacency) && vertexAdjacency[idx] == toVertexID
}

// BuildTransitiveClosure returns transitive closure of a digraph
// by running depth-first search from each vertex.
// https://algs4.cs.princeton.edu/42digraph/TransitiveClosure.java.html
func BuildTransitiveClosure(dgr graph.Graph) TransitiveClosure {
	numEdges := 0
	adjacency := make([][]int, dgr.NumVertices())
	for vertexID := 0; vertexID < dgr.NumVertices(); vertexID++ {
		paths := graph.FindPathsDepthFirst(dgr, vertexID)
		var vertexAdjacency []int
		for adjacentVertexID := 0; adjacentVertexID < dgr.NumVertices(); adjacentVertexID++ {
			if paths.HasPathTo(adjacentVertexID) {
				vertexAdjacency = append(vertexAdjacency, adjacentVertexID)
				numEdges++
			}
		}
		adjacency[vertexID] = vertexAdjacency
	}
	return TransitiveClosure{
		numVertices: dgr.NumVertices(),
		numEdges:    numEdges,
		adjacency:   adjacency,
	}
}
