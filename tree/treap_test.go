package tree

import (
	"fmt"
	"github.com/tiandi111/ds/test"
	"testing"
)

func validateHeap(node Node) bool {
	real, ok := node.(*treapNode)
	if !ok {
		return false
	}
	if real.left != nil && real.left.priority < real.priority {
		return false
	}
	if real.right != nil && real.right.priority < real.priority {
		return false
	}
	return true
}

func validatePointers(node Node) bool {
	real, ok := node.(*treapNode)
	if !ok {
		return false
	}
	if real.left != nil && real.left.parent != real {
		return false
	}
	if real.right != nil && real.right.parent != real {
		return false
	}
	return true
}

func TestTreap_Insert(t *testing.T) {
	treap := NewGenericTreap()
	for i := 0; i < 1000; i++ {
		treap.Insert(test.Cpb{i})
		test.AssertTrue(t, InOrderValidate(treap.root, validateHeap, validatePointers))
	}
}

func TestGenericTreap_Remove(t *testing.T) {
	treap := NewGenericTreap()
	for i := 0; i < 1000; i++ {
		treap.Insert(test.Cpb{i})
	}
	for i := 0; i < 1000; i++ {
		treap.Remove(test.Cpb{i})
		test.AssertTrue(t, InOrderValidate(treap.root, validateHeap, validatePointers))
		test.Assert(t, int64(1000-i-1), treap.Size())
	}
	test.AssertNil(t, treap.root)
}

func TestGenericTreap_Find(t *testing.T) {
	treap := NewGenericTreap()
	for i := 0; i < 10; i++ {
		treap.Insert(test.Cpb{i})
	}
	for i := -10; i < 20; i++ {
		obj := treap.Find(test.Cpb{i})
		if i >= 0 && i < 10 {
			test.AssertNonNil(t, obj)
		} else {
			test.AssertNil(t, obj)
		}

	}
}

func TestTreapDepth(t *testing.T) {
	treap := NewGenericTreap()
	N := 10000
	for i := 0; i < N; i++ {
		treap.Insert(test.Cpb{i})
	}
	fmt.Println(Depth(treap.root))
}
