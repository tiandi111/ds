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
	for i := 1; i < 5000; i++ {
		sl := NewGenericSkipList()
		for j := 0; j < i; j++ {
			sl.Add(test.Cpb{rand.Intn(i)})
			test.AssertTrue(t, sl.head.validate(false))
		}
	}
}

func TestGenericSkipList_level(t *testing.T) {
	sl := NewGenericSkipList()
	cnt := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		l := sl.level()
		cnt[l]++
	}
	fmt.Println(cnt)
}
