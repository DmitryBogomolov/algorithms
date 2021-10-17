package priorityqueue

import (
	"container/heap"
)

func (queue _PriorityQueue) Len() int {
	return len(queue.items)
}

func (queue _PriorityQueue) Swap(i, j int) {
	queue.items[i], queue.items[j] = queue.items[j], queue.items[i]
}

func (queue _PriorityQueue) Less(i, j int) bool {
	return queue.less(queue.items[i], queue.items[j])
}

func (queue *_PriorityQueue) Push(item interface{}) {
	queue.items = append(queue.items, item)
}

func (queue *_PriorityQueue) Pop() interface{} {
	last := queue.Len() - 1
	item := queue.items[last]
	queue.items[last] = nil
	queue.items = queue.items[0:last]
	return item
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
	if queue.Len() == 0 {
		panic("queue is empty")
	}
	return heap.Pop(queue)
}

func (queue _PriorityQueue) Peek() interface{} {
	if queue.Len() == 0 {
		panic("queue is empty")
	}
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
	if less == nil {
		panic("less func is nil")
	}
	return &_PriorityQueue{
		less:  less,
		items: nil,
	}
}
