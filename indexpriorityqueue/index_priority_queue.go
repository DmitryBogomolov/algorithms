package indexpriorityqueue

import (
	"container/heap"
)

// LessFunc is an ordering function.
type LessFunc[T any] func(lhs, rhs T) bool

func (queue _IndexPriorityQueue[T]) Len() int {
	return len(queue.items)
}

func (queue _IndexPriorityQueue[T]) Swap(i, j int) {
	queue.items[i], queue.items[j] = queue.items[j], queue.items[i]
	iKey, jKey := queue.idxToKey[i], queue.idxToKey[j]
	queue.keyToIdx[iKey], queue.keyToIdx[jKey] = j, i
	queue.idxToKey[i], queue.idxToKey[j] = jKey, iKey
}

func (queue _IndexPriorityQueue[T]) Less(i, j int) bool {
	return queue.less(queue.items[i], queue.items[j])
}

func (queue *_IndexPriorityQueue[T]) Push(item interface{}) {
	queue.items = append(queue.items, item.(T))
}

func (queue *_IndexPriorityQueue[T]) Pop() interface{} {
	last := queue.Len() - 1
	item := queue.items[last]
	var stub T
	queue.items[last] = stub
	queue.items = queue.items[0:last]
	return item
}

type _IndexPriorityQueue[T any] struct {
	items    []T
	keyToIdx map[int]int
	idxToKey map[int]int
	less     LessFunc[T]
}

func (queue _IndexPriorityQueue[T]) Size() int {
	return queue.Len()
}

func (queue *_IndexPriorityQueue[T]) Insert(key int, element T) {
	idx, hasKey := queue.keyToIdx[key]
	if hasKey {
		queue.items[idx] = element
		heap.Fix(queue, idx)
	} else {
		idx = queue.Len()
		queue.idxToKey[idx] = key
		queue.keyToIdx[key] = idx
		heap.Push(queue, element)
	}
}

func (queue *_IndexPriorityQueue[T]) Remove() (T, int) {
	if queue.Len() == 0 {
		panic("queue is empty")
	}
	key := queue.idxToKey[0]
	element := heap.Pop(queue).(T)
	delete(queue.idxToKey, queue.keyToIdx[key])
	delete(queue.keyToIdx, key)
	return element, key
}

func (queue _IndexPriorityQueue[T]) Peek() T {
	if queue.Len() == 0 {
		panic("queue is empty")
	}
	return queue.items[0]
}

func (queue *_IndexPriorityQueue[T]) RemoveByKey(key int) T {
	idx := queue.keyToIdx[key]
	element := heap.Remove(queue, idx).(T)
	delete(queue.idxToKey, queue.keyToIdx[key])
	delete(queue.keyToIdx, key)
	return element
}

func (queue _IndexPriorityQueue[T]) PeekKey() int {
	return queue.idxToKey[0]
}

func (queue _IndexPriorityQueue[T]) HasKey(key int) bool {
	_, hasKey := queue.keyToIdx[key]
	return hasKey
}

// IndexPriorityQueue is index priority queue data structure.
type IndexPriorityQueue[T any] interface {
	// Size gets number of elements in a queue.
	Size() int
	// Insert adds element to a queue.
	Insert(key int, element T)
	// Remove removes element from a queue.
	Remove() (T, int)
	// Peek returns first element of a queue.
	Peek() T
	// RemoveByKey removes element by key.
	RemoveByKey(key int) T
	// PeekKey returns first element key.
	PeekKey() int
	// HasKey tells if key exists.
	HasKey(key int) bool
}

// New creates instance of IndexPriorityQueue
func New[T any](less LessFunc[T]) IndexPriorityQueue[T] {
	if less == nil {
		panic("less func is nil")
	}
	return &_IndexPriorityQueue[T]{
		items:    nil,
		keyToIdx: map[int]int{},
		idxToKey: map[int]int{},
		less:     less,
	}
}
