package huffman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	str := "it was the best of times it was the worst of times"
	tmp := Compress([]byte(str))
	x := Decompress(tmp)
	s := string(x)
	assert.Equal(t, str, s)
}
