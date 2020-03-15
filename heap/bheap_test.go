package heap

import (
	"fmt"
	"github.com/tiandi111/ds/test"
	"k8s.io/apimachinery/pkg/util/rand"
	"testing"
	"time"
)

func TestNewGenericBinomialHeap(t *testing.T) {
	test.AssertNonNil(t, NewGenericBinomialHeap())
}

func TestGenericBinomialHeap_Insert(t *testing.T) {
	h := NewGenericBinomialHeap()
	for i := 0; i < 100; i++ {
		h.Insert(test.Cpb{i})
		test.AssertNil(t, h.validateBinomialHeap())
		test.Assert(t, i+1, h.Size())
	}
}

func TestGenericBinomialHeap_Min(t *testing.T) {
	test.AssertNil(t, NewGenericBinomialHeap().Min())
}

func TestGenericBinomialHeap_Min2(t *testing.T) {
	h := NewGenericBinomialHeap()
	for i := 100; i >= 0; i-- {
		h.Insert(test.Cpb{i})
		test.Assert(t, i, h.Min().(test.Cpb).Val)
		test.AssertNil(t, h.validateBinomialHeap())
	}
}

func TestGenericBinomialHeap_Min_RandomTest(t *testing.T) {
	rand.Seed(time.Now().Unix())
	h := NewGenericBinomialHeap()
	min := 200
	for i := 0; i < 100; i++ {
		val := rand.Intn(200) - 100
		h.Insert(test.Cpb{val})
		if val < min {
			min = val
		}
		test.Assert(t, min, h.Min().(test.Cpb).Val)
		test.AssertNil(t, h.validateBinomialHeap())
	}
}

func TestGenericBinomialHeap_DelMin(t *testing.T) {
	test.AssertNil(t, NewGenericBinomialHeap().DelMin())
}

func TestGenericBinomialHeap_DelMin2(t *testing.T) {
	h := NewGenericBinomialHeap()
	for i := 0; i < 100; i++ {
		h.Insert(test.Cpb{i})
	}
	for i := 0; i < 100; i++ {
		test.Assert(t, i, h.DelMin().(test.Cpb).Val)
		test.Assert(t, 100-1-i, h.Size())
		test.AssertNil(t, h.validateBinomialHeap())
	}
	test.AssertNil(t, h.DelMin())
}

func TestGenericBinomialHeap_union_BothNil(t *testing.T) {
	h1 := NewGenericBinomialHeap()
	h2 := NewGenericBinomialHeap()
	h1.union(h2)
	test.AssertNil(t, h1.head)
}

func TestGenericBinomialHeap_union_SingleNil(t *testing.T) {
	h1 := NewGenericBinomialHeapWithValue(test.Cpb{1})
	h2 := NewGenericBinomialHeap()
	test.Assert(t, h1, h1.union(h2))
	h3 := NewGenericBinomialHeap()
	h4 := NewGenericBinomialHeapWithValue(test.Cpb{2})
	test.Assert(t, h4.head, h3.union(h4).head)
}

func TestGenericBinomialHeap_union(t *testing.T) {
	h1 := NewGenericBinomialHeapWithValue(test.Cpb{1})
	h2 := NewGenericBinomialHeapWithValue(test.Cpb{2})
	h1.union(h2)
	test.Assert(t, 1, h1.head.length())
	test.AssertNil(t, h1.validateBinomialHeap())
}

func TestGenericBinomialHeap_union_RandomTest(t *testing.T) {
	for i := 0; i < 20; i++ {
		h1 := generateBinomialHeap(i)
		for j := 0; j < 20; j++ {
			h2 := generateBinomialHeap(i)
			h1.union(h2)
			test.AssertNil(t, h1.validateBinomialHeap())
		}
	}
}

func TestGenericBinomialHeap_Merge_RandomTest(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 1; i < 100; i++ {
		bh1 := NewGenericBinomialHeap()
		bh2 := NewGenericBinomialHeap()
		bh1.head = generateBhNodeListWithMonoIncrDegree(rand.Intn(i))
		bh2.head = generateBhNodeListWithMonoIncrDegree(rand.Intn(i))
		l1 := bh1.head.length()
		l2 := bh2.head.length()
		bh1.merge(bh2)
		test.AssertTrue(t, bh1.head.validateDegree())
		test.Assert(t, l1+l2, bh1.head.length())
	}
}

func TestMergeByDegree_BothNil(t *testing.T) {
	test.AssertNil(t, mergeByDegree(nil, nil))
}

func TestMergeByDegree_OneNil(t *testing.T) {
	h1 := new(bhNode)
	test.Assert(t, h1, mergeByDegree(h1, nil))
	h2 := new(bhNode)
	test.Assert(t, h2, mergeByDegree(nil, h2))
}

func TestMergeByDegree(t *testing.T) {
	n1 := &bhNode{degree: 0}
	n2 := &bhNode{degree: 1}
	merged := mergeByDegree(n1, n2)
	test.Assert(t, n1, merged)
	test.Assert(t, n2, merged.sibling)
}

func TestMergeByDegree_RandomTest(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 1; i < 100; i++ {
		n1 := generateBhNodeListWithMonoIncrDegree(rand.Intn(i))
		n2 := generateBhNodeListWithMonoIncrDegree(rand.Intn(i))
		l1 := n1.length()
		l2 := n2.length()
		merged := mergeByDegree(n1, n2)
		test.AssertTrue(t, merged.validateDegree())
		test.Assert(t, l1+l2, merged.length())
	}
}

func TestBhNode_Link_DiffDegree(t *testing.T) {
	defer func() {
		test.Assert(t, errLinkTreesWithDiffDegree, recover())
	}()
	n1 := &bhNode{degree: 0}
	n2 := &bhNode{degree: 1}
	n1.link(n2)
}

func TestBhNode_Link(t *testing.T) {
	c := new(bhNode)
	n1 := &bhNode{child: c, degree: 1}
	n2 := &bhNode{degree: 1}
	n1.link(n2)
	test.Assert(t, n2, n1.child)
	test.Assert(t, n1, n2.parent)
	test.Assert(t, c, n2.sibling)
	test.Assert(t, 2, n1.degree)
}

func TestReverse_Nil(t *testing.T) {
	test.AssertNil(t, reverse(nil))
}

func TestReverse(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := newBhNode(test.Cpb{0})
		cur := n
		for j := 1; j <= i; j++ {
			cur.sibling = newBhNode(test.Cpb{j})
			cur = cur.sibling
		}
		r := reverse(n)
		for j := i; j >= 0; j-- {
			test.Assert(t, j, r.val.(test.Cpb).Val)
			r = r.sibling
		}
		test.AssertNil(t, r)
	}
}

func generateBinomialTree(k int) *bhNode {
	if k == 0 {
		return newBhNode(test.Cpb{rand.Intn(1000)})
	}
	t1 := generateBinomialTree(k - 1)
	t2 := generateBinomialTree(k - 1)
	if t1.val.CompareTo(t2.val) < 0 {
		t1.link(t2)
		return t1
	}
	t2.link(t1)
	return t2
}

func (n *bhNode) validateBinomialTree() bool {
	child := n.child
	cnt := 0
	for child != nil {
		cnt++
		if n.val.CompareTo(child.val) > 0 || !child.validateBinomialTree() {
			return false
		}
		child = child.sibling
	}
	return n.degree == cnt
}

func generateBinomialHeap(maxDegree int) *GenericBinomialHeap {
	h := NewGenericBinomialHeap()
	cur := h.head
	for i := 0; i < maxDegree; i++ {
		if rand.Intn(2) == 1 {
			if h.head == nil {
				h.head = generateBinomialTree(i)
				cur = h.head
			} else {
				cur.sibling = generateBinomialTree(i)
				cur = cur.sibling
			}
		}
	}
	return h
}

func (h *GenericBinomialHeap) validateBinomialHeap() error {
	cur := h.head
	for cur != nil {
		if !cur.validateBinomialTree() {
			return fmt.Errorf("invalid binomial tree")
		}
		cur = cur.sibling
	}
	return h.head.validateDegreeForBHeap()
}

// two test functions for merge operation

func generateBhNodeListWithMonoIncrDegree(size int) *bhNode {
	head := new(bhNode)
	cur := head
	last := 0
	for i := 0; i < size-1; i++ {
		next := &bhNode{degree: last + rand.Intn(2)}
		cur.sibling = next
		cur = cur.sibling
		last = cur.degree
	}
	return head
}

func (n *bhNode) validateDegree() bool {
	last := -1
	cur := n
	for cur != nil {
		if cur.degree < last {
			return false
		}
		last = cur.degree
		cur = cur.sibling
	}
	return true
}

func (n *bhNode) validateDegreeForBHeap() error {
	last := -1
	cur := n
	for cur != nil {
		if cur.degree == last {
			return fmt.Errorf("degree equals to the last one")
		}
		if cur.degree < last {
			return fmt.Errorf("degree greater than the last one")
		}
		last = cur.degree
		cur = cur.sibling
	}
	return nil
}
