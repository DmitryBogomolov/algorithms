package huffman_test

import (
	"math/rand"
	"testing"

	. "github.com/DmitryBogomolov/algorithms/huffman"
	"github.com/stretchr/testify/assert"
)

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
