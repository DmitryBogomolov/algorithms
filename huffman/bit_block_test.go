package huffman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeBitBlock(size int, bytes ...byte) bitBlock {
	return bitBlock{buffer: bytes, size: size}
}

func TestBitBlockGrow(t *testing.T) {
	block := makeBitBlock(4, 0b00001001)

	block.grow(1)
	assert.Equal(t,
		makeBitBlock(5, 0b00001001),
		block,
	)

	block.grow(3)
	assert.Equal(t,
		makeBitBlock(8, 0b00001001),
		block,
	)

	block.grow(2)
	assert.Equal(t,
		makeBitBlock(10, 0b00001001, 0),
		block,
	)

	block.grow(32)
	assert.Equal(t,
		makeBitBlock(42, 0b00001001, 0, 0, 0, 0, 0),
		block,
	)
}

func TestBitBlockAlign(t *testing.T) {
	check := func(block, expected bitBlock) {
		block.align()
		assert.Equal(t, expected, block)
	}

	check(
		makeBitBlock(3, 0b00000001),
		makeBitBlock(8, 0b00000001),
	)

	check(
		makeBitBlock(8, 0b01000001),
		makeBitBlock(8, 0b01000001),
	)

	check(
		makeBitBlock(14, 10, 0b00101100),
		makeBitBlock(16, 10, 0b00101100),
	)
}

func TestBitBlockAppend(t *testing.T) {
	check := func(block, appendee, expected bitBlock) {
		block.append(appendee)
		assert.Equal(t, expected, block)
	}

	check(
		makeBitBlock(4, 0b00001001),
		makeBitBlock(1, 0b00000001),
		makeBitBlock(5, 0b00011001),
	)

	check(
		makeBitBlock(4, 0b00000101),
		makeBitBlock(2, 0b00000011),
		makeBitBlock(6, 0b00110101),
	)

	check(
		makeBitBlock(4, 0b00001011),
		makeBitBlock(4, 0b00001100),
		makeBitBlock(8, 0b11001011),
	)

	check(
		makeBitBlock(4, 0b00001100),
		makeBitBlock(8, 0b10011001),
		makeBitBlock(12, 0b10011100, 0b00001001),
	)

	check(
		makeBitBlock(4, 0b00001000),
		makeBitBlock(32, 0, 0b10010001, 0, 0b11001100),
		makeBitBlock(36, 0b00001000, 0b00010000, 0b00001001, 0b11000000, 0b00001100),
	)

	check(
		makeBitBlock(19, 10, 20, 0b00000101),
		makeBitBlock(11, 0b11000111, 0b00000101),
		makeBitBlock(30, 10, 20, 0b00111101, 0b00101110),
	)
}

func TestBitBlockAppendBit(t *testing.T) {
	check := func(block bitBlock, b bool, expected bitBlock) {
		block.appendBit(b)
		assert.Equal(t, expected, block)
	}

	check(
		makeBitBlock(4, 0b00001100),
		false,
		makeBitBlock(5, 0b00001100),
	)

	check(
		makeBitBlock(4, 0b00001100),
		true,
		makeBitBlock(5, 0b00011100),
	)
}

func TestBitBlockAppendByte(t *testing.T) {
	check := func(block bitBlock, b byte, expected bitBlock) {
		block.appendByte(b)
		assert.Equal(t, expected, block)
	}

	check(
		makeBitBlock(4, 0b00001001),
		0b11001010,
		makeBitBlock(12, 0b10101001, 0b00001100),
	)

	check(
		makeBitBlock(7, 0b01101001),
		0b10010010,
		makeBitBlock(15, 0b01101001, 0b01001001),
	)
}
