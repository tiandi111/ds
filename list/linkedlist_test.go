package list

import (
	"github.com/tiandi111/ds/test"
	"testing"
)

func TestNewGenericLinkedList(t *testing.T) {
	test.AssertNonNil(t, NewGenericLinkedList())
}

func TestGenericLinkedList_Len(t *testing.T) {
	list := NewGenericLinkedList()
	test.Assert(t, 0, list.Len())
	list.Add(1)
	test.Assert(t, 1, list.Len())
}

func TestGenericLinkedList_Add(t *testing.T) {
	list := NewGenericLinkedList()
	list.Add(1)
	list.Add(2)
	test.Assert(t, 2, list.len)
	test.Assert(t, 1, list.head.val.(int))
	test.Assert(t, 2, list.head.next.val.(int))
}

func TestGenericLinkedList_Get(t *testing.T) {
	defer func() {
		got := recover()
		test.AssertNonNil(t, got)
		test.Assert(t, "index out of bound: 0", got)
	}()
	list := NewGenericLinkedList()
	list.Get(0)
}

func TestGenericLinkedList_Get2(t *testing.T) {
	list := NewGenericLinkedList()
	list.Add(1)
	test.Assert(t, 1, list.Get(0))
}

func TestGenericLinkedList_Del(t *testing.T) {
	defer func() {
		got := recover()
		test.AssertNonNil(t, got)
		test.Assert(t, "index out of bound: 0", got)
	}()
	list := NewGenericLinkedList()
	list.Del(0)
}

func TestGenericLinkedList_Del2(t *testing.T) {
	list := NewGenericLinkedList()
	list.Add(1)
	list.Del(0)
	test.Assert(t, 0, list.Len())
}

func TestGenericLinkedList_Del3(t *testing.T) {
	list := NewGenericLinkedList()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Del(1)
	test.Assert(t, 2, list.Len())
	test.Assert(t, 1, list.Get(0))
	test.Assert(t, 3, list.Get(1))
}

func TestGenericLinkedList_NewIterator(t *testing.T) {
	list := NewGenericLinkedList()
	iter := list.NewIterator()
	test.AssertNonNil(t, iter)
}

func TestGenericLinkedListIterator_HasNext(t *testing.T) {
	iter := NewGenericLinkedList().NewIterator()
	test.Assert(t, false, iter.HasNext())
}

func TestGenericLinkedListIterator_HasNext2(t *testing.T) {
	list := NewGenericLinkedList()
	list.Add(1)
	iter := list.NewIterator()
	test.Assert(t, true, iter.HasNext())
}

func TestGenericLinkedListIterator_Next(t *testing.T) {
	list := NewGenericLinkedList()
	list.Add(1)
	iter := list.NewIterator()
	test.Assert(t, 1, iter.Next().GetValue())
}

func TestGenericLinkedListIterator_Next2(t *testing.T) {
	defer func() {
		got := recover()
		test.AssertNil(t, got)
	}()
	list := NewGenericLinkedList()
	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	iter := list.NewIterator()
	for iter.HasNext() {
		iter.Next()
	}
}

func TestGenericLinkedListIterator_GetValue(t *testing.T) {
	defer func() {
		got := recover()
		test.AssertNonNil(t, got)
		test.Assert(t, "call Next() to get the first element", got)
	}()
	iter := NewGenericLinkedList().NewIterator()
	iter.GetValue()
}

func TestGenericLinkedListIterator_GetValue2(t *testing.T) {
	list := NewGenericLinkedList()
	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	iter := list.NewIterator()
	for iter.HasNext() {
		val := iter.Next().GetValue()
		test.Assert(t, iter.index, val)
	}
}

func BenchmarkGenericLinkedList_Del_Middle(b *testing.B) {
	lList := NewGenericLinkedList()
	for i := 0; i < BenchmarkListSize; i++ {
		lList.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < BenchmarkDelReps; i++ {
		lList.Del(lList.Len() * BenchmarkDelMiddlePosition / BenchmarkDelDeNominator)
	}
}

func BenchmarkGenericLinkedList_Del_Head(b *testing.B) {
	lList := NewGenericLinkedList()
	for i := 0; i < BenchmarkListSize; i++ {
		lList.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < BenchmarkDelReps; i++ {
		lList.Del(lList.Len() * BenchmarkDelHeadPosition / BenchmarkDelDeNominator)
	}
}

func BenchmarkGenericLinkedList_Del_Tail(b *testing.B) {
	lList := NewGenericLinkedList()
	for i := 0; i < BenchmarkListSize; i++ {
		lList.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < BenchmarkDelReps; i++ {
		lList.Del(lList.Len() * BenchmarkDelTailPosition / BenchmarkDelDeNominator)
	}
}
