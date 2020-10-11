package huffman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitScannerReadBit(t *testing.T) {
	scanner := newBitScanner([]byte{0b00110101, 0b00000101, 0b10100110})
	var bits []bool
	for i := 0; i < len(scanner.buffer)*8; i++ {
		bits = append(bits, scanner.readBit())
	}
	assert.Equal(t,
		[]bool{
			true, false, true, false, true, true, false, false,
			true, false, true, false, false, false, false, false,
			false, true, true, false, false, true, false, true,
		},
		bits,
	)
}

func TestBitScannerAlign(t *testing.T) {
	scanner := newBitScanner([]byte{0b00110101, 0b00001011, 0b10100110})

	assert.Equal(t,
		[]bool{true, false, true},
		[]bool{scanner.readBit(), scanner.readBit(), scanner.readBit()},
	)

	scanner.align()
	assert.Equal(t,
		[]bool{true, true, false, true, false},
		[]bool{scanner.readBit(), scanner.readBit(), scanner.readBit(), scanner.readBit(), scanner.readBit()},
	)

	scanner.align()
	assert.Equal(t,
		[]bool{false, true, true},
		[]bool{scanner.readBit(), scanner.readBit(), scanner.readBit()},
	)
}

func TestBitScannerReadByte(t *testing.T) {
	scanner := newBitScanner([]byte{0b00110101, 0b00001011, 0b10100110})

	scanner.readBit()
	assert.Equal(t, byte(0b10011010), scanner.readByte())

	scanner.align()
	assert.Equal(t, byte(0b10100110), scanner.readByte())
}
