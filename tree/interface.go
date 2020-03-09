package tree

import "github.com/tiandi111/ds"

type SortedTree interface {
	Insert(ds.Comparable)
	Remove(ds.Comparable) bool
	Find(ds.Comparable) interface{}
	RangeFind(from, to ds.Comparable) []interface{}
	NewIterator() ds.Iterator
}
