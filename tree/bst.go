package tree

import (
	"github.com/tiandi111/ds"
	"log"
)

type GenericBinarySearchTree struct {
	root *bstNode
	size int
}

func NewGenericBinarySearchTree() *GenericBinarySearchTree {
	return new(GenericBinarySearchTree)
}

func (t *GenericBinarySearchTree) Insert(c ds.Comparable) {
	i := newBstNode(c)
	if t.root == nil {
		t.root = i
		return
	}
	n := t.root
	for n != nil {
		if n.val.CompareTo(c) >= 0 {
			if n.left == nil {
				n.left = i
				return
			}
			n = n.left
		} else {
			if n.right == nil {
				n.right = i
				return
			}
			n = n.right
		}
	}
}

func (t *GenericBinarySearchTree) Remove(c ds.Comparable) bool {
	return t.remove(c)
}

// delete the top-most child of n that is equal to c
// return true if a node is deleted
// return false if
// 	(1)given root is nil
// 	(2)given value is nil
//	(3)no such child that is equal to c
func (t *GenericBinarySearchTree) remove(c ds.Comparable) bool {
	n := t.root
	if n == nil || c == nil {
		return false
	}
	p, d, isLeftChild := findWithParent(n, c)
	if d == nil {
		return false
	}
	switch {
	case d.left == nil && d.right == nil:
		// node is leaf node
		// parent is nil, node is the root
		if p == nil {
			t.root = nil
		} else {
			if isLeftChild {
				p.left = nil
			} else {
				p.right = nil
			}
		}
	case d.left != nil && d.right != nil:
		// node has left and right child
		// replace it with the left most child of its right child
		succ := delLeftMost(d.right)
		d.val = succ.val
		if d.right.left == nil {
			d.right = d.right.right
		}
	case d.left != nil:
		// node only has left child
		// replace it with its left child
		*d = *d.left
	case d.right != nil:
		// node only has right child
		// replace it with its right child
		*d = *d.right
	}
	return true
}

func (t *GenericBinarySearchTree) Find(c ds.Comparable) interface{} {
	if t.root == nil {
		return nil
	}
	node := find(t.root, c)
	return node.value()
}

func (t *GenericBinarySearchTree) RangeFind(from, to ds.Comparable) []interface{} {
	return nil
}

func (t *GenericBinarySearchTree) validate(enableLog bool) bool {
	if t.root == nil {
		return true
	}
	return t.root.validate(enableLog)
}

type bstNode struct {
	left  *bstNode
	right *bstNode
	val   ds.Comparable
}

func newBstNode(v ds.Comparable) *bstNode {
	return &bstNode{val: v}
}

// binary search the top-most node with value c
// nil value is not supported
func find(n *bstNode, c ds.Comparable) *bstNode {
	_, t, _ := findWithParent(n, c)
	return t
}

// find the top-most node with value c
// return (parent, target, isLeftChild)
func findWithParent(n *bstNode, c ds.Comparable) (parent *bstNode, target *bstNode, isLeftChild bool) {
	for n != nil && n.val.CompareTo(c) != 0 {
		parent = n
		if n.val.CompareTo(c) > 0 {
			n = n.left
			isLeftChild = true
		} else {
			n = n.right
			isLeftChild = false
		}
	}
	if n == nil {
		return nil, nil, false
	}
	return parent, n, isLeftChild
}

// delete the left most child and return it
// if the given node is the left most one, only return it
func delLeftMost(n *bstNode) *bstNode {
	if n == nil {
		return nil
	}
	if n.left == nil {
		return n
	}
	for n.left.left != nil {
		n = n.left
	}
	tmp := n.left
	n.left = nil
	return tmp
}

// Validate the tree rooted at n by in-order traversal
func (n *bstNode) validate(enableLog bool) bool {
	stack := make([]*bstNode, 0)
	stack = pushAllLeft(stack, n)
	var last ds.Comparable
	i := 0
	for len(stack) > 0 {
		l := len(stack)
		cur := stack[l-1]
		stack = stack[:l-1]

		if enableLog {
			log.Printf("the %dth element: %v ", i, cur.val)
		}

		if last != nil && cur.val.CompareTo(last) < 0 {
			if enableLog {
				log.Printf("validate failed at the %dth element", i)
			}
			return false
		}

		last = cur.val
		if cur.right != nil {
			stack = pushAllLeft(stack, cur.right)
		}
		i++
	}
	if enableLog {
		log.Printf("valid tree")
	}
	return true
}

func pushAllLeft(stack []*bstNode, n *bstNode) []*bstNode {
	for n != nil {
		stack = append(stack, n)
		n = n.left
	}
	return stack
}

func (n *bstNode) value() ds.Comparable {
	if n == nil {
		return nil
	}
	return n.val
}
