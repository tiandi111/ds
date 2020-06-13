package tree

import (
	"github.com/tiandi111/ds"
)

//type SortedTree interface {
//	Insert(ds.Comparable)
//	Remove(ds.Comparable) interface{}
//	Find(ds.Comparable) interface{}
//	RangeFind(from, to ds.Comparable) []interface{}
//	Size() int
//	NewIterator() ds.Iterator
//}

type GenericBTree struct {
	root   *btnode
	degree int
	size   int
	level  int
}

func NewGenericBTree(d int) *GenericBTree {
	return &GenericBTree{degree: d}
}

func (t *GenericBTree) Insert(v ds.Comparable) {
	t.size++
	if t.root == nil {
		t.root = newbtnode(v, t)
		t.level++
		return
	}
	stack := make([]*btnode, 0)
	cur := t.root
	for !cur.isLeaf() { // means cur node has at least one node
		stack = append(stack, cur)
		cur = cur.nodes[cur.search(v)]
	}
	cur.insertKeyNode(v, nil)
	spill(cur, stack)
}

func (t *GenericBTree) Remove(v ds.Comparable) interface{} {
	if t.root == nil {
		return nil
	}
	stack := make([]*nodeTrace, 0, t.level)
	cur := t.root
	kidx := -1
	// find the key and save the path in stack
	for cur != nil {
		idx := cur.search(v)
		if idx < len(cur.keys) && cur.keys[idx].CompareTo(v) == 0 {
			kidx = idx
			break
		}
		if !cur.isLeaf() {
			stack = append(stack, &nodeTrace{cur, idx})
			cur = cur.nodes[idx]
		} else {
			break
		}
	}
	// key not found
	if kidx == -1 {
		return nil
	}
	target := cur.keys[kidx]
	if !cur.isLeaf() {
		if len(cur.nodes[kidx].keys) >= t.degree {
			cur = replaceWithPredecessor(cur, kidx, &stack)
			kidx = len(cur.keys) - 1
		} else if len(cur.nodes[kidx+1].keys) >= t.degree {
			cur = replaceWithSuccessor(cur, kidx, &stack)
			kidx = 0
		} else {
			merge(cur, kidx)
			if len(cur.keys) < t.degree {
				underflow(cur, stack)
			}
			return t.Remove(v)
		}
	}
	deleteKeyFromLeaf(cur, kidx)
	underflow(cur, stack)
	t.size--
	return target
}

type nodeTrace struct {
	node *btnode
	nidx int
}

// replace the key of the node at index kidx with its predecessor
func replaceWithPredecessor(node *btnode, kidx int, stack *[]*nodeTrace) *btnode {
	next := node.nodes[kidx]
	for !next.isLeaf() {
		*stack = append(*stack, &nodeTrace{next, len(next.nodes) - 1})
		next = next.last()
	}
	node.keys[kidx] = next.max()
	return next
}

func replaceWithSuccessor(node *btnode, kidx int, stack *[]*nodeTrace) *btnode {
	next := node.nodes[kidx+1]
	for !next.isLeaf() {
		*stack = append(*stack, &nodeTrace{next, 0})
		next = next.first()
	}
	node.keys[kidx] = next.min()
	return next
}

func merge(node *btnode, kidx int) *btnode {
	front := node.nodes[kidx]
	back := node.nodes[kidx+1]
	// merge key and back node into front node
	front.keys = append(front.keys, node.keys[kidx])
	front.keys = append(front.keys, back.keys...)
	front.nodes = append(front.nodes, back.nodes...)
	// delete key and back node
	copy(node.keys[kidx:], node.keys[kidx+1:])
	node.keys = node.keys[:len(node.keys)-1]
	copy(node.nodes[kidx+1:], node.nodes[kidx+2:])
	node.nodes = node.nodes[:len(node.nodes)-1]
	return front
}

func deleteKeyFromLeaf(node *btnode, kidx int) {
	if !node.isLeaf() {
		panic("not a leaf node") // panic for testing
	}
	copy(node.keys[kidx:], node.keys[kidx+1:])
	node.keys = node.keys[:len(node.keys)-1]
}

func underflow(node *btnode, stack []*nodeTrace) {
	cur := node
	for len(stack) != 0 {
		if len(cur.keys) >= node.tree.degree-1 {
			return
		}

		parent := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		cidx := parent.nidx
		pnode := parent.node

		if cidx > 0 && len(pnode.nodes[cidx-1].keys) >= node.tree.degree {

			rotateToRight(cur, pnode, cidx)

		} else if cidx < len(pnode.nodes)-1 && len(pnode.nodes[cidx+1].keys) >= node.tree.degree {

			rotateToLeft(cur, pnode, cidx)

		} else {

			if cidx == len(parent.node.nodes)-1 {
				cidx--
			}
			merge(pnode, cidx)
		}
		cur = pnode
	}
	if len(stack) == 0 {
		if len(cur.keys) == 0 {
			cur.tree.level--
			if !cur.isLeaf() {
				cur.tree.root = cur.nodes[0]
			} else {
				cur.tree.root = nil
			}
		}
	}
}

func rotateToRight(cur, pnode *btnode, cidx int) {
	front := pnode.nodes[cidx-1]
	// rotate key
	cur.keys = append(cur.keys, pnode.keys[cidx-1])
	copy(cur.keys[1:], cur.keys)
	cur.keys[0] = pnode.keys[cidx-1]
	pnode.keys[cidx-1] = front.max()
	front.keys = front.keys[:len(front.keys)-1]
	// prepend node
	if !front.isLeaf() {
		cur.nodes = append(cur.nodes, new(btnode))
		copy(cur.nodes[1:], cur.nodes)
		cur.nodes[0] = front.last()
		front.nodes = front.nodes[:len(front.nodes)-1]
	}
}

func rotateToLeft(cur, pnode *btnode, cidx int) {
	back := pnode.nodes[cidx+1]
	// rotate key
	cur.keys = append(cur.keys, pnode.keys[cidx])
	pnode.keys[cidx] = back.keys[0]
	back.keys = back.keys[1:]
	// append node
	if !back.isLeaf() {
		cur.nodes = append(cur.nodes, back.nodes[0])
		back.nodes = back.nodes[1:]
	}
}

func (t *GenericBTree) Find(v ds.Comparable) interface{} {
	cur := t.root
	for cur != nil {
		idx := cur.search(v)
		if idx < len(cur.keys) && cur.keys[idx].CompareTo(v) == 0 {
			return cur.keys[idx]
		}
		if !cur.isLeaf() {
			cur = cur.nodes[idx]
		} else {
			break
		}
	}
	return nil
}

func spill(leaf *btnode, pstack []*btnode) {
	cur := leaf
	for cur.isFull() {
		silbling, key := cur.spilt()
		if len(pstack) != 0 {
			cur = pstack[len(pstack)-1]
			pstack = pstack[0 : len(pstack)-1]
			cur.insertKeyNode(key, silbling)
		} else {
			newRoot := newbtnode(key, cur.tree)
			newRoot.nodes = append(newRoot.nodes, cur)
			newRoot.nodes = append(newRoot.nodes, silbling)
			newRoot.tree.root = newRoot
			newRoot.tree.level++
			break
		}
	}
}

type btnode struct {
	keys  []ds.Comparable
	nodes []*btnode
	tree  *GenericBTree
}

// cap(keys) = 2*t - 1
// t-1  <= len(keys) <= 2*t - 1
func newbtnode(v ds.Comparable, tree *GenericBTree) *btnode {
	node := &btnode{
		keys:  make([]ds.Comparable, 0, 2*tree.degree), // capacity is 2*tree*degree so we pre-allocate an entry for spilled key
		nodes: make([]*btnode, 0, 2*tree.degree),
		tree:  tree,
	}
	node.keys = append(node.keys, v)
	return node
}

func (n *btnode) isLeaf() bool {
	return len(n.nodes) == 0
}

func (n *btnode) isFull() bool {
	return len(n.keys) == 2*n.tree.degree
}

func (n *btnode) min() ds.Comparable {
	return n.keys[0]
}

func (n *btnode) max() ds.Comparable {
	return n.keys[len(n.keys)-1]
}

func (n *btnode) first() *btnode {
	if n.isLeaf() {
		return nil
	}
	return n.nodes[0]
}

func (n *btnode) last() *btnode {
	if n.isLeaf() {
		return nil
	}
	return n.nodes[len(n.nodes)-1]
}

// invariant: len(nodes) == len(keys)+1
// return the correct node index to go
func (n *btnode) search(v ds.Comparable) int {
	for i, e := range n.keys {
		if v.CompareTo(e) <= 0 {
			return i
		}
	}
	return len(n.keys)
}

func (n *btnode) insertKeyNode(v ds.Comparable, node *btnode) {
	at := n.search(v)
	n.keys = append(n.keys, v)
	copy(n.keys[at+1:], n.keys[at:])
	n.keys[at] = v
	if node != nil {
		n.nodes = append(n.nodes, node)
		copy(n.nodes[at+2:], n.nodes[at+1:])
		n.nodes[at+1] = node
	}
}

func (n *btnode) spilt() (*btnode, ds.Comparable) {
	if !n.isFull() {
		return nil, nil
	}
	mid := len(n.keys) / 2
	upkey := n.keys[mid]
	sibling := &btnode{
		keys: n.keys[mid+1:],
		tree: n.tree,
	}
	n.keys = n.keys[:mid]
	if !n.isLeaf() { // then n must has children
		sibling.nodes = n.nodes[mid+1:]
		n.nodes = n.nodes[:mid+1]
	}
	return sibling, upkey
}
