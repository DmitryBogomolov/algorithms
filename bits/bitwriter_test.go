package bits

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitWriter_WriteBit(t *testing.T) {
	var buffer bytes.Buffer
	writer := NewBitWriter(&buffer)

	writer.WriteBit(0)
	writer.WriteBit(1)
	writer.WriteBit(1)
	writer.WriteBit(0)
	assert.Equal(t, []byte(nil), buffer.Bytes())
	writer.Flush()
	assert.Equal(t, []byte{0b0110}, buffer.Bytes())
}

func TestBitWriter_Flush(t *testing.T) {
	var buffer bytes.Buffer
	writer := NewBitWriter(&buffer)

	writer.WriteBit(1)
	writer.WriteBit(0)
	writer.WriteBit(1)
	writer.Flush()
	writer.WriteBit(1)
	writer.WriteBit(1)
	writer.Flush()
	assert.Equal(t, []byte{0b101, 0b11}, buffer.Bytes())
}

func TestBitWriter_FlushAligned(t *testing.T) {
	var buffer bytes.Buffer
	writer := NewBitWriter(&buffer)

	writer.Flush()
	assert.Equal(t, []byte(nil), buffer.Bytes())
}

func TestBitWriter_WriteBits8(t *testing.T) {
	var buffer bytes.Buffer
	writer := NewBitWriter(&buffer)

	writer.WriteBit(1)
	writer.WriteBit(0)
	writer.WriteBit(1)
	writer.WriteBits8(0b00101110)
	writer.WriteBit(1)
	writer.Flush()
	assert.Equal(t, []byte{0b01110101, 0b1001}, buffer.Bytes())
}

func TestBitWriter_WriteBits8Aligned(t *testing.T) {
	var buffer bytes.Buffer
	writer := NewBitWriter(&buffer)

	writer.WriteBits8(125)
	writer.WriteBits8(201)
	writer.Flush()
	assert.Equal(t, []byte{125, 201}, buffer.Bytes())
}

func TestBitWriter_WriteBits16(t *testing.T) {
	var buffer bytes.Buffer
	writer := NewBitWriter(&buffer)

	writer.WriteBit(1)
	writer.WriteBit(0)
	writer.WriteBit(1)
	writer.WriteBits16(0b_00110011_10011001)
	writer.WriteBit(1)
	writer.Flush()
	assert.Equal(t, []byte{0b11001101, 0b10011100, 0b1001}, buffer.Bytes())
}

func TestBitWriter_WriteBits16Aligned(t *testing.T) {
	var buffer bytes.Buffer
	writer := NewBitWriter(&buffer)

	writer.WriteBits16(12001)
	writer.WriteBits16(40022)
	writer.Flush()
	assert.Equal(
		t,
		[]byte{12001 & 0xFF, (12001 >> 8) & 0xFF, 40022 & 0xFF, (40022 >> 8) & 0xFF},
		buffer.Bytes(),
	)
}

func TestBitWriter_WriteBits32(t *testing.T) {
	var buffer bytes.Buffer
	writer := NewBitWriter(&buffer)

	writer.WriteBit(1)
	writer.WriteBit(0)
	writer.WriteBit(1)
	writer.WriteBits32(0b_00110011_01100110_00001111_01010101)
	writer.WriteBit(1)
	writer.Flush()
	assert.Equal(t, []byte{0b10101101, 0b01111010, 0b00110000, 0b10011011, 0b1001}, buffer.Bytes())
}

func TestBitWriter_WriteBits32Aligned(t *testing.T) {
	var buffer bytes.Buffer
	writer := NewBitWriter(&buffer)

	writer.WriteBits32(100123)
	writer.WriteBits32(200801)
	writer.Flush()
	assert.Equal(
		t,
		[]byte{
			100123 & 0xFF,
			(100123 >> 8) & 0xFF,
			(100123 >> 16) & 0xFF,
			(100123 >> 24) & 0xFF,
			200801 & 0xFF,
			(200801 >> 8) & 0xFF,
			(200801 >> 16) & 0xFF,
			(200801 >> 24) & 0xFF,
		},
		buffer.Bytes(),
	)
}
