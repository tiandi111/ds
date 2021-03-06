package hash

// todo: compressed bitmap;
// todo: do operations atomically
type BitMap interface {
	Add(bool) int64
	Set(offset int64) bool
	Clear(offset int64) bool
	Get(offset int64) bool
	Count(start, end int64) int64
	Size() int64
	NewIterator()
}

type BloomFilter interface {
	Set(string)
	Contains(string) bool
}
