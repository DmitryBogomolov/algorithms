package huffman_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/huffman"
	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	data, err := Compress([]byte("Hello World"))
	assert.Equal(t, nil, err)
	assert.Equal(
		t,
		[]byte{0x08, 0x52, 0xce, 0xb2, 0x57, 0xb2, 0xf9, 0x26, 0x52, 0x32, 0x0b, 0x00, 0x00, 0x00, 0xa7, 0x1a, 0x3c, 0xf6},
		data,
	)

	data, err = Compress([]byte("ABRACADABRA!"))
	assert.Equal(t, nil, err)
	assert.Equal(
		t,
		[]byte{0x06, 0x91, 0xc8, 0x90, 0x43, 0x4a, 0x15, 0x02, 0xc, 0x00, 0x00, 0x00, 0x3e, 0x2d, 0x3e, 0x05},
		data,
	)
}

func TestCompressEmptyBuffer(t *testing.T) {
	var ret []byte
	var err error

	ret, err = Compress(nil)
	assert.Equal(t, []byte(nil), ret)
	assert.Equal(t, ErrEmptyData, err)

	ret, err = Compress([]byte{})
	assert.Equal(t, []byte(nil), ret)
	assert.Equal(t, ErrEmptyData, err)

	ret, err = Compress([]byte{1})
	assert.Equal(t, []byte{3, 0, 1, 0, 0, 0}, ret)
	assert.Equal(t, nil, err)
}
