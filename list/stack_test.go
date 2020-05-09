package list

import (
	"fmt"
	"github.com/tiandi111/ds/test"
	"testing"
)

func TestNewGenericStack(t *testing.T) {
	stack := NewGenericStack(0)
	test.AssertNonNil(t, stack)
	test.AssertNonNil(t, stack.arr)
}

func TestStack(t *testing.T) {
	stack := NewGenericStack(0)
}

func TestGenericStack_Push(t *testing.T) {
	stack := NewGenericStack(10)
	for i := 0; i < 100; i++ {
		stack.Push(i)
	}
	for i := 0; i < 100; i++ {
		test.Assert(t, i, stack.arr[i])
	}
}

func TestGenericStack_Pop(t *testing.T) {
	stack := NewGenericStack(10)
	for i := 0; i < 100; i++ {
		stack.Push(i)
	}
	for i := 99; i >= 0; i-- {
		test.Assert(t, i, stack.Pop())
		test.Assert(t, i, stack.Size())
	}
}

func TestGenericStack_PopEmptyStack(t *testing.T) {
	stack := NewGenericStack(0)
	defer func() {
		test.Assert(t, fmt.Errorf("stack is empty"), recover())
	}()
	stack.Pop()
}
