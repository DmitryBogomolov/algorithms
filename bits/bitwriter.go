package bits

import "io"

type _BitWriter struct {
	writer      io.ByteWriter
	currentByte byte
	bitOffset   int
	bitCount    int
}

func NewBitWriter(writer io.ByteWriter) *_BitWriter {
	return &_BitWriter{
		writer: writer,
	}
}

func (writer *_BitWriter) BitCount() int {
	return writer.bitCount
}

func (writer *_BitWriter) writeByte() error {
	err := writer.writer.WriteByte(writer.currentByte)
	writer.currentByte = 0
	return err
}

func (writer *_BitWriter) WriteBit(value byte) error {
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

func (writer *_BitWriter) Flush() error {
	if writer.bitOffset > 0 {
		writer.bitCount += 8 - writer.bitOffset
		writer.bitOffset = 0
		return writer.writeByte()
	}
	return nil
}

func (writer *_BitWriter) WriteUint8(value byte) error {
	writer.currentByte |= value << writer.bitOffset
	err := writer.writeByte()
	writer.currentByte |= value >> (8 - writer.bitOffset)
	return err
}

func (writer *_BitWriter) WriteUint16(value uint16) error {
	err := writer.WriteUint8(byte(value))
	if err == nil {
		err = writer.WriteUint8(byte(value >> 8))
	}
	return err
}

func (writer *_BitWriter) WriteUint32(value uint32) error {
	err := writer.WriteUint16(uint16(value))
	if err == nil {
		writer.WriteUint16(uint16(value >> 16))
	}
	return err
}
