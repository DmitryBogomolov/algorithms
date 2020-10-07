package huffman

type byteCodeTable map[byte]string

func newByteCodeTable() byteCodeTable {
	return byteCodeTable{}
}

func (tbl byteCodeTable) get(b byte) string {
	return tbl[b]
}

func (tbl byteCodeTable) set(b byte, code string) {
	tbl[b] = code
}
