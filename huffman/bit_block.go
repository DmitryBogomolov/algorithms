package huffman

type bitBlock struct {
	buffer []byte
	size   int
}

func getBytesBits(bits int) (int, int) {
	return bits / 8, bits % 8
}

func makeBuffer(bits int) []byte {
	byteCnt, bitCnt := getBytesBits(bits)
	if bitCnt > 0 {
		byteCnt++
	}
	return make([]byte, byteCnt)
}

func (bb *bitBlock) grow(bits int) {
	size := bb.size + bits
	buffer := makeBuffer(size)
	copy(buffer, bb.buffer)
	bb.buffer = buffer
	bb.size = size
}

func (bb *bitBlock) clone() bitBlock {
	buffer := makeBuffer(bb.size)
	copy(buffer, bb.buffer)
	return bitBlock{buffer, bb.size}
}

func (bb *bitBlock) align() {
	_, bits := getBytesBits(bb.size)
	if bits > 0 {
		bb.size += 8 - bits
	}
}

func (bb *bitBlock) append(block bitBlock) {
	byteIdx, bitShift := getBytesBits(bb.size)
	bb.grow(block.size)
	var residue byte
	for i := 0; i < len(block.buffer); i++ {
		srcByte := block.buffer[i]
		dstByte := (srcByte << bitShift) | residue
		residue = srcByte >> (8 - bitShift)
		bb.buffer[byteIdx+i] |= dstByte
	}
	if residue > 0 {
		bb.buffer[byteIdx+len(block.buffer)] = residue
	}
}

var bitBlock0 = bitBlock{size: 1, buffer: []byte{0}}
var bitBlock1 = bitBlock{size: 1, buffer: []byte{1}}

func (bb *bitBlock) appendBit(bit bool) {
	var block bitBlock
	if bit {
		block = bitBlock1
	} else {
		block = bitBlock0
	}
	bb.append(block)
}

func (bb *bitBlock) appendByte(bt byte) {
	var block bitBlock
	block.grow(8)
	block.buffer[0] = bt
	bb.append(block)
}
