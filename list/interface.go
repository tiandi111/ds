package list

type List interface {
	Add(interface{})
	Get(int) interface{}
	Del(int) interface{}
	Len() int
	NewIterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() Iterator
	Index() int
	GetValue() interface{}
}
