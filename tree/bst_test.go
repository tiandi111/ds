package tree

import (
	"fmt"
	"github.com/tiandi111/ds/test"
	"math/rand"
	"testing"
	"time"
)

func TestNewGenericBinarySearchTree(t *testing.T) {
	test.AssertNonNil(t, NewGenericBinarySearchTree())
}

func TestGenericBinarySearchTree_Find(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	test.AssertNil(t, tree.Find(test.Cpb{1}))
}

func TestGenericBinarySearchTree_Find2(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	root, left, right := newBstNode(test.Cpb{2}), newBstNode(test.Cpb{1}), newBstNode(test.Cpb{3})
	root.left, root.right = left, right
	tree.root = root
	test.AssertTrue(t, tree.Find(test.Cpb{1}).(test.Cpb).CompareTo(left.value()) == 0)
	test.AssertTrue(t, tree.Find(test.Cpb{2}).(test.Cpb).CompareTo(root.value()) == 0)
	test.AssertTrue(t, tree.Find(test.Cpb{3}).(test.Cpb).CompareTo(right.value()) == 0)
	test.AssertTrue(t, tree.Find(test.Cpb{4}) == nil)
}

func TestGenericBinarySearchTree_Insert(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	tree.Insert(test.Cpb{0})
	test.AssertTrue(t, tree.root.val.CompareTo(test.Cpb{0}) == 0)
	test.Assert(t, 1, tree.Size())
}

// Random test
func TestGenericBinarySearchTree_Insert2(t *testing.T) {
	rand.Seed(time.Now().Unix())
	tree := NewGenericBinarySearchTree()
	for i := 0; i < 1000; i++ {
		c := test.Cpb{-i + rand.Intn(2*i+1)}
		tree.Insert(c)
		test.AssertTrue(t, tree.Find(c).(test.Cpb).CompareTo(c) == 0)
		test.AssertTrue(t, tree.validate(false))
		test.Assert(t, i+1, tree.Size())
	}
}

func TestGenericBinarySearchTree_Remove(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	test.AssertNil(t, tree.Remove(test.Cpb{1}))
	test.Assert(t, 0, tree.Size())
}

func TestGenericBinarySearchTree_Remove1(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	tree.root = newBstNode(test.Cpb{1})
	test.Assert(t, 1, tree.Remove(test.Cpb{1}).(test.Cpb).Val)
	test.AssertNil(t, tree.root)
}

func TestGenericBinarySearchTree_Remove2(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	tree.root = newBstNode(test.Cpb{1})
	tree.root.right = newBstNode(test.Cpb{2})
	test.Assert(t, 1, tree.Remove(test.Cpb{1}).(test.Cpb).Val)
	test.AssertNil(t, tree.root.right)
	test.AssertTrue(t, tree.validate(false))
}

// Random test
func TestGenericBinarySearchTree_Remove3(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 5000; i++ {
		// the size is i
		tree := generateRandomBST(i)
		expectedSize := i
		for j := 0; j < i+1; j++ {
			c := test.Cpb{j}
			msg := fmt.Sprintf("the %dth random test, the %dth removal, expected %dth element to be removed", i, j, c.Val)
			tree.Remove(c)
			test.AssertNil(t, tree.Find(c), msg)
			test.Assert(t, true, tree.validate(false), msg)
			if c.Val >= 0 && c.Val < i {
				expectedSize--
			}
			test.Assert(t, expectedSize, tree.Size(), msg)
		}
	}
}

func TestGenericBinarySearchTree_RangeFind(t *testing.T) {
	defer func() {
		test.Assert(t, ErrInvalidRangeFindArgs, recover())
	}()
	tree := NewGenericBinarySearchTree()
	tree.RangeFind(test.Cpb{1}, test.Cpb{0})
}

func TestGenericBinarySearchTree_RangeFind2(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	data := tree.RangeFind(test.Cpb{0}, test.Cpb{1})
	test.Assert(t, 0, len(data))
}

// Random test
func TestGenericBinarySearchTree_RangeFind3(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 1; i < 100; i++ {
		tree := generateRandomBST(i)
		from := rand.Intn(i)
		to := from + rand.Intn(i-from)
		data := tree.RangeFind(test.Cpb{from}, test.Cpb{to})
		for j := 0; j < to-from; j++ {
			test.Assert(t, from+j, data[j].(test.Cpb).Val)
		}
	}
}

func TestGenericBinarySearchTree_validate(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	root, left, right := newBstNode(test.Cpb{2}), newBstNode(test.Cpb{1}), newBstNode(test.Cpb{3})
	root.left, root.right = left, right
	tree.root = root
	test.AssertTrue(t, tree.validate(false))
}

func TestGenericBinarySearchTree_validate2(t *testing.T) {
	tree := NewGenericBinarySearchTree()
	root, left, right := newBstNode(test.Cpb{2}), newBstNode(test.Cpb{3}), newBstNode(test.Cpb{3})
	root.left, root.right = left, right
	tree.root = root
	test.AssertFalse(t, tree.validate(false))
}

func TestGenericBinarySearchTree_generateRandomBST(t *testing.T) {
	tree := generateRandomBST(100)
	test.AssertTrue(t, tree.validate(false))
}

// Generate a random bst
// to reduce dependency, we construct bst directly instead of using Insert
func generateRandomBST(size int) *GenericBinarySearchTree {
	rand.Seed(time.Now().Unix())
	tree := NewGenericBinarySearchTree()
	if size < 0 {
		panic("tree size must be non-negative integer")
	}
	tree.root = generateRandomNode(0, size)
	tree.size = size
	return tree
}

func generateRandomNode(from, to int) *bstNode {
	if from >= to {
		return nil
	}
	val := from + int(rand.Int31n(int32(to-from)))
	node := newBstNode(test.Cpb{val})
	node.left = generateRandomNode(from, val)
	node.right = generateRandomNode(val+1, to)
	return node
}

func printBST(root *bstNode) {
	queue := make([]*bstNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == nil {
				fmt.Printf("nil ")
				continue
			}
			fmt.Printf("%v ", cur.val)
			queue = append(queue, cur.left)
			queue = append(queue, cur.right)
		}
		fmt.Printf("\n")
	}
}
