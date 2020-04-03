package hash

import (
	"fmt"
	"github.com/tiandi111/ds/test"
	"testing"
)

func TestNewBitMap64(t *testing.T) {
	bitmap := NewBitMap64(10)
	test.AssertNonNil(t, bitmap)
}

func TestNewBitMap642(t *testing.T) {
	for i := 0; i < 128; i++ {
		bitmap := NewBitMap64(int64(i))
		test.Assert(t, int64(i), bitmap.size)
		test.Assert(t, int64(i%64), bitmap.sizeLow)
		test.Assert(t, int64(i/64), bitmap.sizeHigh)
	}
}

func TestBitMap64_Add(t *testing.T) {
	//bitmap := NewBitMap64(0)

}

func TestBitMap64_Set_Overflow(t *testing.T) {
	defer func() {
		test.Assert(t, fmt.Errorf(errOffsetOutOfBoundFmt, 10), recover())
	}()
	bitmap := NewBitMap64(10)
	bitmap.Set(10)
}

func TestBitMap64_Set(t *testing.T) {
	N := int64(10000)
	bitmap := NewBitMap64(N)
	for i := int64(0); i < N; i++ {
		bitmap.Set(i)
		for j := int64(0); j <= i; j++ {
			test.AssertTrue(t, bitmap.Get(j))
		}
	}
}

func TestBitMap64_Clear(t *testing.T) {
	N := int64(10000)
	bitmap := NewBitMap64(N)
	for i := int64(0); i < N; i++ {
		bitmap.Set(i)
	}
	for i := int64(0); i < N; i++ {
		bitmap.Clear(i)
		for j := int64(0); j <= i; j++ {
			test.AssertFalse(t, bitmap.Get(j))
		}
	}
}
