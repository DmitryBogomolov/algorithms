package graph

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testGraph struct {
	numVertices int
	numEdges    int
	adjacency   [][]int
}

type testEdgeWeightedGraph struct {
	testGraph
	weights [][]float64
}

func newTestGraph(numVertices int, connections ...int) *testGraph {
	graph := testGraph{
		numVertices: numVertices,
		numEdges:    len(connections) / 2,
		adjacency:   make([][]int, numVertices),
	}
	for i := 0; i < len(connections)/2; i++ {
		graph.addEdge(connections[2*i], connections[2*i+1])
	}
	return &graph
}

func (g *testGraph) addEdge(v1, v2 int) {
	g.adjacency[v1] = append(g.adjacency[v1], v2)
	g.adjacency[v2] = append(g.adjacency[v2], v1)
}

func (g *testGraph) addDirectedEdge(v1, v2 int) {
	g.adjacency[v1] = append(g.adjacency[v1], v2)
}

func (g *testGraph) NumVertices() int {
	return g.numVertices
}

func (g *testGraph) NumEdges() int {
	return g.numEdges
}

func (g *testGraph) AdjacentVertices(vertex int) []int {
	return g.adjacency[vertex]
}

func newTestDigraph(numVertices int, connections ...int) *testGraph {
	digraph := testGraph{
		numVertices: numVertices,
		numEdges:    len(connections) / 2,
		adjacency:   make([][]int, numVertices),
	}
	for i := 0; i < len(connections)/2; i++ {
		digraph.addDirectedEdge(connections[2*i], connections[2*i+1])
	}
	return &digraph
}

type testWeightedEdge struct {
	v1, v2 int
	weight float64
}

func (g *testEdgeWeightedGraph) addEdge(v1, v2 int, w float64) {
	g.testGraph.addEdge(v1, v2)
	g.weights[v1] = append(g.weights[v1], w)
	g.weights[v2] = append(g.weights[v2], w)
}

func (g *testEdgeWeightedGraph) addDirectedEdge(v1, v2 int, w float64) {
	g.testGraph.addDirectedEdge(v1, v2)
	g.weights[v1] = append(g.weights[v1], w)
}

func (g *testEdgeWeightedGraph) AdjacentWeights(vertex int) []float64 {
	return g.weights[vertex]
}

func newTestEdgeWeightedGraph(numVertices int, edges []testWeightedEdge) *testEdgeWeightedGraph {
	graph := testEdgeWeightedGraph{
		testGraph: testGraph{
			numVertices: numVertices,
			numEdges:    len(edges),
			adjacency:   make([][]int, numVertices),
		},
		weights: make([][]float64, numVertices),
	}
	for _, edge := range edges {
		graph.addEdge(edge.v1, edge.v2, edge.weight)
	}
	return &graph
}

func newTestEdgeWeightedDigraph(numVertices int, edges []testWeightedEdge) *testEdgeWeightedGraph {
	digraph := testEdgeWeightedGraph{
		testGraph: testGraph{
			numVertices: numVertices,
			numEdges:    len(edges),
			adjacency:   make([][]int, numVertices),
		},
		weights: make([][]float64, numVertices),
	}
	for _, edge := range edges {
		digraph.addDirectedEdge(edge.v1, edge.v2, edge.weight)
	}
	return &digraph
}

func TestAllGraphEdges(t *testing.T) {
	graph := newTestGraph(6,
		0, 1,
		1, 2,
		0, 3,
		3, 2,
		4, 3,
		4, 5,
		5, 0,
	)

	ret := AllGraphEdges(graph)

	assert.Equal(t, []Edge{{0, 1}, {0, 3}, {0, 5}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}, ret)
}

func TestAllDigraphEdges(t *testing.T) {
	graph := newTestDigraph(6,
		0, 1,
		1, 2,
		0, 3,
		3, 0,
		4, 3,
		4, 5,
		5, 0,
	)

	ret := AllDigraphEdges(graph)

	assert.Equal(t, []Edge{{0, 1}, {0, 3}, {1, 2}, {3, 0}, {4, 3}, {4, 5}, {5, 0}}, ret)
}

func TestAllGraphWeights(t *testing.T) {
	graph := newTestEdgeWeightedGraph(6, []testWeightedEdge{
		{0, 1, 1.2},
		{1, 2, 2.3},
		{0, 3, 3.1},
		{3, 2, 4.1},
		{4, 3, 5.1},
		{4, 5, 1.6},
		{5, 0, 2.2},
	})

	ret := AllGraphWeights(graph)

	assert.Equal(t, []float64{1.2, 3.1, 2.2, 2.3, 4.1, 5.1, 1.6}, ret)
	assert.InDelta(t, 19.6, TotalGraphWeight(graph), 0.0001)
}

func TestAllDigraphWeights(t *testing.T) {
	graph := newTestEdgeWeightedDigraph(6, []testWeightedEdge{
		{0, 1, 1.2},
		{1, 2, 2.3},
		{0, 3, 3.1},
		{3, 0, 4.1},
		{4, 3, 5.1},
		{4, 5, 1.6},
		{5, 0, 2.2},
	})

	ret := AllDigraphWeights(graph)

	assert.Equal(t, []float64{1.2, 3.1, 2.3, 4.1, 5.1, 1.6, 2.2}, ret)
	assert.InDelta(t, 19.6, TotalGraphWeight(graph), 0.0001)
}

func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimRight(line, "\n"), nil
}

func loadGraph(filename string) (g *testGraph, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	ret := testGraph{}
	var line string
	line, err = readLine(reader)
	if err != nil {
		return
	}
	numVertices, err := strconv.Atoi(line)
	if err != nil {
		return
	}
	ret.numVertices = numVertices
	line, err = readLine(reader)
	if err != nil {
		return
	}
	numEdges, err := strconv.Atoi(line)
	if err != nil {
		return
	}
	ret.numEdges = numEdges
	ret.adjacency = make([][]int, numVertices)
	for {
		line, e := readLine(reader)
		if e == io.EOF {
			break
		}
		if e != nil {
			err = e
			return
		}
		parts := strings.Split(line, " ")
		v1, e1 := strconv.Atoi(parts[0])
		if e1 != nil {
			err = e1
			return
		}
		v2, e2 := strconv.Atoi(parts[1])
		if err != nil {
			err = e2
			return
		}
		ret.addEdge(v1, v2)
	}
	g = &ret
	return
}

func TestLoadGraph(t *testing.T) {
	const filename = "test-data.txt"
	ioutil.WriteFile(filename, []byte("7\n5\n0 1\n0 2\n1 2\n1 4\n3 4\n"), os.ModePerm)
	defer os.Remove(filename)

	g, err := loadGraph(filename)

	assert.NoError(t, err, "error")
	assert.Equal(t, g.NumVertices(), 7, "vertices")
	assert.Equal(t, g.NumEdges(), 5, "edges")
	assert.Equal(t, g.AdjacentVertices(0), []int{1, 2}, "vertex 0")
	assert.Equal(t, g.AdjacentVertices(1), []int{0, 2, 4}, "vertex 1")
	assert.Equal(t, g.AdjacentVertices(2), []int{0, 1}, "vertex 2")
	assert.Equal(t, g.AdjacentVertices(3), []int{4}, "vertex 3")
	assert.Equal(t, g.AdjacentVertices(4), []int{1, 3}, "vertex 4")
	assert.Equal(t, g.AdjacentVertices(5), []int(nil), "vertex 5")
	assert.Equal(t, g.AdjacentVertices(6), []int(nil), "vertex 6")
}
