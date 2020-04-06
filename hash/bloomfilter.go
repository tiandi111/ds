package hash

import (
	"crypto/sha256"
	"math"
)

type hashing func(data []byte, k int64) []uint64

type BasicBloomFilter struct {
	numOfMembers int64
	numOfHashing int64
	hashing
	bitmap *BitMap64
}

func NewBasicBloomFilter(m, k int64, h hashing) *BasicBloomFilter {
	return &BasicBloomFilter{0, k, h, NewBitMap64(m)}
}

func (b *BasicBloomFilter) Set(key string) {
	bytes := []byte(key)
	hashes := b.hashing(bytes, b.numOfHashing)

	for i := int64(0); i < b.numOfHashing; i++ {

		i := hashes[i] % uint64(b.bitmap.Size())

		b.bitmap.Set(int64(i))

	}

	b.numOfMembers++
}

func (b *BasicBloomFilter) Contains(key string) bool {
	bytes := []byte(key)
	hashes := b.hashing(bytes, b.numOfHashing)

	for i := int64(0); i < b.numOfHashing; i++ {

		i := hashes[i] % uint64(b.bitmap.Size())

		if !b.bitmap.Get(int64(i)) {
			return false
		}

	}
	return true
}

func (b *BasicBloomFilter) ExpectedErrorRate() float64 {
	pUnsetSingle := float64(1 - float64(1/b.bitmap.Size()))

	totalTries := float64(b.numOfMembers * b.numOfHashing)

	pUnset := math.Pow(pUnsetSingle, totalTries)

	return math.Pow(1-pUnset, float64(b.numOfHashing))
}

func (b *BasicBloomFilter) SetBitFraction() float64 {
	setBits := float64(b.bitmap.Count(0, b.bitmap.Size()))
	totalBits := float64(b.bitmap.Size())
	return setBits / totalBits
}

func sha256Hashing(data []byte, k int64) []uint64 {
	hash := sha256.Sum256(data)

	hash1 := readUint64(hash[:8])
	hash2 := readUint64(hash[8:16])
	hashes := make([]uint64, k)

	for i := int64(0); i < k; i++ {
		hashes[i] = hash1 + uint64(i)*hash2
	}

	return hashes
}

func readUint64(b []byte) uint64 {
	_ = b[7]
	return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
}
