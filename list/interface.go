package list

import "github.com/tiandi111/ds"

type List interface {
	Add(interface{})
	Get(int) interface{}
	Del(int) interface{}
	Len() int
	NewIterator() ds.Iterator
}

type Queue interface {
	Append(comparable ds.Comparable)
	First() ds.Comparable
	DelFirst() ds.Comparable
	Size() int
	NewIterator() ds.Iterator
}

type Stack interface {
	Push(comparable ds.Comparable)
	Top() ds.Comparable
	Pop() ds.Comparable
	Size() int
	NewIterator() ds.Iterator
}
