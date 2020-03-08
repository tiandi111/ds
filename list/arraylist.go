package list

type GenericArrayList struct {
	arr *[]interface{}
}

func NewGenericArrayList() *GenericArrayList {
	arr := make([]interface{}, 0)
	return &GenericArrayList{&arr}
}

func (l *GenericArrayList) Len() int {
	return len(*l.arr)
}

func (l *GenericArrayList) Add(v interface{}) {
	*l.arr = append(*l.arr, v)
}

func (l *GenericArrayList) Get(n int) interface{} {
	if n < 0 || n >= l.Len() {
		panic(ErrIndexOutOfBound(n))
	}
	return (*l.arr)[n]
}

func (l *GenericArrayList) Del(n int) interface{} {
	if n < 0 || n >= l.Len() {
		panic(ErrIndexOutOfBound(n))
	}
	*l.arr = append((*l.arr)[:n], (*l.arr)[n+1:]...)
	return nil
}

func (l *GenericArrayList) NewIterator() *GenericArrayListIterator {
	return &GenericArrayListIterator{-1, l.arr}
}

type GenericArrayListIterator struct {
	index int
	arr   *[]interface{}
}

func (i *GenericArrayListIterator) HasNext() bool {
	return i.index < len(*i.arr)-1
}

func (i *GenericArrayListIterator) Next() *GenericArrayListIterator {
	i.index++
	if !i.HasNext() {
		panic(ErrIndexOutOfBound(i.index))
	}
	return i
}

func (i *GenericArrayListIterator) GetValue() interface{} {
	return (*i.arr)[i.index]
}
