package digraph

import (
	"algorithms/graph/graph"
	"sort"
)

// TransitiveClosure is a transitive closure of a digraph.
// Transitive closure of a digraph is another digraph with the same set of vertices
// but with an edge from "v" to "w" iif "w" is reachable from "v".
type TransitiveClosure struct {
	graph.Graph
}

// Reachable tells if there is a directed path between vertices.
func (transitiveClosure TransitiveClosure) Reachable(fromVertexID int, toVertexID int) bool {
	vertexAdjacency := transitiveClosure.AdjacentVertices(fromVertexID)
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
	return TransitiveClosure{graph.NewImplGraph(dgr.NumVertices(), numEdges, adjacency)}
}
