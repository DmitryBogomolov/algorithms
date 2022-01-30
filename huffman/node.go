package huffman

type _Node struct {
	item         byte
	frequency    int
	lNode, rNode *_Node
}

func (n _Node) isLeaf() bool {
	return n.lNode == nil && n.rNode == nil
}
