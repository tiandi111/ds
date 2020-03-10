package tree

import "github.com/tiandi111/ds"

type SortedTree interface {
	Insert(ds.Comparable)
	Remove(ds.Comparable) interface{}
	Find(ds.Comparable) interface{}
	RangeFind(from, to ds.Comparable) []interface{}
	Size() int
	NewIterator() ds.Iterator
}
