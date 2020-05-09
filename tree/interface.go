package tree

import (
	"github.com/tiandi111/ds"
	"github.com/tiandi111/ds/util"
)

type SortedTree interface {
	Insert(ds.Comparable)
	Remove(ds.Comparable) interface{}
	Find(ds.Comparable) interface{}
	RangeFind(from, to ds.Comparable) []interface{}
	Size() int
	NewIterator() ds.Iterator
}

type Node interface {
	Left() Node
	Right() Node
	Value() ds.Comparable
}

func Depth(root Node) int {
	if util.IsNil(root) {
		return 0
	}
	return util.MaxInt(Depth(root.Left()), Depth(root.Right())) + 1
}
