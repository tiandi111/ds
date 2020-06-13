package tree

import (
	"errors"
	"fmt"
	"testing"

	"github.com/tiandi111/ds/test"

	"github.com/tiandi111/ds"
)

func TestGenericBTree(t *testing.T) {
	for d := 2; d < 10; d++ {
		size := 2000
		btree := NewGenericBTree(d)
		// test insert
		for i := 0; i < size; i++ {
			btree.Insert(test.Cpb{i})
			test.AssertNil(t, validateBTree(btree))
		}
		// test find
		for i := 0; i < size; i++ {
			test.AssertNonNil(t, btree.Find(test.Cpb{i}), fmt.Sprintf("target: %d", i))
		}
		test.AssertNil(t, btree.Find(test.Cpb{-1}))
		test.AssertNil(t, btree.Find(test.Cpb{size + 1}))
		// print info
		fmt.Printf("level:%d size:%d\n", btree.level, btree.size)
		//printBtree(btree)
		// test remove
		toRemove := make(map[int]struct{})
		for i := 0; i < size; i++ {
			toRemove[i] = struct{}{}
		}
		for i, _ := range toRemove {
			//printBtree(btree)
			test.Assert(t, i, btree.Remove(test.Cpb{i}).(test.Cpb).Val)
			//printBtree(btree)
			test.AssertNil(t, validateBTree(btree))
			test.AssertNil(t, btree.Find(test.Cpb{i}))
			delete(toRemove, i)
			//fmt.Println(i)
		}
		test.AssertNil(t, btree.Remove(test.Cpb{-1}))
		test.AssertNil(t, btree.Remove(test.Cpb{size + 1}))
		test.Assert(t, 0, btree.size)
	}
}

func validateBTree(t *GenericBTree) error {
	if t.root == nil {
		return nil
	}
	level := 0
	q := []*btnode{t.root}
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			if err := validateBtnode(cur, level == t.level-1, t.root == cur); err != nil {
				return err
			}
			for _, node := range cur.nodes {
				q = append(q, node)
			}
		}
		level++
	}
	return nil
}

// key[i+1] >= key[i]
// key[i-1] <= nodes[i].min < key[i]
func validateBtnode(node *btnode, isleaf, isroot bool) error {
	if !isleaf && !isroot && len(node.keys) < node.tree.degree-1 {
		return fmt.Errorf("key number %d less than degree-1: %d", len(node.keys), node.tree.degree-1)
	}
	if isleaf && len(node.nodes) != 0 {
		return errors.New("leaf node has child")
	}
	if !isleaf && len(node.nodes) != len(node.keys)+1 {
		return fmt.Errorf("node number %d != key number+1 %d", len(node.nodes), len(node.keys)+1)
	}
	var last ds.Comparable
	for i, key := range node.keys {
		if last == nil {
			if !isleaf && node.nodes[i].max().CompareTo(key) >= 0 {
				return fmt.Errorf("unexpected node index")
			}
		} else {
			if key.CompareTo(last) < 0 {
				return fmt.Errorf("unexpected key index")
			}
			if !isleaf && (node.nodes[i].min().CompareTo(last) < 0 ||
				node.nodes[i].max().CompareTo(key) >= 0) {
				return fmt.Errorf("unexpected node index")
			}
		}
		last = key
	}
	if !isleaf && node.last().min().CompareTo(last) < 0 {
		return fmt.Errorf("unexpected node index")
	}
	return nil
}

func printBtree(tree *GenericBTree) {
	if tree.root == nil {
		return
	}
	q := make([]*btnode, 0)
	q = append(q, tree.root)
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			fmt.Printf("%v/", cur.keys)
			for _, child := range cur.nodes {
				q = append(q, child)
			}
		}
		fmt.Println()
	}
}
