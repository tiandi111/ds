package list

import "github.com/tiandi111/ds"

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

func (l *GenericArrayList) Add(v ds.Comparable) {
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

func (l *GenericArrayList) Find(v ds.Comparable) int {
	lo, hi := 0, l.Len()
	for lo < hi {
		mid := lo + (hi-lo)/2
		pivot := l.Get(mid).(ds.Comparable)
		ret := v.CompareTo(pivot)
		if ret > 0 {
			lo = mid + 1
		} else if ret < 0 {
			hi = mid
		} else {
			return mid
		}
	}
	return lo
}

func (l *GenericArrayList) swap(i, j int) {
	(*l.arr)[i], (*l.arr)[j] = (*l.arr)[j], (*l.arr)[i]
}

func (l *GenericArrayList) Sort() *GenericArrayList {
	l.sortUntil(0, l.Len())
	return l
}

func (l *GenericArrayList) sortUntil(st, end int) {
	if st >= end-1 {
		return
	}
	p := l.pivot(0, end)
	l.sortUntil(0, p)
	l.sortUntil(p+1, end)
}

func (l *GenericArrayList) pivot(st, end int) int {
	if st >= end-1 {
		return st
	}
	pivot := l.Get(end - 1).(ds.Comparable)
	i, j := st, end-2
	for i <= j {
		if l.Get(i).(ds.Comparable).CompareTo(pivot) < 0 {
			i++
		} else {
			l.swap(i, j)
			j--
		}
	}
	l.swap(i, end-1)
	return i
}

func (l *GenericArrayList) NewIterator() *GenericArrayListIterator {
	return &GenericArrayListIterator{-1, l.arr}
}

// todo: test iterator
type GenericArrayListIterator struct {
	index int
	arr   *[]interface{}
}

func (i *GenericArrayListIterator) HasNext() bool {
	return i.index <= len(*i.arr)-1
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
