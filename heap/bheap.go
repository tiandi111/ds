package heap

import (
	"fmt"
	"github.com/tiandi111/ds"
)

var (
	errLinkTreesWithDiffDegree = fmt.Errorf("link trees with different degree")
	errNilMinNode              = fmt.Errorf("minNode is nil, shouldn't happen")
)

type GenericBinomialHeap struct {
	head *bhNode
	size int
}

func NewGenericBinomialHeap() *GenericBinomialHeap {
	return new(GenericBinomialHeap)
}

func NewGenericBinomialHeapWithValue(c ds.Comparable) *GenericBinomialHeap {
	return &GenericBinomialHeap{newBhNode(c), 1}
}

func newGenericBinomialHeapWithHead(h *bhNode) *GenericBinomialHeap {
	return &GenericBinomialHeap{h, 1}
}

func (h *GenericBinomialHeap) Insert(c ds.Comparable) {
	h2 := NewGenericBinomialHeapWithValue(c)
	h.union(h2)
	h.size++
}

func (h *GenericBinomialHeap) Min() interface{} {
	var min ds.Comparable
	cur := h.head
	for cur != nil {
		if min == nil || cur.val.CompareTo(min) < 0 {
			min = cur.val
		}
		cur = cur.sibling
	}
	return min
}

func (h *GenericBinomialHeap) DelMin() interface{} {
	h.size--
	min := h.Min()
	if min == nil {
		return nil
	}
	var prev, minNode *bhNode
	cur := h.head
	for cur != nil {
		if cur.val.CompareTo(min.(ds.Comparable)) == 0 {
			minNode = cur
			if prev == nil {
				h.head = cur.sibling
			} else {
				prev.sibling = cur.sibling
			}
			cur.sibling = nil
			break
		}
		prev = cur
		cur = cur.sibling
	}
	if minNode == nil {
		panic(errNilMinNode)
	}
	leftMostChild := minNode.child
	// clean pointers to avoid memory leak
	minNode.parent = nil
	minNode.sibling = nil
	minNode.child = nil
	r := reverse(leftMostChild)
	h.union(newGenericBinomialHeapWithHead(r))
	return min
}

func (h *GenericBinomialHeap) Size() int {
	return h.size
}

func (h *GenericBinomialHeap) NewIterator() ds.Iterator {
	return nil
}

func (h *GenericBinomialHeap) Union(u *GenericBinomialHeap) *GenericBinomialHeap {
	h.size += u.size
	return h.union(u)
}

func (h *GenericBinomialHeap) union(u *GenericBinomialHeap) *GenericBinomialHeap {
	h.merge(u)
	if h.head == nil {
		return h
	}
	var prev *bhNode
	cur := h.head
	next := h.head.sibling
	for next != nil {
		switch {
		// only link cur and next
		// move forward:
		// case1: cur.degree != next.degree
		// case2: next.sibling != nil && cur.degree == next.sibling.degree
		case cur.degree != next.degree || (next.sibling != nil && cur.degree == next.sibling.degree):
			prev = cur
			cur = next
		// link:
		// case3 && case4: cur.degree == next.degree
		case cur.val.CompareTo(next.val) <= 0:
			cur.sibling = next.sibling
			cur.link(next)
		case cur.val.CompareTo(next.val) > 0:
			if prev == nil {
				h.head = next
			} else {
				prev.sibling = next
			}
			next.link(cur)
			cur = next
		}
		next = cur.sibling
	}
	return h
}

// merge m to h
func (h *GenericBinomialHeap) merge(m *GenericBinomialHeap) {
	h.head = mergeByDegree(h.head, m.head)
}

type bhNode struct {
	parent  *bhNode
	sibling *bhNode
	child   *bhNode // left most child
	degree  int     // the number of children
	val     ds.Comparable
}

func newBhNode(c ds.Comparable) *bhNode {
	return &bhNode{val: c}
}

// this operation links two B[k-1] trees to a B[k] tree
// link o to n
func (n *bhNode) link(o *bhNode) {
	if n.degree != o.degree {
		panic(errLinkTreesWithDiffDegree)
	}
	o.parent = n
	o.sibling = n.child
	n.child = o
	n.degree++
}

func mergeByDegree(n1, n2 *bhNode) *bhNode {
	dummy := new(bhNode)
	cur := dummy

	for n1 != nil || n2 != nil {
		switch {
		case n1 != nil && n2 != nil:
			if n1.degree <= n2.degree {
				cur.sibling = n1
				n1 = n1.sibling
			} else {
				cur.sibling = n2
				n2 = n2.sibling
			}
			cur = cur.sibling
		case n1 != nil:
			cur.sibling = n1
			return dummy.sibling
		case n2 != nil:
			cur.sibling = n2
			return dummy.sibling
		}
	}
	return dummy.sibling
}

func reverse(n *bhNode) *bhNode {
	if n == nil {
		return nil
	}
	var prev *bhNode
	cur := n
	next := n.sibling
	for next != nil {
		// clean parent to avoid memory leak
		cur.parent = nil
		tmp := next.sibling
		next.sibling = cur
		cur.sibling = prev
		prev = cur
		cur = next
		next = tmp
	}
	return cur
}

func (n *bhNode) length() int {
	len := 0
	cur := n
	for cur != nil {
		len++
		cur = cur.sibling
	}
	return len
}
