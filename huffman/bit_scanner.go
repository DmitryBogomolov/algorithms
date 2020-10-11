package huffman

type bitScanner struct {
	buffer     []byte
	byteOffset int
	bitOffset  int
}

func newBitScanner(buffer []byte) *bitScanner {
	return &bitScanner{buffer: buffer}
}

func (bs *bitScanner) readBit() bool {
	bit := (bs.buffer[bs.byteOffset] & (1 << bs.bitOffset)) > 0
	bs.bitOffset++
	if bs.bitOffset == 8 {
		bs.byteOffset++
		bs.bitOffset = 0
	}
	return bit
}

func (bs *bitScanner) align() {
	if bs.bitOffset > 0 {
		bs.byteOffset++
		bs.bitOffset = 0
	}
}

func (bs *bitScanner) readByte() byte {
	b := bs.buffer[bs.byteOffset]
	if bs.bitOffset > 0 {
		b1 := b >> bs.bitOffset
		b2 := bs.buffer[bs.byteOffset+1] << (8 - bs.bitOffset)
		b = b1 | b2
	}
	bs.byteOffset++
	return b
}
