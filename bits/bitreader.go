package bits

import "io"

type _BitReader struct {
	reader      io.ByteReader
	currentByte byte
	bitOffset   int
	bitCount    int
}

func NewBitReader(reader io.ByteReader) *_BitReader {
	return &_BitReader{
		reader: reader,
	}
}

func (reader *_BitReader) readByte() error {
	nextByte, err := reader.reader.ReadByte()
	if err != nil {
		return err
	}
	reader.currentByte = nextByte
	return nil
}

func (reader *_BitReader) BitCount() int {
	return reader.bitCount
}

func (reader *_BitReader) ReadBit() (byte, error) {
	if reader.bitOffset == 0 {
		if err := reader.readByte(); err != nil {
			return 0, err
		}
	}
	mask := bitMasks[reader.bitOffset]
	bit := reader.currentByte & mask
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

func (reader *_BitReader) Flush() (byte, error) {
	if reader.bitOffset == 0 {
		return 0, nil
	}
	ret := reader.currentByte >> reader.bitOffset
	reader.currentByte = 0
	reader.bitCount += 8 - reader.bitOffset
	reader.bitOffset = 0
	return ret, nil
}

func (reader *_BitReader) ReadUint8() (uint8, error) {
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

func (reader *_BitReader) ReadUint16() (uint16, error) {
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

func (reader *_BitReader) ReadUint32() (uint32, error) {
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
