package heap

import (
	"github.com/tiandi111/ds"
	"github.com/tiandi111/ds/list"
	"github.com/tiandi111/ds/test"
	"math/rand"
	"testing"
	"time"
)

const (
	HeapInsertionSize = 100
)

func TestNewGenericArrayHeap(t *testing.T) {
	test.AssertNonNil(t, NewGenericArrayHeap())
}

func TestGenericArrayHeap_Size(t *testing.T) {
	test.Assert(t, 0, NewGenericArrayHeap().Size())
}

func TestGenericArrayHeap_Insert(t *testing.T) {
	h := NewGenericArrayHeap()
	for i := 0; i < HeapInsertionSize; i++ {
		h.Insert(test.Cpb{i})
		test.AssertTrue(t, h.validate())
		test.Assert(t, i+1, h.Size())
		test.Assert(t, i, h.arr[0].(test.Cpb).Val)
		test.Assert(t, 1, h.frequency(test.Cpb{i}))
	}
}

// random test
func TestGenericArrayHeap_Insert_Random_Test(t *testing.T) {
	rand.Seed(time.Now().Unix())
	h := NewGenericArrayHeap()
	m := make(map[int]int)
	for i := 0; i < HeapInsertionSize; i++ {
		val := rand.Intn(2*HeapInsertionSize) - HeapInsertionSize
		m[val]++
		h.Insert(test.Cpb{val})
		test.Assert(t, m[val], h.frequency(test.Cpb{val}))
		test.AssertTrue(t, h.validate())
	}
}

func TestGenericArrayHeap_Max_EmptyHeap(t *testing.T) {
	h := NewGenericArrayHeap()
	test.AssertNil(t, h.Max())
}

func TestGenericArrayHeap_Max(t *testing.T) {
	h := NewGenericArrayHeap()
	for i := 0; i < HeapInsertionSize; i++ {
		h.Insert(test.Cpb{i})
		test.Assert(t, i, h.Max().(test.Cpb).Val)
	}
}

func TestGenericArrayHeap_DelMax_EmptyHeap(t *testing.T) {
	h := NewGenericArrayHeap()
	test.AssertNil(t, h.DelMax())
}

func TestGenericArrayHeap_DelMax(t *testing.T) {
	h := NewGenericArrayHeap()
	for i := 0; i < HeapInsertionSize; i++ {
		h.Insert(test.Cpb{i})
	}
	for i := HeapInsertionSize - 1; i >= 0; i-- {
		test.Assert(t, i, h.DelMax().(test.Cpb).Val)
		test.AssertTrue(t, h.validate())
		test.Assert(t, i, h.Size())
		if i > 0 {
			test.Assert(t, i-1, h.Max().(test.Cpb).Val)
		}
	}
}

func TestUp(t *testing.T) {
	a := test.Cpb{0}
	b := test.Cpb{1}
	c := test.Cpb{2}
	arr := []ds.Comparable{b, a, c}
	up(arr, 2)
	test.Assert(t, c, arr[0])
	test.Assert(t, a, arr[1])
	test.Assert(t, b, arr[2])
}

func TestDown(t *testing.T) {
	a := test.Cpb{1}
	b := test.Cpb{2}
	c := test.Cpb{3}
	arr := []ds.Comparable{a, b, c}
	down(arr, 0)
	test.Assert(t, c, arr[0])
	test.Assert(t, b, arr[1])
	test.Assert(t, a, arr[2])
}

func TestSwap_Overflow1(t *testing.T) {
	defer func() {
		test.Assert(t, list.ErrIndexOutOfBound(0).Error(), recover().(error).Error())
	}()
	arr := make([]ds.Comparable, 0)
	swap(arr, 0, 1)
}

func TestSwap_Overflow2(t *testing.T) {
	defer func() {
		test.Assert(t, list.ErrIndexOutOfBound(1).Error(), recover().(error).Error())
	}()
	arr := make([]ds.Comparable, 1)
	swap(arr, 0, 1)
}

func TestSwap(t *testing.T) {
	a := test.Cpb{0}
	b := test.Cpb{1}
	arr := []ds.Comparable{a, b}
	swap(arr, 0, 1)
	test.AssertTrue(t, arr[0].CompareTo(b) == 0)
	test.AssertTrue(t, arr[1].CompareTo(a) == 0)
}

func (h *GenericArrayHeap) validate() bool {
	for i := range h.arr {
		left := 2*i + 1
		right := 2*i + 2
		if (left < len(h.arr) && h.arr[i].CompareTo(h.arr[left]) < 0) ||
			(right < len(h.arr) && h.arr[i].CompareTo(h.arr[right]) < 0) {
			return false
		}
	}
	return true
}

func (h *GenericArrayHeap) frequency(c ds.Comparable) int {
	f := 0
	for _, val := range h.arr {
		if val == c {
			f++
		}
	}
	return f
}
