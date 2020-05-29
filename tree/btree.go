package tree

import "github.com/tiandi111/ds"

//type SortedTree interface {
//	Insert(ds.Comparable)
//	Remove(ds.Comparable) interface{}
//	Find(ds.Comparable) interface{}
//	RangeFind(from, to ds.Comparable) []interface{}
//	Size() int
//	NewIterator() ds.Iterator
//}

type GenericBTree struct {
	root *btnode
	degree int
	size int
}

func NewGenericBTree(d int) *GenericBTree {
	return &GenericBTree{degree: d}
}

func (t *GenericBTree) Insert(v ds.Comparable) {
	if t.root == nil {
		t.root = newbtnode(v, t)
		return
	}
	stack := make([]*btnode, 0)
	cur := t.root
	for !cur.isLeaf() { // means cur node has at least one node
		stack = append(stack, cur)
		cur = cur.nodes[cur.search(v)]
	}
	if !cur.isFull() {
		cur.insertKey(v)
	} else {



	}
}

type btnode struct {
	keys []ds.Comparable
	nodes []*btnode
	tree *GenericBTree
}

// cap(keys) = 2*t - 1
// t-1  <= len(keys) <= 2*t - 1
func newbtnode(v ds.Comparable, tree *GenericBTree) *btnode {
	node := &btnode{
		keys:make([]ds.Comparable, 0, 2*tree.degree-1),
		nodes:make([]*btnode, 0, 2*tree.degree),
		tree:tree,
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

// invariant: len(nodes) == len(keys)+1
func(n *btnode) search(v ds.Comparable) int {
	for i, e := range n.keys {
		if v.CompareTo(e) < 0 {
			return i
		}
	}
	return len(n.keys)
}

// assume (1) n is nut full (2) n is a leaf node
func (n *btnode) insertKey(v ds.Comparable) {
	at := n.search(v)
	copy(n.keys[at+1:], n.keys[at:])
}