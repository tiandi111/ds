package benchmark

import (
	"github.com/tiandi111/ds/list"
	"github.com/tiandi111/ds/test"
	"github.com/tiandi111/ds/tree"
	"k8s.io/apimachinery/pkg/util/rand"
)

const (
	K             = 1000
	M             = 1000 * K
	InsertionReps = M
	SearchSize    = M
	SearchReps    = 500 * K
)

func SkipListInsertion_HighP() {
	InitSkipList(InsertionReps, 2)
}

func SkipListInsertion_LowP() {
	InitSkipList(InsertionReps, 3)
}

func BstInsertion() {
	InitBst(InsertionReps)
}

// todo: skip list search is significantly slower than bst search, investigate on this
func SkipListSearch(sl *list.GenericSkipList, reps int) {
	for i := 0; i < reps; i++ {
		sl.Get(test.Cpb{rand.Intn(sl.Len())})
	}
}

func BstSearch(bst *tree.GenericBinarySearchTree, reps int) {
	for i := 0; i < reps; i++ {
		bst.Find(test.Cpb{rand.Intn(bst.Size())})
	}
}

func InitSkipList(size int, n int) *list.GenericSkipList {
	sl := list.NewGenericSkipListWithN(n)
	for i := 0; i < size; i++ {
		sl.Add(test.Cpb{rand.Intn(size)})
	}
	return sl
}

func InitBst(size int) *tree.GenericBinarySearchTree {
	bst := tree.NewGenericBinarySearchTree()
	for i := 0; i < size; i++ {
		bst.Insert(test.Cpb{rand.Intn(size)})
	}
	return bst
}
