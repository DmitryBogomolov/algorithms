package bits

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
