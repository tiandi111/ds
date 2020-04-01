package list

import (
	"fmt"
	"github.com/tiandi111/ds/test"
	"k8s.io/apimachinery/pkg/util/rand"
	"testing"
)

func TestNewGenericSkipList(t *testing.T) {
	test.AssertNonNil(t, NewGenericSkipList())
}

func TestGenericSkipList_Add(t *testing.T) {
	for i := 1; i < 100; i++ {
		sl := NewGenericSkipList()
		for j := 0; j < i; j++ {
			sl.Add(test.Cpb{rand.Intn(i)})
			test.AssertTrue(t, sl.head.validate(false))
			test.Assert(t, j+1, sl.Len())
		}
	}
}

func TestGenericSkipList_Get(t *testing.T) {
	test.AssertNil(t, NewGenericSkipList().Get(test.Cpb{1}))
}

func TestGenericSkipList_Get2(t *testing.T) {
	sl := NewGenericSkipList()
	size := 100
	for i := 0; i < size; i++ {
		c := test.Cpb{i}
		sl.Add(c)
		sl.head.validate(false)
	}
	for i := -1; i < size+1; i++ {
		got := sl.Get(test.Cpb{i})
		if i >= 0 && i < size {
			test.Assert(t, i, got.(test.Cpb).Val)
		} else {
			test.AssertNil(t, got)
		}
	}
}

func TestGenericSkipList_Del(t *testing.T) {
	for i := 1; i < 1000; i++ {
		sl := NewGenericSkipList()
		for j := 1; j <= i; j++ {
			sl.Add(test.Cpb{j})
		}
		expectedSize := i
		for j := 1; j <= i+1; j++ {
			del := test.Cpb{j}
			sl.Del(del)
			test.AssertNil(t, sl.Get(del))
			test.AssertTrue(t, sl.head.validate(false))
			if del.Val >= 1 && del.Val <= i {
				expectedSize--
			}
			test.Assert(t, expectedSize, sl.Len(), fmt.Sprintf("the %dth removal", j))
		}
	}
}

func TestGenericSkipList_Len(t *testing.T) {
	test.Assert(t, 0, NewGenericSkipList().Len())
}

func TestGenericSkipList_level(t *testing.T) {
	sl := NewGenericSkipList()
	cnt := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		l := sl.level()
		cnt[l]++
	}
}
