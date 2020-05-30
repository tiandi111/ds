package tree

import (
	"errors"
	"fmt"
	"github.com/tiandi111/ds"
	"github.com/tiandi111/ds/test"
	"testing"
)

func TestGenericBTree_Insert(t *testing.T) {
	btree := NewGenericBTree(2)
	for i := 0; i < 10; i++ {
		btree.Insert(test.Cpb{i})
		test.AssertNil(t, validateBTree(btree), fmt.Sprintf("iter%d", i))
		test.Assert(t, i+1, btree.size)
	}
	printBtree(btree)
}

func validateBTree(t *GenericBTree) error {
	if t.root == nil {
		return nil
	}
	q := []*btnode{t.root}
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		if err := validateBtnode(cur, t.root == cur); err != nil {
			return err
		}
		for _, node := range cur.nodes {
			q = append(q, node)
		}
	}
	return nil
}

// key[i+1] >= key[i]
// key[i-1] <= nodes[i].min < key[i]
func validateBtnode(node *btnode, isroot bool) error {
	isleaf := node.isLeaf()
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
			last = key
		} else {
			if key.CompareTo(last) < 0 {
				return fmt.Errorf("unexpected key index")
			}
			if !isleaf && (node.nodes[i].min().CompareTo(last) < 0 ||
				node.nodes[i].max().CompareTo(key) >= 0) {
				return fmt.Errorf("unexpected node index")
			}
		}
	}
	if !isleaf && node.last().min().CompareTo(last) >= 0 {
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
			fmt.Print(cur.keys, "/")
			for _, node := range cur.nodes {
				q = append(q, node)
			}
		}
		fmt.Println()
	}
}
