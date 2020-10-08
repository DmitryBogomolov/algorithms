package huffman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeBitBlock(size int, bytes ...byte) bitBlock {
	return bitBlock{buffer: bytes, size: size}
}

func TestBitBlockShift(t *testing.T) {
	block := makeBitBlock(4, 0b00001001)

	assert.Equal(t,
		makeBitBlock(5, 0b00010010),
		block.shift(1),
	)
	assert.Equal(t,
		makeBitBlock(6, 0b00100100),
		block.shift(2),
	)
	assert.Equal(t,
		makeBitBlock(8, 0b10010000),
		block.shift(4),
	)
	assert.Equal(t,
		makeBitBlock(9, 0b00100000, 0b00000001),
		block.shift(5),
	)
	assert.Equal(t,
		makeBitBlock(16, 0, 0b10010000),
		block.shift(12),
	)
	assert.Equal(t,
		makeBitBlock(20, 0, 0, 0b00001001),
		block.shift(16),
	)
	assert.Equal(t,
		makeBitBlock(23, 0, 0, 0b01001000),
		block.shift(19),
	)
}
