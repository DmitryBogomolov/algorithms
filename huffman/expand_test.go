package huffman_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	. "github.com/DmitryBogomolov/algorithms/huffman"
	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	{
		var buffer bytes.Buffer
		err := Expand(
			bytes.NewBuffer([]byte{
				0x08, 0x52, 0xce, 0xb2, 0x57, 0xb2, 0xf9, 0x26, 0x52, 0x32, 0x0b, 0x00, 0x00, 0x00, 0xa7, 0x1a, 0x3c, 0xf6,
			}),
			&buffer,
		)
		assert.Equal(t, nil, err)
		assert.Equal(t, []byte("Hello World"), buffer.Bytes())
	}
	{
		var buffer bytes.Buffer
		err := Expand(
			bytes.NewBuffer([]byte{
				0x06, 0x91, 0xc8, 0x90, 0x43, 0x4a, 0x15, 0x02, 0xc, 0x00, 0x00, 0x00, 0x3e, 0x2d, 0x3e, 0x05,
			}),
			&buffer,
		)
		assert.Equal(t, nil, err)
		assert.Equal(t, []byte("ABRACADABRA!"), buffer.Bytes())
	}
}

func TestExpandEmptyBuffer(t *testing.T) {
	{
		var buffer bytes.Buffer
		err := Expand(bytes.NewBuffer(nil), &buffer)
		assert.Equal(t, io.EOF, errors.Unwrap(err))
		assert.Equal(t, []byte(nil), buffer.Bytes())
	}
	{
		var buffer bytes.Buffer
		err := Expand(bytes.NewBuffer([]byte{}), &buffer)
		assert.Equal(t, io.EOF, errors.Unwrap(err))
		assert.Equal(t, []byte(nil), buffer.Bytes())
	}
	{
		var buffer bytes.Buffer
		err := Expand(bytes.NewBuffer([]byte{1}), &buffer)
		assert.Equal(t, io.EOF, errors.Unwrap(err))
		assert.Equal(t, []byte(nil), buffer.Bytes())
	}
}
