package graph

// ConnectedComponents is a collection of connected components in a graph.
// Connected component is a set of vertices connected by edges.
type ConnectedComponents struct {
	count      int
	components []int
}

// Count gets number of connected components.
func (connectedComponents ConnectedComponents) Count() int {
	return connectedComponents.count
}

// Connected tells if two vertices are connected.
func (connectedComponents ConnectedComponents) Connected(vertexID1 int, vertexID2 int) bool {
	return connectedComponents.ComponentID(vertexID1) == connectedComponents.ComponentID(vertexID2)
}

// ComponentID returns component to which vertex belongs.
func (connectedComponents ConnectedComponents) ComponentID(vertexID int) int {
	return connectedComponents.components[vertexID]
}

// Component returns vertices of a connected component.
func (connectedComponents ConnectedComponents) Component(componentID int) []int {
	var vertices []int
	for i := 0; i < len(connectedComponents.components); i++ {
		if connectedComponents.ComponentID(i) == componentID {
			vertices = append(vertices, i)
		}
	}
	return vertices
}

func newConnectedComponents(numVertices int) ConnectedComponents {
	return ConnectedComponents{
		components: make([]int, numVertices),
	}
}

func findConnectedComponentsCore(
	connectedComponents *ConnectedComponents, marked []bool, graph Graph, vertexID int,
) {
	marked[vertexID] = true
	connectedComponents.components[vertexID] = connectedComponents.count
	for _, adjacentVertexID := range graph.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			findConnectedComponentsCore(connectedComponents, marked, graph, adjacentVertexID)
		}
	}
}

// FindConnectedComponents returns connected components in a graph.
// https://algs4.cs.princeton.edu/41graph/CC.java.html
func FindConnectedComponents(graph Graph) ConnectedComponents {
	numVertices := graph.NumVertices()
	connectedComponents := newConnectedComponents(numVertices)
	marked := make([]bool, numVertices)
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		if !marked[vertexID] {
			findConnectedComponentsCore(&connectedComponents, marked, graph, vertexID)
			connectedComponents.count++
		}
	}
	return connectedComponents
}
