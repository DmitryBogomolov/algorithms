package unionfind

// UnionFind is a union-find structure.
type UnionFind struct {
	count  int
	parent []int
	rank   []byte
}

// Count gets number of sets.
func (uf *UnionFind) Count() int {
	return uf.count
}

// Find returns set identifier for "p".
func (uf *UnionFind) Find(p int) int {
	if p < 0 || p > len(uf.parent)-1 {
		panic("out of range")
	}
	i := p
	parent := uf.parent
	for i != parent[i] {
		parent[i] = parent[parent[i]]
		i = parent[i]
	}
	return i
}

// Union adds connection between "p" and "q".
func (uf *UnionFind) Union(p, q int) {
	pRoot, qRoot := uf.Find(p), uf.Find(q)
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
	uf.count--
}

// Connected tells if "p" and "q" are connected (i.e. in the same set).
func (uf *UnionFind) Connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// New creates instance of UnionFind.
func New(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]byte, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &UnionFind{
		count:  size,
		parent: parent,
		rank:   rank,
	}
}
