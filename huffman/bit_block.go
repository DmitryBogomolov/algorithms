package huffman

type bitBlock struct {
	buffer []byte
	size   int
}

func getBytesBits(bits int) (int, int) {
	return bits / 8, bits % 8
}

func getBufferSize(bits int) int {
	byteCnt, bitCnt := getBytesBits(bits)
	if bitCnt > 0 {
		byteCnt++
	}
	return byteCnt
}

func (bb bitBlock) shift(step int) bitBlock {
	if step < 1 {
		panic("shift step must be positive")
	}
	size := bb.size + step
	buffer := make([]byte, getBufferSize(size))
	byteIdx, bitShift := getBytesBits(step)
	var residue byte
	for i := 0; i < len(bb.buffer); i++ {
		srcByte := bb.buffer[i]
		dstByte := (srcByte << bitShift) | residue
		residue = srcByte >> (8 - bitShift)
		buffer[byteIdx+i] = dstByte
	}
	if residue > 0 {
		buffer[byteIdx+len(bb.buffer)] = residue
	}
	return bitBlock{
		buffer: buffer,
		size:   size,
	}
}
