package ds

import (
	"fmt"
	"io"
)

type Element interface {
	Comparable
	Serializable
	fmt.Stringer
}

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

type Serializable interface {
	Serialize(io.Writer) error
	Deserialize(io.Reader) error
}

type Deserializer func(io.Reader) (Element, error)
