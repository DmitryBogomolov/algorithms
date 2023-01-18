package priorityqueue

import (
	"container/heap"
)

func (queue _PriorityQueue[T]) Len() int {
	return len(queue.items)
}

func (queue _PriorityQueue[T]) Swap(i, j int) {
	queue.items[i], queue.items[j] = queue.items[j], queue.items[i]
}

func (queue _PriorityQueue[T]) Less(i, j int) bool {
	return queue.less(queue.items[i], queue.items[j])
}

func (queue *_PriorityQueue[T]) Push(item interface{}) {
	queue.items = append(queue.items, item.(T))
}

func (queue *_PriorityQueue[T]) Pop() interface{} {
	last := queue.Len() - 1
	item := queue.items[last]
	var stub T
	queue.items[last] = stub
	queue.items = queue.items[0:last]
	return item
}

// LessFunc is an ordering function.
type LessFunc[T any] func(lhs, rhs T) bool

type _PriorityQueue[T any] struct {
	items []T
	less  LessFunc[T]
}

func (queue _PriorityQueue[T]) Size() int {
	return queue.Len()
}

func (queue *_PriorityQueue[T]) Insert(element T) {
	heap.Push(queue, element)
}

func (queue *_PriorityQueue[T]) Remove() T {
	if queue.Len() == 0 {
		panic("queue is empty")
	}
	return heap.Pop(queue).(T)
}

func (queue _PriorityQueue[T]) Peek() T {
	if queue.Len() == 0 {
		panic("queue is empty")
	}
	return queue.items[0]
}

// PriorityQueue is priority queue data structure.
type PriorityQueue[T any] interface {
	// Size gets number of element in a queue.
	Size() int
	// Insert adds element to a queue.
	Insert(element T)
	// Remove removes element from a queue.
	Remove() T
	// Peek returns first element of a queue.
	Peek() T
}

// New create instance of PriorityQueue.
func New[T any](less LessFunc[T]) PriorityQueue[T] {
	if less == nil {
		panic("less func is nil")
	}
	return &_PriorityQueue[T]{
		less:  less,
		items: nil,
	}
}
