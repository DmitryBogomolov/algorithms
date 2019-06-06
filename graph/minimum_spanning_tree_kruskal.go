package graph

type unionFind struct {
	parent []int
	rank   []byte
}

func (uf *unionFind) find(p int) int {
	i := p
	parent := uf.parent
	for i != parent[i] {
		i = parent[i]
	}
	return i
}

func (uf *unionFind) union(p, q int) {
	pRoot, qRoot := uf.find(p), uf.find(q)
	if pRoot == qRoot {
		return
	}
	pRank, qRank := uf.rank[pRoot], uf.rank[qRoot]
	if pRank < qRank {
		uf.parent[pRoot] = qRoot
	} else {
		uf.parent[qRoot] = pRoot
	}
	if pRank == qRank {
		uf.rank[pRoot]++
	}
}

func (uf *unionFind) connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func newUnionFind(size int) *unionFind {
	parent := make([]int, size)
	rank := make([]byte, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &unionFind{
		parent: parent,
		rank:   rank,
	}
}

// MinimumSpanningTreeKruskal computes minimum spanning tree for a graph using Kruskal's algorithm.
func MinimumSpanningTreeKruskal(graph EdgeWeightedGraph) EdgeWeightedGraph {
	return nil
}
