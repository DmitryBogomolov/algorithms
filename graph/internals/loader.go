package internals

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimRight(line, "\n"), nil
}

func loadGraph(filename string) (g *TestGraph, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	ret := TestGraph{}
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
		ret.AddEdge(v1, v2)
	}
	g = &ret
	return
}
