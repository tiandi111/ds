package list

import (
	"fmt"
	"testing"

	"github.com/tiandi111/ds"

	"github.com/tiandi111/ds/test"
)

func TestNewGenericArrayList(t *testing.T) {
	test.AssertNonNil(t, NewGenericArrayList())
}

func TestGenericArrayList_Add(t *testing.T) {
	list := NewGenericArrayList()
	list.Add(test.Cpb{1})
	test.Assert(t, 1, list.Len())
}

func TestGenericArrayList_Len(t *testing.T) {
	test.Assert(t, 0, NewGenericArrayList().Len())
}

func TestGenericArrayList_Get(t *testing.T) {
	list := NewGenericArrayList()
	list.Add(test.Cpb{1})
	test.Assert(t, test.Cpb{1}, list.Get(0))
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
	list.Add(test.Cpb{1})
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

func TestGenericArrayList_Sort(t *testing.T) {
	test1 := []int{}
	test.AssertTrue(t, validate(genComparableArrayList(test1).Sort()))

	test2 := []int{1}
	test.AssertTrue(t, validate(genComparableArrayList(test2).Sort()))

	test3 := []int{2, 1}
	test.AssertTrue(t, validate(genComparableArrayList(test3).Sort()))

	test4 := []int{1, 3, 2}
	test.AssertTrue(t, validate(genComparableArrayList(test4).Sort()))

	test5 := []int{-1, 2}
	test.AssertTrue(t, validate(genComparableArrayList(test5).Sort()))

	test6 := []int{3, -1, -2}
	test.AssertTrue(t, validate(genComparableArrayList(test6).Sort()))

	test7 := []int{3, -1, -2, 100, 29, 56, -200}
	test.AssertTrue(t, validate(genComparableArrayList(test7).Sort()))
}

func validate(l *GenericArrayList) bool {
	defer func() {
		fmt.Printf("%v\n", l.arr)
	}()
	for i := 1; i < l.Len(); i++ {
		last := l.Get(i - 1).(ds.Comparable)
		this := l.Get(i).(ds.Comparable)
		if last.CompareTo(this) > 0 {

			return false
		}
	}
	return true
}

func genComparableArrayList(arr []int) *GenericArrayList {
	l := NewGenericArrayList()
	for _, e := range arr {
		l.Add(test.Cpb{e})
	}
	return l
}

func TestGenericArrayList_Find(t *testing.T) {
	test1 := []int{1, 2, 3}
	test.Assert(t, 1, genComparableArrayList(test1).Find(test.Cpb{2}))
	test.Assert(t, 3, genComparableArrayList(test1).Find(test.Cpb{4}))
	test.Assert(t, 0, genComparableArrayList(test1).Find(test.Cpb{0}))

	test4 := []int{1, 2}
	test.Assert(t, 1, genComparableArrayList(test4).Find(test.Cpb{2}))
	test.Assert(t, 0, genComparableArrayList(test4).Find(test.Cpb{1}))
	test.Assert(t, 0, genComparableArrayList(test4).Find(test.Cpb{-1}))
}

func BenchmarkGenericArrayList_Del_Middle(b *testing.B) {
	aList := NewGenericArrayList()
	for i := 0; i < BenchmarkListSize; i++ {
		aList.Add(test.Cpb{i})
	}
	b.ResetTimer()
	for i := 0; i < BenchmarkDelReps; i++ {
		aList.Del(aList.Len() * BenchmarkDelMiddlePosition / BenchmarkDelDeNominator)
	}
}

func BenchmarkGenericArrayList_Del_Head(b *testing.B) {
	aList := NewGenericArrayList()
	for i := 0; i < BenchmarkListSize; i++ {
		aList.Add(test.Cpb{i})
	}
	b.ResetTimer()
	for i := 0; i < BenchmarkDelReps; i++ {
		aList.Del(aList.Len() * BenchmarkDelHeadPosition / BenchmarkDelDeNominator)
	}
}

func BenchmarkGenericArrayList_Del_Tail(b *testing.B) {
	aList := NewGenericArrayList()
	for i := 0; i < BenchmarkListSize; i++ {
		aList.Add(test.Cpb{i})
	}
	b.ResetTimer()
	for i := 0; i < BenchmarkDelReps; i++ {
		aList.Del(aList.Len() * BenchmarkDelTailPosition / BenchmarkDelDeNominator)
	}
}
