package heap

import (
	"github.com/tiandi111/ds"
	"github.com/tiandi111/ds/list"
)

type GenericArrayHeap struct {
	arr []ds.Comparable
}

func NewGenericArrayHeap() *GenericArrayHeap {
	h := &GenericArrayHeap{make([]ds.Comparable, 0, 8)}
	return h
}

func (h *GenericArrayHeap) Insert(c ds.Comparable) {
	h.insert(c)
}

func (h *GenericArrayHeap) Max() interface{} {
	return h.max()
}

func (h *GenericArrayHeap) DelMax() interface{} {
	return h.delMax()
}

func (h *GenericArrayHeap) Size() int {
	return len(h.arr)
}

func (h *GenericArrayHeap) NewIterator(c ds.Comparable) ds.Iterator {
	return nil
}

func (h *GenericArrayHeap) insert(c ds.Comparable) {
	h.arr = append(h.arr, c)
	up(h.arr, len(h.arr)-1)
}

func (h *GenericArrayHeap) max() interface{} {
	if h.Size() == 0 {
		return nil
	}
	return h.arr[0]
}

func (h *GenericArrayHeap) delMax() interface{} {
	if h.Size() == 0 {
		return nil
	}
	max := h.arr[0]
	swap(h.arr, 0, len(h.arr)-1)
	h.arr = h.arr[:len(h.arr)-1]
	down(h.arr, 0)
	return max
}

func up(arr []ds.Comparable, idx int) {
	cur := idx
	for cur >= 0 && cur < len(arr) {
		parent := (cur - 1) / 2
		if parent >= 0 && arr[cur].CompareTo(arr[parent]) > 0 {
			swap(arr, cur, parent)
			cur = parent
		} else {
			return
		}
	}
}

func down(arr []ds.Comparable, idx int) {
	cur := idx
	for cur >= 0 && cur < len(arr) {
		left := cur*2 + 1
		right := cur*2 + 2
		switch {
		// when cur element less than its left child,
		// we choose the max of its children and swap
		case left < len(arr) && arr[cur].CompareTo(arr[left]) < 0:
			next := left
			if right < len(arr) && arr[left].CompareTo(arr[right]) < 0 {
				next = right
			}
			swap(arr, cur, next)
			cur = next

		case right < len(arr) && arr[cur].CompareTo(arr[right]) < 0:
			swap(arr, cur, right)
			cur = right

		default:
			return
		}
	}
}

func swap(arr []ds.Comparable, a, b int) {
	if a < 0 || a >= len(arr) {
		panic(list.ErrIndexOutOfBound(a))
	}
	if b < 0 || b >= len(arr) {
		panic(list.ErrIndexOutOfBound(b))
	}
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}
