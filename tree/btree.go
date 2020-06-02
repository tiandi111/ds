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
	stack := make([]*btnode, 0, t.level)
	cur := t.root
	kidx := -1
	// find the key and save the path as a stack
	for cur != nil {
		stack = append(stack, cur)
		idx := cur.search(v)
		if idx < len(cur.keys) && cur.keys[idx].CompareTo(v) == 0 {
			kidx = idx
			break
		}
		if !cur.isLeaf() {
			cur = cur.nodes[idx]
		} else {
			break
		}
	}
	// key not found
	if kidx == -1 {
		return nil
	}
	if !cur.isLeaf() {
		if len(cur.nodes[kidx].keys) >= t.degree {
			next := cur.nodes[kidx]
			for !next.isLeaf() {
				stack = append(stack, next)
				next = next.last()
			}
			predecessor := next.max()
			cur.keys[kidx] = predecessor
			cur = next
		} else if len(cur.nodes[kidx+1].keys) >= t.degree {
			next := cur.nodes[kidx+1]
			for !next.isLeaf() {
				stack = append(stack, next)
				next = next.first()
			}
			succesor := next.min()
			cur.keys[kidx] = succesor
			cur = next
		} else {
			front := cur.nodes[kidx]
			back := cur.nodes[kidx+1]
			front.keys = append(front.keys, cur.keys[kidx])
			front.keys = append(front.keys, back.keys...)
			front.nodes = append(front.nodes, back.nodes...)
			cur.keys[kidx:] = cur.keys[kidx+1:]
			cur.keys = cur.keys[:len(cur.keys)-1]
			cur.nodes[kidx+1:] = cur.nodes[kidx+2:]
			cur.nodes = cur.nodes[:len(cur.nodes)-1]
			// todo: remove keys from front
		}
	}
	return nil
}

func replaceByPredecessor() {

}

func replaceBySuccessor() {

}

func merge() {

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
// return the right node index to go
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
