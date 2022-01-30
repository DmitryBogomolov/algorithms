package huffman_test

import (
	"errors"
	"io"
	"testing"

	. "github.com/DmitryBogomolov/algorithms/huffman"
	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	data, err := Expand(
		[]byte{0x08, 0x52, 0xce, 0xb2, 0x57, 0xb2, 0xf9, 0x26, 0x52, 0x32, 0x0b, 0x00, 0x00, 0x00, 0xa7, 0x1a, 0x3c, 0xf6},
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte("Hello World"), data)

	data, err = Expand(
		[]byte{0x06, 0x91, 0xc8, 0x90, 0x43, 0x4a, 0x15, 0x02, 0xc, 0x00, 0x00, 0x00, 0x3e, 0x2d, 0x3e, 0x05},
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte("ABRACADABRA!"), data)
}

func TestExpandEmptyBuffer(t *testing.T) {
	var ret []byte
	var err error

	ret, err = Expand(nil)
	assert.Equal(t, []byte(nil), ret)
	assert.Equal(t, ErrEmptyData, err)

	ret, err = Expand([]byte{})
	assert.Equal(t, []byte(nil), ret)
	assert.Equal(t, ErrEmptyData, err)

	ret, err = Expand([]byte{1})
	assert.Equal(t, []byte(nil), ret)
	assert.Equal(t, io.EOF, errors.Unwrap(err))
}
