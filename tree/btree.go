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
	cur.insertKey(v)
	spill(cur, stack)
}

func spill(leaf *btnode, pstack []*btnode) {
	cur := leaf
	for cur.isFull() {
		silbling, key := cur.spilt()
		if len(pstack) != 0 {
			cur = pstack[len(pstack)-1]
			at := cur.search(key)
			copy(cur.keys[at+1:], cur.keys[at:])
			copy(cur.nodes[at+2:], cur.nodes[at+1:])
			cur.keys[at] = key
			cur.nodes[at+1] = silbling
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
	return len(n.nodes) == 2*n.tree.degree
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
func (n *btnode) search(v ds.Comparable) int {
	for i, e := range n.keys {
		if v.CompareTo(e) < 0 {
			return i
		}
	}
	return len(n.keys)
}

func (n *btnode) insertKey(v ds.Comparable) {
	at := n.search(v)
	copy(n.keys[at+1:], n.keys[at:])
}

func (n *btnode) insertNode(node *btnode) {

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
