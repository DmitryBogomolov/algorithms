package bits_test

import (
	"bytes"
	"io"
	"testing"

	. "github.com/DmitryBogomolov/algorithms/bits"
	"github.com/stretchr/testify/assert"
)

func TestBitReader_ReadBit(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{0b11001})
	reader := NewBitReader(buffer)

	checkBit := func(expected byte) {
		bit, _ := reader.ReadBit()
		assert.Equal(t, expected, bit)
	}

	checkBit(1)
	checkBit(0)
	checkBit(0)
	checkBit(1)
	checkBit(1)

	checkBit(0)
	checkBit(0)
	checkBit(0)
	_, err := reader.ReadBit()
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, 8, reader.BitCount())
}

func TestBitReader_Flush(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{0b10101})
	reader := NewBitReader(buffer)

	reader.ReadBit()
	reader.ReadBit()
	reader.ReadBit()
	val, err := reader.Flush()
	assert.Equal(t, byte(0b10), val)
	assert.Equal(t, nil, err)
	assert.Equal(t, 8, reader.BitCount())
}

func TestBitReader_FlushAligned(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	reader := NewBitReader(buffer)

	val, err := reader.Flush()
	assert.Equal(t, byte(0), val)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, reader.BitCount())
}

func TestBitReader_ReadUint8(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{0b10011001, 0b101})
	reader := NewBitReader(buffer)
	var val uint8
	var err error

	reader.ReadBit()
	reader.ReadBit()
	val, err = reader.ReadUint8()
	assert.Equal(t, byte(0b01100110), val)
	assert.Equal(t, nil, err)
	val, err = reader.ReadUint8()
	assert.Equal(t, byte(0b1), val)
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, 16, reader.BitCount())
}

func TestBitReader_ReadUint8Aligned(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{137, 243})
	reader := NewBitReader(buffer)
	var val uint8
	var err error

	val, err = reader.ReadUint8()
	assert.Equal(t, byte(137), val)
	assert.Equal(t, nil, err)
	val, err = reader.ReadUint8()
	assert.Equal(t, nil, err)
	assert.Equal(t, byte(243), val)
	val, err = reader.ReadUint8()
	assert.Equal(t, byte(0), val)
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, 16, reader.BitCount())
}

func TestBitReader_ReadUint16(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{0b10010011, 0b00111001, 0b10100101})
	reader := NewBitReader(buffer)
	var val uint16
	var err error

	reader.ReadBit()
	reader.ReadBit()
	val, err = reader.ReadUint16()
	assert.Equal(t, uint16(0b_01001110_01100100), val)
	assert.Equal(t, nil, err)
	val, err = reader.ReadUint16()
	assert.Equal(t, uint16(0b101001), val)
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, 24, reader.BitCount())
}

func TestBitReader_ReadUint16Aligned(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{12991 & 0xFF, (12991 >> 8) & 0xFF, 41011 & 0xFF, (41011 >> 8) & 0xFF})
	reader := NewBitReader(buffer)
	var val uint16
	var err error

	val, err = reader.ReadUint16()
	assert.Equal(t, uint16(12991), val)
	assert.Equal(t, nil, err)
	val, err = reader.ReadUint16()
	assert.Equal(t, uint16(41011), val)
	assert.Equal(t, nil, err)
	val, err = reader.ReadUint16()
	assert.Equal(t, uint16(0), val)
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, 32, reader.BitCount())
}

func TestBitReader_ReadUint32(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{0b01100110, 0b00110011, 0b10011001, 0b10010110, 0b11001100})
	reader := NewBitReader(buffer)
	var val uint32
	var err error

	reader.ReadBit()
	reader.ReadBit()
	val, err = reader.ReadUint32()
	assert.Equal(t, uint32(0b_00100101_10100110_01001100_11011001), val)
	assert.Equal(t, nil, err)
	val, err = reader.ReadUint32()
	assert.Equal(t, uint32(0b110011), val)
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, 40, reader.BitCount())
}

func TestBitReader_ReadUint32Aligned(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{
		123100 & 0xFF, (123100 >> 8) & 0xFF, (123100 >> 16) & 0xFF, (123100 >> 24) & 0xFF,
		431121 & 0xFF, (431121 >> 8) & 0xFF, (431121 >> 16) & 0xFF, (431121 >> 24) & 0xFF,
	})
	reader := NewBitReader(buffer)
	var val uint32
	var err error

	val, err = reader.ReadUint32()
	assert.Equal(t, uint32(123100), val)
	assert.Equal(t, nil, err)
	val, err = reader.ReadUint32()
	assert.Equal(t, uint32(431121), val)
	assert.Equal(t, nil, err)
	val, err = reader.ReadUint32()
	assert.Equal(t, uint32(0), val)
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, 64, reader.BitCount())
}
