package tests

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
