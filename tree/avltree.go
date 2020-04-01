package tree

import (
	"github.com/tiandi111/ds"
	"github.com/tiandi111/ds/util"
)

const (
	LH = 1 // left subtree higher
	EH = 0
	RH = -1
)

type GenericAVLTree struct {
	root *avlNode
	size int64
}

func (t *GenericAVLTree) Insert(comparable ds.Comparable) {

}

func (t *GenericAVLTree) insert(c ds.Comparable) {
	t.size++
	if t.root == nil {
		t.root = newAVLNode(c)
		return
	}
	var parent, ancestor *avlNode
	var leftChildOfParent, leftChildOfAncestor bool
	cur := t.root
	for cur != nil {
		if c.CompareTo(cur.val) <= 0 {
			if cur.left == nil {
				cur.left = newAVLNode(c)
				if leftChildOfParent && cur.bal == 0 && parent.bal == 1 {
					parent = parent.rotateRight()
				} else if !leftChildOfParent && cur.bal == 0 && parent.bal == -1 {
					parent.right = cur.rotateRight()
					parent = parent.rotateLeft()
				}
			} else {
				ancestor = parent
				parent = cur
				cur = cur.left
				leftChildOfAncestor = leftChildOfParent
				leftChildOfParent = true
			}
		} else {
			if cur.right == nil {
				cur.right = newAVLNode(c)
				if leftChildOfParent && cur.bal == 0 && parent.bal == 1 {
					parent.left = cur.rotateLeft()
					parent = parent.rotateRight()
				} else if !leftChildOfParent && cur.bal == 0 && parent.bal == -1 {
					parent = parent.rotateLeft()
				}
			} else {
				ancestor = parent
				parent = cur
				cur = cur.right
				leftChildOfAncestor = leftChildOfParent
				leftChildOfParent = false
			}
		}
	}
	if ancestor == nil {
		t.root = parent
	} else if leftChildOfAncestor {
		ancestor.left = parent
	} else {
		ancestor.right = parent
	}
}

func (t *GenericAVLTree) Delete(comparable ds.Comparable) interface{} {
	return nil
}

func (t *GenericAVLTree) Find(comparable ds.Comparable) interface{} {
	return nil
}

func (t *GenericAVLTree) RangeFind(from, to ds.Comparable) {

}

func (t *GenericAVLTree) Size() int64 {
	return t.size
}

func (t *GenericAVLTree) NewIterator() ds.Iterator {
	return nil
}

type avlNode struct {
	// balance factor
	bal   int
	left  *avlNode
	right *avlNode
	val   ds.Comparable
}

func newAVLNode(c ds.Comparable) *avlNode {
	return &avlNode{val: c}
}

func (n *avlNode) insertLeft() {

}

// rotate n to the left and return the new root
func (root *avlNode) rotateLeft() *avlNode {
	if root.right == nil {
		return root
	}
	var rootBal, rightChildBal int
	if root.right.bal >= 0 {
		// root.bal = h0-1-h1
		// right.bal = h1-h2
		rightChildBal = 1 + util.MaxInt(root.bal+root.right.bal+1, root.right.bal) // max(h0+1, h1+1)-h2
		rootBal = root.bal + 1
	} else {
		// root.bal = h0-1-h2
		// right.bal = h1-h2
		rightChildBal = 1 + util.MaxInt(root.bal+1, root.right.bal) // 1+max(h0, h1)-h2
		rootBal = root.bal + 1 - root.right.bal                     // h0-h1
	}
	root.bal = rootBal
	root.right.bal = rightChildBal
	newRoot := root.right
	rightChild := root.right.left
	newRoot.left = root
	root.right = rightChild
	return newRoot
}

// rotate n to the right and return the new root
func (root *avlNode) rotateRight() *avlNode {
	if root.left == nil {
		return root
	}
	var rootBal, leftChildBal int
	if root.left.bal >= 0 {
		// root.bal = 1+h1-h0
		// left.bal = h1-h2
		leftChildBal = util.MinInt(root.left.bal, root.bal-1) - 1 // h1-1+min(-h2, -h0)
		rootBal = root.bal - 1 - root.left.bal                    // h2-h0
	} else {
		// root.bal = 1+h2-h0
		// left.bal = h1-h2
		leftChildBal = util.MinInt(root.left.bal, root.bal+root.left.bal-1) - 1 // h1-1+min(-h2, -h0)
		rootBal = root.bal - 1                                                  // h2-h0
	}
	root.bal = rootBal
	root.left.bal = leftChildBal
	newRoot := root.left
	leftChild := root.left.right
	newRoot.right = root
	root.left = leftChild
	return newRoot
}
