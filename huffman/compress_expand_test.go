package huffman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressExpand(t *testing.T) {
	sample := []byte("it was the best of times it was the worst of times")
	tmp, _ := Compress(sample)
	x, _ := Expand(tmp)
	assert.Equal(t, sample, x)
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
