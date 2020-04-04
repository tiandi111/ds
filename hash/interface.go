package hash

// todo: compressed bitmap
type BitMap interface {
	Add(bool) int64
	Set(offset int64) bool
	Clear(offset int64) bool
	Get(offset int64) bool
	Count(start, end int64) int64
	Size() int64
	NewIterator()
}
