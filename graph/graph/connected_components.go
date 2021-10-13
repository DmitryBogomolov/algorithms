package graph

import (
	"algorithms/graph/internals"
	"fmt"
)

// ConnectedComponents is a collection of connected components in a graph.
// Connected component is a set of vertices connected by edges.
type ConnectedComponents struct {
	count        int
	componentIDs []int
}

// Count gets number of connected components.
func (connectedComponents ConnectedComponents) Count() int {
	return connectedComponents.count
}

// ComponentID returns component to which vertex belongs.
func (connectedComponents ConnectedComponents) ComponentID(vertexID int) int {
	if vertexID < 0 || vertexID > len(connectedComponents.componentIDs)-1 {
		panic(fmt.Sprintf("vertex '%d' is out of range", vertexID))
	}
	return connectedComponents.componentIDs[vertexID]
}

// Connected tells if two vertices are connected.
func (connectedComponents ConnectedComponents) Connected(vertexID1 int, vertexID2 int) bool {
	return connectedComponents.ComponentID(vertexID1) == connectedComponents.ComponentID(vertexID2)
}

// Component returns vertices of a connected component.
func (connectedComponents ConnectedComponents) Component(componentID int) []int {
	if componentID < 0 || componentID > connectedComponents.count-1 {
		panic(fmt.Sprintf("component '%d' is out of range", componentID))
	}
	var vertices []int
	for vertexID := 0; vertexID < len(connectedComponents.componentIDs); vertexID++ {
		if connectedComponents.componentIDs[vertexID] == componentID {
			vertices = append(vertices, vertexID)
		}
	}
	return vertices
}

// NewConnectedComponents creates instance.
func NewConnectedComponents(count int, componentIDs []int) ConnectedComponents {
	return ConnectedComponents{count: count, componentIDs: componentIDs}
}

func findConnectedComponentsCore(
	componentCount *int, componentIDs []int,
	graph Graph, marked []bool, vertexID int,
) {
	marked[vertexID] = true
	componentIDs[vertexID] = *componentCount
	for _, adjacentVertexID := range graph.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			findConnectedComponentsCore(componentCount, componentIDs, graph, marked, adjacentVertexID)
		}
	}
}

// FindConnectedComponents returns connected components in a graph.
// https://algs4.cs.princeton.edu/41graph/CC.java.html
func FindConnectedComponents(graph Graph) ConnectedComponents {
	componentCount := 0
	componentIDs := make([]int, graph.NumVertices())
	internals.ResetList(componentIDs)
	marked := make([]bool, graph.NumVertices())
	for vertexID := 0; vertexID < graph.NumVertices(); vertexID++ {
		if !marked[vertexID] {
			findConnectedComponentsCore(&componentCount, componentIDs, graph, marked, vertexID)
			componentCount++
		}
	}
	return NewConnectedComponents(componentCount, componentIDs)
}
