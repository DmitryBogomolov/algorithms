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

func (g *testGraph) addEdge(v1, v2 int) {
	g.adjacency[v1] = append(g.adjacency[v1], v2)
	g.adjacency[v2] = append(g.adjacency[v2], v1)
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
