package priorityqueue_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/priorityqueue"
	"github.com/stretchr/testify/assert"
)

func TestMinPriorityQueue(t *testing.T) {
	queue := New(func(lhs, rhs interface{}) bool {
		return lhs.(int) < rhs.(int)
	})

	assert.Equal(t, queue.Size(), 0)

	queue.Insert(3)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, queue.Peek().(int), 3)

	queue.Insert(4)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek().(int), 3)

	queue.Insert(2)
	assert.Equal(t, queue.Size(), 3)
	assert.Equal(t, queue.Peek().(int), 2)

	queue.Insert(1)
	assert.Equal(t, queue.Size(), 4)
	assert.Equal(t, queue.Peek().(int), 1)

	assert.Equal(t, queue.Remove().(int), 1)
	assert.Equal(t, queue.Size(), 3)
	assert.Equal(t, queue.Peek().(int), 2)

	assert.Equal(t, queue.Remove().(int), 2)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek().(int), 3)

	assert.Equal(t, queue.Remove().(int), 3)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, queue.Peek().(int), 4)

	assert.Equal(t, queue.Remove().(int), 4)
	assert.Equal(t, queue.Size(), 0)
}

func TestMaxPriorityQueue(t *testing.T) {
	queue := New(func(lhs interface{}, rhs interface{}) bool {
		return lhs.(float64) > rhs.(float64)
	})

	assert.Equal(t, queue.Size(), 0)

	queue.Insert(1.2)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, queue.Peek().(float64), 1.2)

	queue.Insert(2.3)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek().(float64), 2.3)

	queue.Insert(1.9)
	assert.Equal(t, queue.Size(), 3)
	assert.Equal(t, queue.Peek().(float64), 2.3)

	queue.Insert(3.1)
	assert.Equal(t, queue.Size(), 4)
	assert.Equal(t, queue.Peek().(float64), 3.1)

	assert.Equal(t, queue.Remove().(float64), 3.1)
	assert.Equal(t, queue.Size(), 3)
	assert.Equal(t, queue.Peek().(float64), 2.3)

	assert.Equal(t, queue.Remove().(float64), 2.3)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek().(float64), 1.9)

	assert.Equal(t, queue.Remove().(float64), 1.9)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, queue.Peek().(float64), 1.2)

	assert.Equal(t, queue.Remove().(float64), 1.2)
	assert.Equal(t, queue.Size(), 0)
}
