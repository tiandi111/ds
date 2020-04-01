package tree

import (
	"github.com/tiandi111/ds/test"
	"testing"
)

func TestAvlNode_RotateLeft_NilRightChild(t *testing.T) {
	n := new(avlNode)
	test.Assert(t, n, n.rotateLeft())
}

//  ①         ②
//   \        / \
//   ②  =>  ①  ③
//    \
//    ③
func TestAvlNode_RotateLeft(t *testing.T) {
	root := newAVLNode(test.Cpb{1})
	root.bal = -2
	right := newAVLNode(test.Cpb{2})
	right.bal = -1
	rightOfRight := newAVLNode(test.Cpb{3})
	root.right = right
	root.right.right = rightOfRight

	newRoot := root.rotateLeft()
	test.Assert(t, right, newRoot)
	test.Assert(t, root, newRoot.left)
	test.Assert(t, rightOfRight, newRoot.right)
	test.Assert(t, 0, newRoot.bal)
	test.Assert(t, 0, newRoot.left.bal)
	test.Assert(t, 0, newRoot.right.bal)
}

// ①        ③
//  \       /
//  ③  => ①
//  /       \
//②         ②
func TestAvlNode_RotateLeft2(t *testing.T) {
	root := newAVLNode(test.Cpb{1})
	root.bal = -2
	right := newAVLNode(test.Cpb{3})
	right.bal = 1
	leftOfRight := newAVLNode(test.Cpb{2})
	root.right = right
	root.right.left = leftOfRight

	newRoot := root.rotateLeft()
	test.Assert(t, right, newRoot)
	test.Assert(t, root, newRoot.left)
	test.Assert(t, leftOfRight, newRoot.left.right)
	test.Assert(t, 2, newRoot.bal)
	test.Assert(t, -1, newRoot.left.bal)
	test.Assert(t, 0, newRoot.left.right.bal)
}

func TestAvlNode_RotateRight_NilLeftChild(t *testing.T) {
	n := new(avlNode)
	test.Assert(t, n, n.rotateRight())
}

//     ③      ②
//    /       / \
//   ②  =>  ①  ③
//  /
// ①
func TestAvlNode_RotateRight(t *testing.T) {
	root := newAVLNode(test.Cpb{3})
	root.bal = 2
	left := newAVLNode(test.Cpb{2})
	left.bal = 1
	leftOfLeft := newAVLNode(test.Cpb{1})
	root.left = left
	root.left.left = leftOfLeft

	newRoot := root.rotateRight()
	test.Assert(t, left, newRoot)
	test.Assert(t, leftOfLeft, newRoot.left)
	test.Assert(t, root, newRoot.right)
	test.Assert(t, 0, newRoot.bal)
	test.Assert(t, 0, newRoot.left.bal)
	test.Assert(t, 0, newRoot.right.bal)
}

//   ③   ①
//  /      \
// ①  =>  ③
//  \     /
//  ②   ②
func TestAvlNode_RotateRight2(t *testing.T) {
	root := newAVLNode(test.Cpb{3})
	root.bal = 2
	left := newAVLNode(test.Cpb{1})
	left.bal = -1
	rightOfLeft := newAVLNode(test.Cpb{2})
	root.left = left
	root.left.right = rightOfLeft

	newRoot := root.rotateRight()
	test.Assert(t, left, newRoot)
	test.Assert(t, root, newRoot.right)
	test.Assert(t, rightOfLeft, newRoot.right.left)
	test.Assert(t, -2, newRoot.bal)
	test.Assert(t, 1, newRoot.right.bal)
	test.Assert(t, 0, newRoot.right.left.bal)
}
