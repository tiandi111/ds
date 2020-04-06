package hash

import (
	"github.com/tiandi111/ds/test"
	"k8s.io/apimachinery/pkg/util/rand"
	"strconv"
	"testing"
)

func TestNewBasicBloomFilter(t *testing.T) {
	test.AssertNonNil(t, NewBasicBloomFilter(10, 1, sha256Hashing))
}

func TestBasicBloomFilter_Set(t *testing.T) {
	bf := NewBasicBloomFilter(10, 2, sha256Hashing)
	bf.Set("bloomfilter")
	test.AssertTrue(t, bf.Contains("bloomfilter"))
}

func TestBasicBloomFilter_Set2(t *testing.T) {
	bf := NewBasicBloomFilter(100, 5, sha256Hashing)
	for i := 0; i < 1000; i++ {
		val := strconv.Itoa(i)
		bf.Set(val)
		test.AssertTrue(t, bf.Contains(val))
	}
}

func TestBasicBloomFilter_Contains(t *testing.T) {
	bf := NewBasicBloomFilter(10, 2, sha256Hashing)
	test.AssertFalse(t, bf.Contains("This key doesn't exist"))
}

func TestBasicBloomFilter_Contains2(t *testing.T) {
	N := 100
	m := make(map[int]bool)
	bf := NewBasicBloomFilter(285, 2, sha256Hashing)
	for i := 0; i < N; i++ {
		val := strconv.Itoa(i)
		if rand.Intn(2) == 0 {
			bf.Set(val)
			m[i] = true
		}
	}
	for i := 0; i < N; i++ {
		if !bf.Contains(strconv.Itoa(i)) {
			test.AssertFalse(t, m[i])
		}
	}
}
