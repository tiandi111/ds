package list

import (
	"github.com/tiandi111/ds/test"
	"testing"
)

func TestNewGenericArrayList(t *testing.T) {
	test.AssertNonNil(t, NewGenericArrayList())
}

func TestGenericArrayList_Add(t *testing.T) {
	list := NewGenericArrayList()
	list.Add(1)
	test.Assert(t, 1, list.Len())
}

func TestGenericArrayList_Len(t *testing.T) {
	test.Assert(t, 0, NewGenericArrayList().Len())
}

func TestGenericArrayList_Get(t *testing.T) {
	list := NewGenericArrayList()
	list.Add(1)
	test.Assert(t, 1, list.Get(0))
}

func TestGenericArrayList_Del(t *testing.T) {
	defer func() {
		got := recover()
		test.Assert(t, ErrIndexOutOfBound(0).Error(), got.(error).Error())
	}()
	list := NewGenericArrayList()
	list.Del(0)
}

func TestGenericArrayList_Del2(t *testing.T) {
	list := NewGenericArrayList()
	list.Add(1)
	list.Del(0)
	test.Assert(t, 0, list.Len())
}

func TestGenericArrayList_Del3(t *testing.T) {
	defer func() {
		got := recover()
		test.Assert(t, ErrIndexOutOfBound(-1).Error(), got.(error).Error())
	}()
	list := NewGenericArrayList()
	list.Del(-1)

}

func BenchmarkGenericArrayList_Del_Middle(b *testing.B) {
	aList := NewGenericArrayList()
	for i := 0; i < BenchmarkListSize; i++ {
		aList.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < BenchmarkDelReps; i++ {
		aList.Del(aList.Len() * BenchmarkDelMiddlePosition / BenchmarkDelDeNominator)
	}
}

func BenchmarkGenericArrayList_Del_Head(b *testing.B) {
	aList := NewGenericArrayList()
	for i := 0; i < BenchmarkListSize; i++ {
		aList.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < BenchmarkDelReps; i++ {
		aList.Del(aList.Len() * BenchmarkDelHeadPosition / BenchmarkDelDeNominator)
	}
}

func BenchmarkGenericArrayList_Del_Tail(b *testing.B) {
	aList := NewGenericArrayList()
	for i := 0; i < BenchmarkListSize; i++ {
		aList.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < BenchmarkDelReps; i++ {
		aList.Del(aList.Len() * BenchmarkDelTailPosition / BenchmarkDelDeNominator)
	}
}
