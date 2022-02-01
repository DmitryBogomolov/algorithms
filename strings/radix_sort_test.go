package strings_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/strings"
	"github.com/stretchr/testify/assert"
)

func TestRadixLSDSort(t *testing.T) {
	alph := NewRangeAlphabet('a', 'z')

	{
		RadixLSDSort(nil, 0, alph)
	}
	{
		items := []string{"test"}
		RadixLSDSort(items, 10, alph)
		assert.Equal(t, []string{"test"}, items)
	}
	{
		items := []string{
			"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
		}
		RadixLSDSort(items, 3, alph)
		assert.Equal(
			t,
			[]string{
				"ace", "add", "bad", "bed", "bee", "cab", "dab", "dad", "ebb", "fad", "fed", "fee",
			},
			items,
		)
	}
	{
		items := []string{
			"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
		}
		RadixLSDSort(items, 2, alph)
		assert.Equal(
			t, []string{
				"dab", "cab", "fad", "bad", "dad", "ebb", "ace", "add", "fed", "bed", "fee", "bee",
			},
			items,
		)
	}
	{
		items := []string{
			"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
		}
		RadixLSDSort(items, 1, alph)
		assert.Equal(
			t,
			[]string{
				"dab", "cab", "ebb", "add", "fad", "bad", "dad", "fed", "bed", "fee", "bee", "ace",
			},
			items,
		)
	}
	{
		items := []string{
			"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
		}
		RadixLSDSort(items, 0, alph)
		assert.Equal(
			t,
			[]string{
				"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
			},
			items,
		)
	}
}

func TestRadixMSDSort(t *testing.T) {
	alph := NewRangeAlphabet('a', 'z')

	{
		RadixMSDSort(nil, 0, alph)
	}
	{
		items := []string{"test"}
		RadixMSDSort(items, 10, alph)
		assert.Equal(t, []string{"test"}, items)
	}
	{
		items := []string{
			"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
		}
		RadixMSDSort(items, 3, alph)
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
		RadixMSDSort(items, 2, alph)
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
		RadixMSDSort(items, 1, alph)
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
		RadixMSDSort(items, 0, alph)
		assert.Equal(
			t,
			[]string{
				"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
			},
			items,
		)
	}
}
