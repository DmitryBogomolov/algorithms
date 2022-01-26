package strings_test

import (
	"strings"
	"testing"

	. "github.com/DmitryBogomolov/algorithms/strings"
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

func makeTestTrie() *Trie {
	trie := NewTrie(testAlphabet{})
	for i, str := range strings.Split("she sells sea shells by the sea shore", " ") {
		trie.Put(str, i)
	}
	return trie
}

func TestAlphabet(t *testing.T) {
	var alphabet testAlphabet
	assert.Equal(t, 0, alphabet.ToIndex('a'))
	assert.Equal(t, 25, alphabet.ToIndex('z'))
	assert.Equal(t, 'b', alphabet.ToSymbol(1))
	assert.Equal(t, 'y', alphabet.ToSymbol(24))
}

func TestSize(t *testing.T) {
	trie := makeTestTrie()

	assert.Equal(t, 7, trie.Size())
}

func TestGet(t *testing.T) {
	trie := makeTestTrie()

	assert.Equal(t, 0, trie.Get("she"))
	assert.Equal(t, 1, trie.Get("sells"))
	assert.Equal(t, 3, trie.Get("shells"))
	assert.Equal(t, 4, trie.Get("by"))
	assert.Equal(t, 5, trie.Get("the"))
	assert.Equal(t, 6, trie.Get("sea"))
	assert.Equal(t, 7, trie.Get("shore"))
	assert.Equal(t, nil, trie.Get("sh"))
	assert.Equal(t, nil, trie.Get("sher"))
}

func TestDel(t *testing.T) {
	trie := makeTestTrie()

	trie.Del("shells")
	assert.Equal(t, nil, trie.Get("shells"))

	trie.Del("sea")
	trie.Del("sea")
	assert.Equal(t, nil, trie.Get("sea"))
}

func TestKeys(t *testing.T) {
	trie := makeTestTrie()

	assert.Equal(t, []string{"by", "sea", "sells", "she", "shells", "shore", "the"}, trie.KeysWithPrefix(""))
}

func TestKeysWithPrefix(t *testing.T) {
	trie := makeTestTrie()

	assert.Equal(t, []string{"by", "sea", "sells", "she", "shells", "shore", "the"}, trie.KeysWithPrefix(""))
	assert.Equal(t, []string{"she", "shells", "shore"}, trie.KeysWithPrefix("sh"))
	assert.Equal(t, []string{"she", "shells"}, trie.KeysWithPrefix("she"))
	assert.Equal(t, []string(nil), trie.KeysWithPrefix("tt"))
}

func TestKeysThatMatch(t *testing.T) {
	trie := makeTestTrie()

	assert.Equal(t, []string(nil), trie.KeysThatMatch(""))
	assert.Equal(t, []string{"by"}, trie.KeysThatMatch("by"))
	assert.Equal(t, []string{"by"}, trie.KeysThatMatch(".y"))
	assert.Equal(t, []string{"by"}, trie.KeysThatMatch("b."))
	assert.Equal(t, []string{"sea", "she", "the"}, trie.KeysThatMatch("..."))
}

func TestLongestPrefix(t *testing.T) {
	trie := makeTestTrie()

	assert.Equal(t, "she", trie.LongestPrefix("she"))
	assert.Equal(t, "she", trie.LongestPrefix("shell"))
	assert.Equal(t, "shells", trie.LongestPrefix("shellsort"))
	assert.Equal(t, "she", trie.LongestPrefix("shelters"))
}
