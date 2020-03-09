package ds

type Comparable interface {
	// return
	// positive if this>other
	// zero if this == other
	// negative this < other
	CompareTo(other Comparable) int
}

type Iterator interface {
	HasNext() bool
	Next() Iterator
	Index() int
	GetValue() interface{}
}
