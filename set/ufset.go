package set

import (
	"fmt"
	"github.com/tiandi111/ds/list"
)

const (
	ModeQuickFind  = 0
	ModeQuickUnion = 1
)

type BasicUnionFindSet struct {
	ufset
	// N is the number of elements
	// mode0: favor find, O(1) for find, O(logN) for Union
	// mode1: favor union, O(1) for union, O(logN) for Find
	mode int
}

func NewBasicUnionFindSet(size, mode int) *BasicUnionFindSet {
	var set ufset
	switch mode {
	case ModeQuickFind:
		set = newQuickFindSet(size)
	case ModeQuickUnion:
		set = newQuickUnionSet(size)
	default:
		panic(fmt.Errorf("invalid mode"))
	}
	return &BasicUnionFindSet{set, mode}
}

func (s *BasicUnionFindSet) Find(n int) int {
	return s.find(n)
}

func (s *BasicUnionFindSet) Union(a, b int) int {
	return s.union(a, b)
}

func (s *BasicUnionFindSet) Connected(a, b int) bool {
	return s.find(a) == s.find(b)
}

func (s *BasicUnionFindSet) Count(n int) int {
	return s.count()
}

type ufset interface {
	find(int) int
	union(int, int) int
	count() int
}

// quickFindSet
// O(1) for find
// O(N) for union
type quickFindSet struct {
	set []int
	cnt int
}

func newQuickFindSet(size int) *quickFindSet {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i
	}
	return &quickFindSet{arr, size}
}

func (s *quickFindSet) find(a int) int {
	s.checkIndex(a)
	return s.set[a]
}

func (s *quickFindSet) union(a, b int) int {
	s.checkIndex(a)
	s.checkIndex(b)
	ra := s.find(a)
	rb := s.find(b)
	if ra == rb {
		return ra
	}
	for i, root := range s.set {
		if root == ra {
			s.set[i] = rb
		}
	}
	return rb
}

func (s *quickFindSet) count() int {
	return s.cnt
}

func (s *quickFindSet) checkIndex(i int) {
	if i < 0 || i >= len(s.set) {
		panic(list.ErrIndexOutOfBound(i))
	}
}

// quickUnionSet
// H is tree height
// O(H) for find
// O(H) for union
type quickUnionSet struct {
	set []int
	cnt int
}

func newQuickUnionSet(size int) *quickUnionSet {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i
	}
	return &quickUnionSet{arr, size}
}

func (s *quickUnionSet) find(a int) int {
	s.checkIndex(a)
	r := a
	for r != s.set[r] {
		r = s.set[r]
	}
	return r
}

func (s *quickUnionSet) union(a, b int) int {
	s.checkIndex(a)
	s.checkIndex(b)
	ra := s.find(a)
	rb := s.find(b)
	if ra == rb {
		return ra
	}
	s.set[ra] = rb
	return rb
}

func (s *quickUnionSet) count() int {
	return s.cnt
}

func (s *quickUnionSet) checkIndex(i int) {
	if i < 0 || i >= len(s.set) {
		panic(list.ErrIndexOutOfBound(i))
	}
}
