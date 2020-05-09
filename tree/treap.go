package tree

import (
	"github.com/tiandi111/ds"
	"math/rand"
)

const (
	MaxPriority = 1<<63 - 1
)

type GenericTreap struct {
	root *treapNode
	size int64
}

func NewGenericTreap() *GenericTreap {
	return new(GenericTreap)
}

func (t *GenericTreap) Insert(v ds.Comparable) {
	t.size++
	i := NewGenericTreapNode(v)
	n := t.root
	for n != nil {
		if i.val.CompareTo(n.val) <= 0 {
			if n.left == nil {
				n.left = i
				i.parent = n
				break
			}
			n = n.left
		} else {
			if n.right == nil {
				n.right = i
				i.parent = n
				break
			}
			n = n.right
		}
	}
	i.Up()
	if i.parent == nil {
		t.root = i
	}
}

func (t *GenericTreap) Remove(v ds.Comparable) interface{} {
	n := t.findNode(v)
	if n == nil {
		return nil
	}
	// remove root node, we need a new one
	if n == t.root {
		if n.IsLeafNode() {
			// no child
			t.root = nil
		} else if n.left == nil || n.right.priority < n.left.priority {
			// left rotate, set right child as new root
			t.root = t.root.right
		} else if n.right == nil || n.left.priority <= n.right.priority {
			// right rotate, set left child as new root
			t.root = t.root.left
		}
	}
	n.Down()
	n.SelfDelete()
	t.size--
	return n.val
}

func (t *GenericTreap) Find(v ds.Comparable) interface{} {
	node := t.findNode(v)
	if node == nil {
		return nil
	}
	return node.val
}

func (t *GenericTreap) findNode(v ds.Comparable) *treapNode {
	n := t.root
	for n != nil && v.CompareTo(n.val) != 0 {
		if v.CompareTo(n.val) < 0 {
			n = n.left
		} else {
			n = n.right
		}
	}
	if n == nil {
		return nil
	}
	return n
}

func (t *GenericTreap) RangeFind(from, to ds.Comparable) []interface{} {
	return nil
}

func (t *GenericTreap) Size() int64 {
	return t.size
}

func (t *GenericTreap) NewIterator(v ds.Comparable) ds.Iterator {
	return nil
}

// invariant: a.priority >= a.left.priority && a.priority >= a.right.priority
type treapNode struct {
	left     *treapNode
	right    *treapNode
	parent   *treapNode
	val      ds.Comparable
	priority int64
}

func NewGenericTreapNode(v ds.Comparable) *treapNode {
	return &treapNode{val: v, priority: rand.Int63n(MaxPriority)}
}

func (n *treapNode) Left() Node {
	// so we can compare the return Node interface with nil
	if n.left == nil {
		return nil
	}
	return n.left
}

func (n *treapNode) Right() Node {
	// so we can compare the return Node interface with nil
	if n.right == nil {
		return nil
	}
	return n.right
}

func (n *treapNode) Value() ds.Comparable {
	return n.val
}

func (n *treapNode) Up() {
	for n.parent != nil && n.priority < n.parent.priority {
		if n.IsLeftChildOf(n.parent) {
			n.parent.RightRotate()
		} else {
			n.parent.LeftRotate()
		}
	}
}

func (n *treapNode) Down() {
	for !n.IsLeafNode() {
		switch {
		case n.left == nil:
			n.LeftRotate()
		case n.right == nil:
			n.RightRotate()
		case n.left.priority <= n.right.priority:
			n.RightRotate()
		case n.left.priority > n.right.priority:
			n.LeftRotate()
		}
	}
}

func (n *treapNode) IsLeftChildOf(p *treapNode) bool {
	return p != nil && p.left == n
}

func (n *treapNode) IsRightChildOf(p *treapNode) bool {
	return p != nil && p.right == n
}

func (n *treapNode) IsLeafNode() bool {
	return n.left == nil && n.right == nil
}

func (n *treapNode) RightRotate() {
	if n.left == nil {
		return
	}
	newLeft := n.left.right
	if newLeft != nil {
		newLeft.parent = n
	}

	n.left.right = n
	n.left.parent = n.parent
	if n.IsLeftChildOf(n.parent) {
		n.parent.left = n.left
	}
	if n.IsRightChildOf(n.parent) {
		n.parent.right = n.left
	}

	newParent := n.left
	n.left = newLeft
	n.parent = newParent
}

func (n *treapNode) LeftRotate() {
	if n.right == nil {
		return
	}
	newRight := n.right.left
	if newRight != nil {
		newRight.parent = n
	}

	n.right.left = n
	n.right.parent = n.parent
	if n.IsLeftChildOf(n.parent) {
		n.parent.left = n.right
	}
	if n.IsRightChildOf(n.parent) {
		n.parent.right = n.right
	}

	newParent := n.right
	n.right = newRight
	n.parent = newParent
}

// delete itself and its subtrees from the tree
func (n *treapNode) SelfDelete() {
	if n.IsLeftChildOf(n.parent) {
		n.parent.left = nil
	} else if n.IsRightChildOf(n.parent) {
		n.parent.right = nil
	}
	n.parent = nil
}
