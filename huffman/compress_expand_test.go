package huffman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressExpand(t *testing.T) {
	sample := []byte("it was the best of times it was the worst of times")
	tmp := Compress(sample)
	x := Expand(tmp)
	assert.Equal(t, sample, x)
}
