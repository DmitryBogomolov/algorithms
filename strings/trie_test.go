package strings

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testAlphabet struct {
}

func (a testAlphabet) Size() int {
	return 'z' - 'a' + 1
}

func (a testAlphabet) ToIndex(symbol rune) int {
	return int(symbol - 'a')
}

func (a testAlphabet) ToSymbol(idx int) rune {
	return rune('a' + idx)
}

func TestTrie(t *testing.T) {
	var alphabet testAlphabet

	assert.Equal(t, 0, alphabet.ToIndex('a'))
	assert.Equal(t, 25, alphabet.ToIndex('z'))
	assert.Equal(t, 'b', alphabet.ToSymbol(1))
	assert.Equal(t, 'y', alphabet.ToSymbol(24))

	trie := NewTrie(alphabet)

	for i, str := range strings.Split("she sells sea shells by the sea shore", " ") {
		trie.Put(str, i)
	}

	assert.Equal(t, 7, trie.Size())

	assert.Equal(t, 0, trie.Get("she"))
	assert.Equal(t, 1, trie.Get("sells"))
	assert.Equal(t, 3, trie.Get("shells"))
	assert.Equal(t, 4, trie.Get("by"))
	assert.Equal(t, 5, trie.Get("the"))
	assert.Equal(t, 6, trie.Get("sea"))
	assert.Equal(t, 7, trie.Get("shore"))

	assert.Equal(t, NoValue, trie.Get(""))
	assert.Equal(t, NoValue, trie.Get("sh"))
	assert.Equal(t, NoValue, trie.Get("shoree"))

	trie.Del("shells")
	assert.Equal(t, NoValue, trie.Get("shells"))

}
