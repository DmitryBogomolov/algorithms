package priorityqueue_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/priorityqueue"
	"github.com/stretchr/testify/assert"
)

func TestMinPriorityQueue(t *testing.T) {
	queue := New(func(lhs, rhs int) bool {
		return lhs < rhs
	})

	assert.Equal(t, queue.Size(), 0)

	queue.Insert(3)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, queue.Peek(), 3)

	queue.Insert(4)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek(), 3)

	queue.Insert(2)
	assert.Equal(t, queue.Size(), 3)
	assert.Equal(t, queue.Peek(), 2)

	queue.Insert(1)
	assert.Equal(t, queue.Size(), 4)
	assert.Equal(t, queue.Peek(), 1)

	assert.Equal(t, queue.Remove(), 1)
	assert.Equal(t, queue.Size(), 3)
	assert.Equal(t, queue.Peek(), 2)

	assert.Equal(t, queue.Remove(), 2)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek(), 3)

	assert.Equal(t, queue.Remove(), 3)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, queue.Peek(), 4)

	assert.Equal(t, queue.Remove(), 4)
	assert.Equal(t, queue.Size(), 0)
}

func TestMaxPriorityQueue(t *testing.T) {
	queue := New(func(lhs, rhs float64) bool {
		return lhs > rhs
	})

	assert.Equal(t, queue.Size(), 0)

	queue.Insert(1.2)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, queue.Peek(), 1.2)

	queue.Insert(2.3)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek(), 2.3)

	queue.Insert(1.9)
	assert.Equal(t, queue.Size(), 3)
	assert.Equal(t, queue.Peek(), 2.3)

	queue.Insert(3.1)
	assert.Equal(t, queue.Size(), 4)
	assert.Equal(t, queue.Peek(), 3.1)

	assert.Equal(t, queue.Remove(), 3.1)
	assert.Equal(t, queue.Size(), 3)
	assert.Equal(t, queue.Peek(), 2.3)

	assert.Equal(t, queue.Remove(), 2.3)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek(), 1.9)

	assert.Equal(t, queue.Remove(), 1.9)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, queue.Peek(), 1.2)

	assert.Equal(t, queue.Remove(), 1.2)
	assert.Equal(t, queue.Size(), 0)
}
