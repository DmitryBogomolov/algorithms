package huffman_test

import (
	"math/rand"
	"testing"

	. "github.com/DmitryBogomolov/algorithms/huffman"
	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	data, err := Compress([]byte("Hello World"))
	assert.Equal(t, nil, err)
	assert.Equal(
		t,
		[]byte{0x7c, 0x13, 0x64, 0x99, 0x6c, 0x24, 0x8b, 0xe4, 0x55, 0x39, 0xb, 0x00, 0x00, 0x00, 0xeb, 0x42, 0xc7, 0x37},
		data,
	)
}

func TestExpand(t *testing.T) {
	data, err := Expand(
		[]byte{0x7c, 0x13, 0x64, 0x99, 0x6c, 0x24, 0x8b, 0xe4, 0x55, 0x39, 0xb, 0x00, 0x00, 0x00, 0xeb, 0x42, 0xc7, 0x37},
	)
	assert.Equal(t, nil, err)
	assert.Equal(t, []byte("Hello World"), data)
}

func TestCompressExpand(t *testing.T) {
	sample := []byte("it was the best of times it was the worst of times")
	compressed, compressErr := Compress(sample)
	expanded, expandErr := Expand(compressed)
	assert.Equal(t, sample, expanded)
	assert.Equal(t, nil, compressErr)
	assert.Equal(t, nil, expandErr)
}

func TestCompressExpandLarge(t *testing.T) {
	sample := make([]byte, 9e6)
	r := rand.New(rand.NewSource(1244123))

	for i := 0; i < 5; i++ {
		r.Read(sample)

		compressed, compressErr := Compress(sample)
		expanded, expandErr := Expand(compressed)

		assert.Equal(t, sample, expanded)
		assert.Equal(t, nil, compressErr)
		assert.Equal(t, nil, expandErr)
	}
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
	assert.Equal(t, ErrDataCorrupted, err)
}
