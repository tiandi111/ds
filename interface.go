package ds

type Comparable interface {
	// return
	// this > other if this.CompareTo(other) > 0
	// this = other if this.CompareTo(other) == 0
	// this < otehr if this.CompareTo(other) < 0
	CompareTo(other Comparable) int
}

type Iterator interface {
	HasNext() bool
	Next() Iterator
	Index() int
	GetValue() interface{}
}
