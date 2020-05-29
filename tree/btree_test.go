package tree

import (
	"errors"
	"fmt"
	"github.com/tiandi111/ds"
	"github.com/tiandi111/ds/test"
	"testing"
)

func TestGenericBTree_Insert(t *testing.T) {
	btree := NewGenericBTree(3)
	for i := 0; i < 100; i++ {
		btree.Insert(test.Cpb{i})
		test.AssertNil(t, validateBTree(btree))
	}
}

func validateBTree(t *GenericBTree) error {
	if t.root == nil {
		return nil
	}
	level := 0
	q := []*btnode{t.root}
	for len(q) != 0 {
		cur := q[0]
		if err := validateBtnode(cur, level == t.level-1, t.root == cur); err != nil {
			return err
		}
		for _, node := range cur.nodes {
			q = append(q, node)
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
			last = key
		} else {
			if key.CompareTo(last) < 0 {
				return fmt.Errorf("unexpected key index")
			}
			if !isleaf && node.nodes[i].min().CompareTo(last) < 0 ||
				node.nodes[i].max().CompareTo(key) >= 0 {
				return fmt.Errorf("unexpected node index")
			}
		}
	}
	if !isleaf && node.last().min().CompareTo(last) >= 0 {
		return fmt.Errorf("unexpected node index")
	}
	return nil
}
