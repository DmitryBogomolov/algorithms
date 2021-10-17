package indexpriorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinIndexPriorityQueue(t *testing.T) {
	queue := New(func(lhs, rhs interface{}) bool {
		return lhs.(int) < rhs.(int)
	})

	assert.Equal(t, queue.Size(), 0)

	queue.Insert(11, 8)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, queue.Peek().(int), 8)
	assert.Equal(t, queue.PeekKey(), 11)

	queue.Insert(12, 5)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek().(int), 5)
	assert.Equal(t, queue.PeekKey(), 12)

	queue.Insert(13, 9)
	assert.Equal(t, queue.Size(), 3)
	assert.Equal(t, queue.Peek().(int), 5)
	assert.Equal(t, queue.PeekKey(), 12)

	queue.Insert(14, 4)
	assert.Equal(t, queue.Size(), 4)
	assert.Equal(t, queue.Peek().(int), 4)
	assert.Equal(t, queue.PeekKey(), 14)

	var element interface{}
	var key int
	element, key = queue.Remove()
	assert.Equal(t, element.(int), 4)
	assert.Equal(t, key, 14)
	assert.Equal(t, queue.Size(), 3)
	assert.Equal(t, queue.Peek().(int), 5)
	assert.Equal(t, queue.PeekKey(), 12)

	element, key = queue.Remove()
	assert.Equal(t, element.(int), 5)
	assert.Equal(t, key, 12)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek().(int), 8)
	assert.Equal(t, queue.PeekKey(), 11)

	element, key = queue.Remove()
	assert.Equal(t, element.(int), 8)
	assert.Equal(t, key, 11)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, queue.Peek().(int), 9)
	assert.Equal(t, queue.PeekKey(), 13)

	element, key = queue.Remove()
	assert.Equal(t, element.(int), 9)
	assert.Equal(t, key, 13)
	assert.Equal(t, queue.Size(), 0)
}

func TestMaxIndexPriorityQueue(t *testing.T) {
	queue := New(func(lhs, rhs interface{}) bool {
		return lhs.(float64) > rhs.(float64)
	})

	assert.Equal(t, queue.Size(), 0)

	queue.Insert(101, 3.2)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, queue.Peek().(float64), 3.2)
	assert.Equal(t, queue.PeekKey(), 101)

	queue.Insert(102, 4.1)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek().(float64), 4.1)
	assert.Equal(t, queue.PeekKey(), 102)

	queue.Insert(103, 3.5)
	assert.Equal(t, queue.Size(), 3)
	assert.Equal(t, queue.Peek().(float64), 4.1)
	assert.Equal(t, queue.PeekKey(), 102)

	queue.Insert(104, 4.4)
	assert.Equal(t, queue.Size(), 4)
	assert.Equal(t, queue.Peek().(float64), 4.4)
	assert.Equal(t, queue.PeekKey(), 104)

	var element interface{}
	var key int
	element, key = queue.Remove()
	assert.Equal(t, element.(float64), 4.4)
	assert.Equal(t, key, 104)
	assert.Equal(t, queue.Size(), 3)
	assert.Equal(t, queue.Peek().(float64), 4.1)
	assert.Equal(t, queue.PeekKey(), 102)

	element, key = queue.Remove()
	assert.Equal(t, element.(float64), 4.1)
	assert.Equal(t, key, 102)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek().(float64), 3.5)
	assert.Equal(t, queue.PeekKey(), 103)

	element, key = queue.Remove()
	assert.Equal(t, element.(float64), 3.5)
	assert.Equal(t, key, 103)
	assert.Equal(t, queue.Size(), 1)
	assert.Equal(t, queue.Peek().(float64), 3.2)
	assert.Equal(t, queue.PeekKey(), 101)

	element, key = queue.Remove()
	assert.Equal(t, element.(float64), 3.2)
	assert.Equal(t, key, 101)
	assert.Equal(t, queue.Size(), 0)
}

func TestChangeByKey(t *testing.T) {
	queue := New(func(lhs, rhs interface{}) bool {
		return lhs.(float64) < rhs.(float64)
	})

	queue.Insert(1, 1.3)
	queue.Insert(2, 2.4)
	assert.Equal(t, queue.HasKey(1), true)
	assert.Equal(t, queue.HasKey(2), true)
	assert.Equal(t, queue.HasKey(3), false)

	queue.Insert(1, 3.1)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek().(float64), 2.4)
	assert.Equal(t, queue.PeekKey(), 2)

	queue.Insert(2, 2.2)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek().(float64), 2.2)
	assert.Equal(t, queue.PeekKey(), 2)

	queue.Insert(2, 4.2)
	assert.Equal(t, queue.Size(), 2)
	assert.Equal(t, queue.Peek().(float64), 3.1)
	assert.Equal(t, queue.PeekKey(), 1)

	queue.Insert(3, 2.1)
	assert.Equal(t, queue.Size(), 3)
	assert.Equal(t, queue.Peek().(float64), 2.1)
	assert.Equal(t, queue.PeekKey(), 3)
}

func TestRemoveByKey(t *testing.T) {
	queue := New(func(lhs, rhs interface{}) bool {
		return lhs.(float64) < rhs.(float64)
	})

	queue.Insert(1, 1.3)
	queue.Insert(2, 2.4)
	queue.Insert(3, 3.5)

	assert.Equal(t, queue.HasKey(1), true)
	assert.Equal(t, queue.HasKey(2), true)
	assert.Equal(t, queue.HasKey(3), true)
	assert.Equal(t, queue.HasKey(4), false)

	assert.Equal(t, queue.RemoveByKey(2).(float64), 2.4)
	assert.Equal(t, queue.HasKey(2), false)
	assert.Equal(t, queue.Size(), 2)

	assert.Equal(t, queue.RemoveByKey(3).(float64), 3.5)
	assert.Equal(t, queue.HasKey(3), false)
	assert.Equal(t, queue.Size(), 1)

	assert.Equal(t, queue.RemoveByKey(1).(float64), 1.3)
	assert.Equal(t, queue.HasKey(1), false)
	assert.Equal(t, queue.Size(), 0)
}
