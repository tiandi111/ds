package test

import "github.com/tiandi111/ds"

// Cpb implements ds.Comparable
// only used for testing
type Cpb struct {
	Val int
}

func (c Cpb) CompareTo(other ds.Comparable) int {
	return c.Val - other.(Cpb).Val
}
