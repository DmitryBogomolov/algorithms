package huffman_test

import (
	"bytes"
	"math/rand"
	"testing"

	. "github.com/DmitryBogomolov/algorithms/huffman"
	"github.com/stretchr/testify/assert"
)

func TestCompressExpand(t *testing.T) {
	sample := []byte("it was the best of times it was the worst of times")

	var compressBuffer bytes.Buffer
	compressErr := Compress(bytes.NewBuffer(sample), &compressBuffer)
	var expandBuffer bytes.Buffer
	expandErr := Expand(&compressBuffer, &expandBuffer)

	assert.Equal(t, nil, compressErr)
	assert.Equal(t, nil, expandErr)
	assert.Equal(t, sample, expandBuffer.Bytes())
}

func TestCompressExpandLarge(t *testing.T) {
	sample := make([]byte, 1e6)
	r := rand.New(rand.NewSource(1244123))

	for i := 0; i < 5; i++ {
		r.Read(sample)

		var compressBuffer bytes.Buffer
		compressErr := Compress(bytes.NewBuffer(sample), &compressBuffer)
		var expandBuffer bytes.Buffer
		expandErr := Expand(&compressBuffer, &expandBuffer)

		assert.Equal(t, nil, compressErr)
		assert.Equal(t, nil, expandErr)
		assert.Equal(t, sample, expandBuffer.Bytes())
	}
}
