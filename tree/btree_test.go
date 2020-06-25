package tree

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/tiandi111/ds/test"

	"github.com/tiandi111/ds"
)

func TestGenericBTree(t *testing.T) {
	for d := 2; d < 3; d++ {
		size := 10
		btree := NewGenericBTree(d, test.SerializableCpbDeserializer)
		// test insert
		toInsert := make(map[int]struct{})
		for i := 0; i < size; i++ {
			toInsert[i] = struct{}{}
		}
		for i, _ := range toInsert {
			btree.Insert(&test.SerializableCpb{i})
			test.AssertNil(t, validateBTree(btree))
			delete(toInsert, i)
		}
		// test find
		for i := 0; i < size; i++ {
			test.AssertNonNil(t, btree.Find(&test.SerializableCpb{i}), fmt.Sprintf("target: %d", i))
		}
		test.AssertNil(t, btree.Find(&test.SerializableCpb{-1}))
		test.AssertNil(t, btree.Find(&test.SerializableCpb{size + 1}))
		// print info
		fmt.Printf("level:%d size:%d\n", btree.level, btree.size)
		printBtree(btree)
		// test remove
		toRemove := make(map[int]struct{})
		for i := 0; i < size; i++ {
			toRemove[i] = struct{}{}
		}
		for i, _ := range toRemove {
			//printBtree(btree)
			test.Assert(t, i, btree.Remove(&test.SerializableCpb{i}).(*test.SerializableCpb).Val)
			//printBtree(btree)
			test.AssertNil(t, validateBTree(btree))
			test.AssertNil(t, btree.Find(&test.SerializableCpb{i}))
			delete(toRemove, i)
			//fmt.Println(i)
		}
		test.AssertNil(t, btree.Remove(&test.SerializableCpb{-1}))
		test.AssertNil(t, btree.Remove(&test.SerializableCpb{size + 1}))
		test.Assert(t, 0, btree.size)
	}
}

func TestGenericBTree_Serialization(t *testing.T) {
	btree := NewGenericBTree(3, test.SerializableCpbDeserializer)
	for i := 0; i < 100; i++ {
		btree.Insert(&test.SerializableCpb{i})
	}
	var buf bytes.Buffer
	err := btree.Serialize(&buf)
	test.AssertNil(t, err)
	data, err := ioutil.ReadAll(&buf)
	test.AssertNil(t, err)
	fmt.Println(string(data))

	btree1 := NewGenericBTree(3, test.SerializableCpbDeserializer)
	err = btree1.Deserialize(bytes.NewBuffer(data))
	test.AssertNil(t, err)
	printBtree(btree1)
	test.AssertNil(t, validateBTree(btree1))
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
