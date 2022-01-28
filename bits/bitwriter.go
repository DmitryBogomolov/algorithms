package bits

import (
	"io"
)

const (
	bit0 uint8 = 1 << iota
	bit1
	bit2
	bit3
	bit4
	bit5
	bit6
	bit7
)

var bitMasks = []byte{bit0, bit1, bit2, bit3, bit4, bit5, bit6, bit7}

type _BitWriter struct {
	writer      io.ByteWriter
	currentByte byte
	bitOffset   int
}

func NewBitWriter(writer io.ByteWriter) *_BitWriter {
	return &_BitWriter{
		writer: writer,
	}
}

func (writer *_BitWriter) WriteBit(value byte) {
	mask := bitMasks[writer.bitOffset]
	if value > 0 {
		writer.currentByte |= mask
	} else {
		writer.currentByte &= ^mask
	}
	writer.bitOffset++
	if writer.bitOffset == 8 {
		writer.Flush()
	}
}

func (writer *_BitWriter) WriteBits8(value byte) {
	writer.currentByte |= value << writer.bitOffset
	writer.writeByte()
	writer.currentByte |= value >> (8 - writer.bitOffset)
}

func (writer *_BitWriter) WriteBits16(value uint16) {
	writer.WriteBits8(byte(value))
	writer.WriteBits8(byte(value >> 8))
}

func (writer *_BitWriter) WriteBits32(value uint32) {
	writer.WriteBits16(uint16(value))
	writer.WriteBits16(uint16(value >> 16))
}

func (writer *_BitWriter) writeByte() {
	err := writer.writer.WriteByte(writer.currentByte)
	if err != nil {
		panic(err)
	}
	writer.currentByte = 0
}

func (writer *_BitWriter) Flush() {
	if writer.bitOffset > 0 {
		writer.writeByte()
		writer.bitOffset = 0
	}
}
