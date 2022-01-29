package bits

import "io"

// BitWriter writes bit by bit to an underlying writer.
type BitWriter struct {
	writer      io.ByteWriter
	currentByte byte
	bitOffset   int
	bitCount    int
}

// NewBitWriter creates BitWriter instance.
func NewBitWriter(writer io.ByteWriter) *BitWriter {
	return &BitWriter{
		writer: writer,
	}
}

// BitCount gets amount of written bits.
func (writer *BitWriter) BitCount() int {
	return writer.bitCount
}

func (writer *BitWriter) writeByte() error {
	err := writer.writer.WriteByte(writer.currentByte)
	writer.currentByte = 0
	return err
}

// WriteBit writes a bit.
func (writer *BitWriter) WriteBit(value byte) error {
	mask := bitMasks[writer.bitOffset]
	if value > 0 {
		writer.currentByte |= mask
	} else {
		writer.currentByte &= ^mask
	}
	writer.bitCount++
	writer.bitOffset++
	if writer.bitOffset == 8 {
		return writer.Flush()
	}
	return nil
}

// Flush alignes to the next byte.
func (writer *BitWriter) Flush() error {
	if writer.bitOffset > 0 {
		writer.bitCount += 8 - writer.bitOffset
		writer.bitOffset = 0
		return writer.writeByte()
	}
	return nil
}

// WriteUint8 writes an uint8.
func (writer *BitWriter) WriteUint8(value byte) error {
	writer.currentByte |= value << writer.bitOffset
	err := writer.writeByte()
	writer.currentByte |= value >> (8 - writer.bitOffset)
	return err
}

// WriteUint16 writes an uint16.
func (writer *BitWriter) WriteUint16(value uint16) error {
	err := writer.WriteUint8(byte(value))
	if err == nil {
		err = writer.WriteUint8(byte(value >> 8))
	}
	return err
}

// WriteUint32 writes an uint32.
func (writer *BitWriter) WriteUint32(value uint32) error {
	err := writer.WriteUint16(uint16(value))
	if err == nil {
		writer.WriteUint16(uint16(value >> 16))
	}
	return err
}
