package list

import "github.com/tiandi111/ds"

type List interface {
	Add(interface{})
	Get(int) interface{}
	Del(int) interface{}
	Len() int
	NewIterator() ds.Iterator
}
