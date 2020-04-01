package hash

type BitMap interface {
	Set(offset int64)
	Get(offset int64) bool
	Count(start, end int64) int64
	AND(offsets ...int64)
	OR()
	XOR()
	NOT()
	NewIterator()
}
