package benchmark

import (
	"testing"
)

func BenchmarkSkipListInsertion_HighP(b *testing.B) {
	SkipListInsertion_HighP()
}

func BenchmarkSkipListInsertion_LowP(b *testing.B) {
	SkipListInsertion_LowP()
}

func BenchmarkBstInsertion(b *testing.B) {
	BstInsertion()
}

func BenchmarkSkipListSearch_HighP(b *testing.B) {
	sl := InitSkipList(SearchSize, 3)
	b.ResetTimer()
	SkipListSearch(sl, SearchReps)
}

func BenchmarkSkipListSearch_LowP(b *testing.B) {
	sl := InitSkipList(SearchSize, 4)
	b.ResetTimer()
	SkipListSearch(sl, SearchReps)
}

func BenchmarkBstSearch(b *testing.B) {
	bst := InitBst(SearchSize)
	b.ResetTimer()
	BstSearch(bst, SearchReps)
}
