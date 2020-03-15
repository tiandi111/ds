package heap

import "github.com/tiandi111/ds"

type MaxHeap interface {
	Insert(comparable ds.Comparable)
	Max() interface{}
	DelMax() interface{}
	Size() int
	NewIterator() ds.Iterator
}

type MinHeap interface {
	Insert(comparable ds.Comparable)
	Min() interface{}
	DelMin() interface{}
	Size() int
	NewIterator() ds.Iterator
}

type MergeableHeap interface {
	MinHeap
	Union(heap MergeableHeap)
}
