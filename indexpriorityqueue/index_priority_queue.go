package indexpriorityqueue

import (
	"container/heap"
)

// LessFunc is an ordering function.
type LessFunc func(lhs, rhs interface{}) bool

func (queue _IndexPriorityQueue) Len() int {
	return len(queue.items)
}

func (queue _IndexPriorityQueue) Swap(i, j int) {
	queue.items[i], queue.items[j] = queue.items[j], queue.items[i]
	iKey, jKey := queue.idxToKey[i], queue.idxToKey[j]
	queue.keyToIdx[iKey], queue.keyToIdx[jKey] = j, i
	queue.idxToKey[i], queue.idxToKey[j] = jKey, iKey
}

func (queue _IndexPriorityQueue) Less(i, j int) bool {
	return queue.less(queue.items[i], queue.items[j])
}

func (queue *_IndexPriorityQueue) Push(item interface{}) {
	queue.items = append(queue.items, item)
}

func (queue *_IndexPriorityQueue) Pop() interface{} {
	last := queue.Len() - 1
	item := queue.items[last]
	queue.items[last] = nil
	queue.items = queue.items[0:last]
	return item
}

type _IndexPriorityQueue struct {
	items    []interface{}
	keyToIdx map[int]int
	idxToKey map[int]int
	less     LessFunc
}

func (queue _IndexPriorityQueue) Size() int {
	return queue.Len()
}

func (queue *_IndexPriorityQueue) Insert(key int, element interface{}) {
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

func (queue *_IndexPriorityQueue) Remove() (interface{}, int) {
	if queue.Len() == 0 {
		panic("queue is empty")
	}
	key := queue.idxToKey[0]
	element := heap.Pop(queue)
	delete(queue.idxToKey, queue.keyToIdx[key])
	delete(queue.keyToIdx, key)
	return element, key
}

func (queue _IndexPriorityQueue) Peek() interface{} {
	if queue.Len() == 0 {
		panic("queue is empty")
	}
	return queue.items[0]
}

func (queue *_IndexPriorityQueue) RemoveByKey(key int) interface{} {
	idx := queue.keyToIdx[key]
	element := heap.Remove(queue, idx)
	delete(queue.idxToKey, queue.keyToIdx[key])
	delete(queue.keyToIdx, key)
	return element
}

func (queue _IndexPriorityQueue) PeekKey() int {
	return queue.idxToKey[0]
}

func (queue _IndexPriorityQueue) HasKey(key int) bool {
	_, hasKey := queue.keyToIdx[key]
	return hasKey
}

// IndexPriorityQueue is index priority queue data structure.
type IndexPriorityQueue interface {
	// Size gets number of elements in a queue.
	Size() int
	// Insert adds element to a queue.
	Insert(key int, element interface{})
	// Remove removes element from a queue.
	Remove() (interface{}, int)
	// Peek returns first element of a queue.
	Peek() interface{}
	// RemoveByKey removes element by key.
	RemoveByKey(key int) interface{}
	// PeekKey returns first element key.
	PeekKey() int
	// HasKey tells if key exists.
	HasKey(key int) bool
}

// New creates instance of IndexPriorityQueue
func New(less LessFunc) IndexPriorityQueue {
	if less == nil {
		panic("less func is nil")
	}
	return &_IndexPriorityQueue{
		items:    nil,
		keyToIdx: map[int]int{},
		idxToKey: map[int]int{},
		less:     less,
	}
}
