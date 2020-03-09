package tree

import (
	"fmt"
	"github.com/tiandi111/ds"
	"github.com/tiandi111/ds/test"
	"math/rand"
	"testing"
	"time"
)

type cpb struct {
	int
}

func (c cpb) CompareTo(other ds.Comparable) int {
	return c.int - other.(cpb).int
}

func TestNewGenericBinarySearchTree(t *testing.T) {
	test.AssertNonNil(t, NewGenericBinarySearchTree())
}

func TestGenericBinarySearchTree_Find(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	test.AssertNil(t, tree.Find(cpb{1}))
}

func TestGenericBinarySearchTree_Find2(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	root, left, right := newBstNode(cpb{2}), newBstNode(cpb{1}), newBstNode(cpb{3})
	root.left, root.right = left, right
	tree.root = root
	test.AssertTrue(t, tree.Find(cpb{1}).(cpb).CompareTo(left.value()) == 0)
	test.AssertTrue(t, tree.Find(cpb{2}).(cpb).CompareTo(root.value()) == 0)
	test.AssertTrue(t, tree.Find(cpb{3}).(cpb).CompareTo(right.value()) == 0)
	test.AssertTrue(t, tree.Find(cpb{4}) == nil)
}

func TestGenericBinarySearchTree_Insert(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	tree.Insert(cpb{0})
	test.AssertTrue(t, tree.root.val.CompareTo(cpb{0}) == 0)
}

// Random test
func TestGenericBinarySearchTree_Insert2(t *testing.T) {
	rand.Seed(time.Now().Unix())
	tree := NewGenericBinarySearchTree()
	for i := 0; i < 1000; i++ {
		c := cpb{rand.Intn(100)}
		tree.Insert(c)
		test.AssertTrue(t, tree.Find(c).(cpb).CompareTo(c) == 0)
		test.AssertTrue(t, tree.validate(false))
	}
}

func TestGenericBinarySearchTree_Remove(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	test.AssertFalse(t, tree.Remove(cpb{1}))
}

func TestGenericBinarySearchTree_Remove1(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	tree.root = newBstNode(cpb{1})
	test.AssertTrue(t, tree.Remove(cpb{1}))
	test.AssertNil(t, tree.root)
}

func TestGenericBinarySearchTree_Remove2(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	tree.root = newBstNode(cpb{1})
	tree.root.right = newBstNode(cpb{2})
	test.AssertTrue(t, tree.Remove(cpb{2}))
	test.AssertNil(t, tree.root.right)
}

// Random test
func TestGenericBinarySearchTree_Remove3(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 1; i < 1000; i++ {
		// the size is i
		tree := generateRandomBST(i)
		for j := 0; j < i+1; j++ {
			c := cpb{j}
			//log.Printf("the %dth random test, the %dth removal, expected %dth element to be removed", i, j, c.int)
			tree.Remove(c)
			test.AssertNil(t, tree.Find(c), fmt.Sprintf("the %dth random test, the %dth removal, expected %dth element to be removed", i, j, c.int))
			test.Assert(t, true, tree.validate(false))
		}
	}
}

func TestGenericBinarySearchTree_validate(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	root, left, right := newBstNode(cpb{2}), newBstNode(cpb{1}), newBstNode(cpb{3})
	root.left, root.right = left, right
	tree.root = root
	test.AssertTrue(t, tree.validate(false))
}

func TestGenericBinarySearchTree_validate2(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	root, left, right := newBstNode(cpb{2}), newBstNode(cpb{3}), newBstNode(cpb{3})
	root.left, root.right = left, right
	tree.root = root
	test.AssertFalse(t, tree.validate(false))
}

func TestGenericBinarySearchTree_generateRandomBST(t *testing.T) {
	tree := generateRandomBST(100)
	test.AssertTrue(t, tree.validate(false))
}

func generateRandomBST(size int) *GenericBinarySearchTree {
	rand.Seed(time.Now().Unix())
	tree := NewGenericBinarySearchTree()
	if size < 0 {
		panic("tree size must be non-negative integer")
	}
	tree.root = generateRandomNode(0, size)
	return tree
}

func generateRandomNode(from, to int) *bstNode {
	if from >= to {
		return nil
	}
	val := from + int(rand.Int31n(int32(to-from)))
	node := newBstNode(cpb{val})
	node.left = generateRandomNode(from, val)
	node.right = generateRandomNode(val+1, to)
	return node
}
