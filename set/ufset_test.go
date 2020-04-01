package set

import (
	"github.com/tiandi111/ds/test"
	"testing"
)

func TestNewBasicUnionFindSet(t *testing.T) {
	set := NewBasicUnionFindSet(10, ModeQuickFind)
	test.AssertNonNil(t, set)
	_, ok := set.ufset.(*quickFindSet)
	test.AssertTrue(t, ok)
}

func TestNewBasicUnionFindSet_1(t *testing.T) {
	set := NewBasicUnionFindSet(10, ModeQuickUnion)
	test.AssertNonNil(t, set)
	_, ok := set.ufset.(*quickUnionSet)
	test.AssertTrue(t, ok)
}

func TestNewQuickFindSet(t *testing.T) {
	set := newQuickFindSet(10)
	test.AssertNonNil(t, set)
	test.Assert(t, 10, len(set.set))
	test.Assert(t, 10, set.cnt)
}

func TestQuickFindSet_Find(t *testing.T) {
	set := newQuickFindSet(10)
	for i := 0; i < 10; i++ {
		test.Assert(t, i, set.find(i))
	}
}

func TestQuickFindSet_Union(t *testing.T) {
	//set := newQuickFindSet(10)

}
