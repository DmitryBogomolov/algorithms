package strings_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/strings"
	"github.com/stretchr/testify/assert"
)

func TestRadixMSD(t *testing.T) {
	alph := NewRangeAlphabet('a', 'z')

	{
		RadixMSD(nil, 0, alph)
	}
	{
		items := []string{"test"}
		RadixMSD(items, 10, alph)
		assert.Equal(t, []string{"test"}, items)
	}
	{
		items := []string{
			"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
		}
		RadixMSD(items, 3, alph)
		assert.Equal(
			t,
			[]string{
				"cab", "dab", "ebb", "bad", "dad", "fad", "add", "bed", "fed", "ace", "bee", "fee",
			},
			items,
		)
	}
	{
		items := []string{
			"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
		}
		RadixMSD(items, 2, alph)
		assert.Equal(
			t, []string{
				"bad", "cab", "dab", "dad", "fad", "ebb", "ace", "add", "bee", "bed", "fee", "fed",
			},
			items,
		)
	}
	{
		items := []string{
			"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
		}
		RadixMSD(items, 1, alph)
		assert.Equal(
			t,
			[]string{
				"add", "ace", "bad", "bee", "bed", "cab", "dab", "dad", "ebb", "fad", "fee", "fed",
			},
			items,
		)
	}
	{
		items := []string{
			"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
		}
		RadixMSD(items, 0, alph)
		assert.Equal(
			t,
			[]string{
				"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
			},
			items,
		)
	}
}
