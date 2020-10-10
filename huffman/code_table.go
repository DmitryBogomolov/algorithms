package huffman

type byteCodeTable map[byte]*bitBlock

func newByteCodeTable() byteCodeTable {
	return byteCodeTable{}
}

func (tbl byteCodeTable) get(b byte) *bitBlock {
	return tbl[b]
}

func (tbl byteCodeTable) set(b byte, code *bitBlock) {
	tbl[b] = code
}
