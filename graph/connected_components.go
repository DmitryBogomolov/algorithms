package graph

// ConnectedComponents represents connected components of a graph.
type ConnectedComponents struct {
	count      int
	components []int
}

// Count returns number of connected components.
func (cc ConnectedComponents) Count() int {
	return cc.count
}

// Connected tells if two vertices are connected.
func (cc ConnectedComponents) Connected(vertex1 int, vertex2 int) bool {
	return cc.components[vertex1] == cc.components[vertex2]
}

// ComponentID returns the component to which vertex belongs.
func (cc ConnectedComponents) ComponentID(vertex int) int {
	return cc.components[vertex]
}

// Component returns vertices of connected component.
func (cc ConnectedComponents) Component(component int) []int {
	var ret []int
	for i := 0; i < len(cc.components); i++ {
		if cc.components[i] == component {
			ret = append(ret, i)
		}
	}
	return ret
}

func findConnectedComponentsCore(cc *ConnectedComponents, marked []bool, graph Graph, current int) {
	marked[current] = true
	cc.components[current] = cc.count
	for _, child := range graph.AdjacentVertices(current) {
		if !marked[child] {
			findConnectedComponentsCore(cc, marked, graph, child)
		}
	}
}

// FindConnectedComponents finds connected components in a graph.
func FindConnectedComponents(graph Graph) ConnectedComponents {
	numVertices := graph.NumVertices()
	result := ConnectedComponents{
		components: make([]int, numVertices),
	}
	marked := make([]bool, numVertices)
	for v := 0; v < numVertices; v++ {
		if !marked[v] {
			findConnectedComponentsCore(&result, marked, graph, v)
			result.count++
		}
	}
	return result
}
