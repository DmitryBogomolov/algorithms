package bits

import "io"

// BitReader reads bit by bit from an underlying reader.
type BitReader struct {
	reader      io.ByteReader
	currentByte byte
	bitOffset   int
	bitCount    int
}

// NewBitReader creates BitReader instance.
func NewBitReader(reader io.ByteReader) *BitReader {
	return &BitReader{
		reader: reader,
	}
}

// BitCount gets amount of read bits.
func (reader *BitReader) BitCount() int {
	return reader.bitCount
}

func (reader *BitReader) readByte() error {
	nextByte, err := reader.reader.ReadByte()
	reader.currentByte = nextByte
	return err
}

// ReadBit reads a bit.
func (reader *BitReader) ReadBit() (byte, error) {
	if reader.bitOffset == 0 {
		if err := reader.readByte(); err != nil {
			return 0, err
		}
	}
	bit := reader.currentByte & bitMasks[reader.bitOffset]
	reader.bitCount++
	reader.bitOffset++
	if reader.bitOffset == 8 {
		reader.bitOffset = 0
	}
	if bit > 0 {
		return 1, nil
	}
	return 0, nil
}

// Flush alignes to the next byte.
func (reader *BitReader) Flush() (byte, error) {
	if reader.bitOffset == 0 {
		return 0, nil
	}
	ret := reader.currentByte >> reader.bitOffset
	reader.currentByte = 0
	reader.bitCount += 8 - reader.bitOffset
	reader.bitOffset = 0
	return ret, nil
}

// ReadUint8 reads an uint8.
func (reader *BitReader) ReadUint8() (uint8, error) {
	if reader.bitOffset == 0 {
		if err := reader.readByte(); err != nil {
			return 0, err
		}
	}
	ret := reader.currentByte >> reader.bitOffset
	reader.bitCount += 8 - reader.bitOffset
	if reader.bitOffset > 0 {
		if err := reader.readByte(); err != nil {
			return ret, err
		}
		ret |= reader.currentByte << (8 - reader.bitOffset)
		reader.bitCount += reader.bitOffset
	}
	return ret, nil
}

// ReadUint16 reads an uint16.
func (reader *BitReader) ReadUint16() (uint16, error) {
	var ret uint16
	var tmp uint8
	var err error
	tmp, err = reader.ReadUint8()
	ret = uint16(tmp)
	if err != nil {
		return ret, err
	}
	tmp, err = reader.ReadUint8()
	ret |= uint16(tmp) << 8
	return ret, err
}

// ReadUint32 reads an uint32.
func (reader *BitReader) ReadUint32() (uint32, error) {
	var ret uint32
	var tmp uint16
	var err error
	tmp, err = reader.ReadUint16()
	ret = uint32(tmp)
	if err != nil {
		return ret, err
	}
	tmp, err = reader.ReadUint16()
	ret |= uint32(tmp) << 16
	return ret, err
}
