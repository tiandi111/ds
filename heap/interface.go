package heap

import "github.com/tiandi111/ds"

// max heap
type Heap interface {
	Insert(comparable ds.Comparable)
	// get the peak value
	Max() interface{}
	DelMax() interface{}
	Size() int
	NewIterator() ds.Iterator
}
