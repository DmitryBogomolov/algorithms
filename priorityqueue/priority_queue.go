package priorityqueue

import (
	"container/heap"
)

func (queue _PriorityQueue) Len() int {
	return len(queue.items)
}

func (queue _PriorityQueue) Swap(i, j int) {
	items := queue.items
	items[i], items[j] = items[j], items[i]
}

func (queue _PriorityQueue) Less(i, j int) bool {
	items := queue.items
	return queue.less(items[i], items[j])
}

func (queue *_PriorityQueue) Push(item interface{}) {
	queue.items = append(queue.items, item)
}

func (queue *_PriorityQueue) Pop() interface{} {
	last := queue.Len() - 1
	value := queue.items[last]
	queue.items[last] = nil
	queue.items = queue.items[0:last]
	return value
}

// LessFunc is an ordering function.
type LessFunc func(lhs, rhs interface{}) bool

type _PriorityQueue struct {
	items []interface{}
	less  LessFunc
}

func (queue _PriorityQueue) Size() int {
	return queue.Len()
}

func (queue *_PriorityQueue) Insert(element interface{}) {
	heap.Push(queue, element)
}

func (queue *_PriorityQueue) Remove() interface{} {
	return heap.Pop(queue)
}

func (queue _PriorityQueue) Peek() interface{} {
	return queue.items[0]
}

// PriorityQueue is priority queue data structure.
type PriorityQueue interface {
	// Size gets number of element in a queue.
	Size() int
	// Insert adds element to a queue.
	Insert(element interface{})
	// Remove removes element from a queue.
	Remove() interface{}
	// Peek returns first element of a queue.
	Peek() interface{}
}

// New create instance of PriorityQueue.
func New(less LessFunc) PriorityQueue {
	queue := _PriorityQueue{
		less:  less,
		items: nil,
	}
	return &queue
}
