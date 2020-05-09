package list

import (
	"fmt"
	"github.com/tiandi111/ds/test"
	"testing"
)

//TODO: check dlinkedlist inner structure(correctness of prev, next pointers)

func TestNewGenericDoublyLinkedList(t *testing.T) {
	test.AssertNonNil(t, NewGenericDoublyLinkedList())
}

func TestGenericDoublyLinkedList_Add(t *testing.T) {
	dl := NewGenericDoublyLinkedList()
	for i := 0; i < 10; i++ {
		dl.Add(i)
		test.Assert(t, i+1, dl.Len())
		test.Assert(t, i, dl.tail.Value().(int))
		cur := dl.head
		for j := 0; j <= i; j++ {
			test.Assert(t, j, cur.Value())
			cur = cur.Next()
		}
	}
}

func TestGenericDoublyLinkedList_InsertAfter(t *testing.T) {

}

func TestGenericDoublyLinkedList_Get_Overflow(t *testing.T) {
	dl := NewGenericDoublyLinkedList()
	defer func() {
		r := recover()
		test.Assert(t, r, fmt.Sprintf("index out of bound: 0"))
	}()
	dl.Get(0)
}

func TestGenericDoublyLinkedList_Get(t *testing.T) {
	dl := NewGenericDoublyLinkedList()
	for i := 0; i < 10; i++ {
		dl.Add(i)
		for j := 0; j <= i; j++ {
			test.Assert(t, j, dl.Get(j))
		}
	}
}

func TestGenericDoublyLinkedList_Del_Head(t *testing.T) {
	dl := NewGenericDoublyLinkedList()
	for i := 0; i < 10; i++ {
		dl.Add(i)
	}
	defer func() {
		r := recover()
		test.Assert(t, r, fmt.Sprintf("index out of bound: 0"))
	}()
	for i := 0; i < 10; i++ {
		dl.Del(0)
		test.Assert(t, 10-i-1, dl.Len())
		test.Assert(t, i+1, dl.Get(0))
	}
}
