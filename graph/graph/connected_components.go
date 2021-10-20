package graph

import (
	"fmt"

	"github.com/DmitryBogomolov/algorithms/graph/internal/utils"
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
	gr Graph, marked []bool, componentCount *int, componentIDs []int,
	vertexID int,
) {
	marked[vertexID] = true
	componentIDs[vertexID] = *componentCount
	for _, adjacentVertexID := range gr.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			findConnectedComponentsCore(gr, marked, componentCount, componentIDs, adjacentVertexID)
		}
	}
}

// FindConnectedComponents returns connected components in a graph.
// https://algs4.cs.princeton.edu/41graph/CC.java.html
func FindConnectedComponents(gr Graph) ConnectedComponents {
	componentCount := 0
	componentIDs := make([]int, gr.NumVertices())
	utils.ResetList(componentIDs)
	marked := make([]bool, gr.NumVertices())
	for vertexID := 0; vertexID < gr.NumVertices(); vertexID++ {
		if !marked[vertexID] {
			findConnectedComponentsCore(gr, marked, &componentCount, componentIDs, vertexID)
			componentCount++
		}
	}
	return NewConnectedComponents(componentCount, componentIDs)
}
