package hash

import (
	"fmt"
)

const (
	offsetMask             = 63
	uint64Mask             = 0xffffffffffffffff
	errOffsetOutOfBoundFmt = "offset out of bound: %d"
	errInvalidRangeFmt     = "invalid range: [%d, %d)"
)

var (
	setMap       = [64]uint64{1, 1 << 1, 1 << 2, 1 << 3, 1 << 4, 1 << 5, 1 << 6, 1 << 7, 1 << 8, 1 << 9, 1 << 10, 1 << 11, 1 << 12, 1 << 13, 1 << 14, 1 << 15, 1 << 16, 1 << 17, 1 << 18, 1 << 19, 1 << 20, 1 << 21, 1 << 22, 1 << 23, 1 << 24, 1 << 25, 1 << 26, 1 << 27, 1 << 28, 1 << 29, 1 << 30, 1 << 31, 1 << 32, 1 << 33, 1 << 34, 1 << 35, 1 << 36, 1 << 37, 1 << 38, 1 << 39, 1 << 40, 1 << 41, 1 << 42, 1 << 43, 1 << 44, 1 << 45, 1 << 46, 1 << 47, 1 << 48, 1 << 49, 1 << 50, 1 << 51, 1 << 52, 1 << 53, 1 << 54, 1 << 55, 1 << 56, 1 << 57, 1 << 58, 1 << 59, 1 << 60, 1 << 61, 1 << 62, 1 << 63}
	clearMap     = [64]uint64{1 ^ uint64Mask, 1<<1 ^ uint64Mask, 1<<2 ^ uint64Mask, 1<<3 ^ uint64Mask, 1<<4 ^ uint64Mask, 1<<5 ^ uint64Mask, 1<<6 ^ uint64Mask, 1<<7 ^ uint64Mask, 1<<8 ^ uint64Mask, 1<<9 ^ uint64Mask, 1<<10 ^ uint64Mask, 1<<11 ^ uint64Mask, 1<<12 ^ uint64Mask, 1<<13 ^ uint64Mask, 1<<14 ^ uint64Mask, 1<<15 ^ uint64Mask, 1<<16 ^ uint64Mask, 1<<17 ^ uint64Mask, 1<<18 ^ uint64Mask, 1<<19 ^ uint64Mask, 1<<20 ^ uint64Mask, 1<<21 ^ uint64Mask, 1<<22 ^ uint64Mask, 1<<23 ^ uint64Mask, 1<<24 ^ uint64Mask, 1<<25 ^ uint64Mask, 1<<26 ^ uint64Mask, 1<<27 ^ uint64Mask, 1<<28 ^ uint64Mask, 1<<29 ^ uint64Mask, 1<<30 ^ uint64Mask, 1<<31 ^ uint64Mask, 1<<32 ^ uint64Mask, 1<<33 ^ uint64Mask, 1<<34 ^ uint64Mask, 1<<35 ^ uint64Mask, 1<<36 ^ uint64Mask, 1<<37 ^ uint64Mask, 1<<38 ^ uint64Mask, 1<<39 ^ uint64Mask, 1<<40 ^ uint64Mask, 1<<41 ^ uint64Mask, 1<<42 ^ uint64Mask, 1<<43 ^ uint64Mask, 1<<44 ^ uint64Mask, 1<<45 ^ uint64Mask, 1<<46 ^ uint64Mask, 1<<47 ^ uint64Mask, 1<<48 ^ uint64Mask, 1<<49 ^ uint64Mask, 1<<50 ^ uint64Mask, 1<<51 ^ uint64Mask, 1<<52 ^ uint64Mask, 1<<53 ^ uint64Mask, 1<<54 ^ uint64Mask, 1<<55 ^ uint64Mask, 1<<56 ^ uint64Mask, 1<<57 ^ uint64Mask, 1<<58 ^ uint64Mask, 1<<59 ^ uint64Mask, 1<<60 ^ uint64Mask, 1<<61 ^ uint64Mask, 1<<62 ^ uint64Mask, 1<<63 ^ uint64Mask}
	countOneMap  = [256]uint8{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8}
	mask8bitsMap = [8]uint64{0xff, 0xff00, 0xff0000, 0xff000000, 0xff00000000, 0xff0000000000, 0xff000000000000, 0xff00000000000000}
)

type BitMap64 struct {
	bmap     []uint64
	size     int64
	sizeLow  int64
	sizeHigh int64
}

func NewBitMap64(size int64) *BitMap64 {
	len := size / 64
	bmap := make([]uint64, len+1)
	return &BitMap64{bmap, size, size & offsetMask, size >> 6}
}

func (m *BitMap64) Add(v bool) int64 {
	if v {
		m.bmap[m.sizeHigh] |= setMap[m.sizeLow]
	}
	m.sizeLow = (m.sizeLow + 1) & offsetMask
	if m.sizeLow == 0 {
		m.bmap = append(m.bmap, 0)
		m.sizeHigh++
	}
	m.size++
	return m.size
}

func (m *BitMap64) Set(offset int64) bool {
	m.checkOffset(offset)
	hi := offset >> 6
	low := offset & offsetMask
	old := m.bmap[hi]&setMap[low] != 0
	m.bmap[hi] |= setMap[low]
	return old
}

func (m *BitMap64) Clear(offset int64) bool {
	m.checkOffset(offset)
	hi := offset >> 6
	low := offset & offsetMask
	old := m.bmap[hi]&setMap[low] != 0
	m.bmap[hi] &= clearMap[low]
	return old
}

func (m *BitMap64) Get(offset int64) bool {
	m.checkOffset(offset)
	return m.bmap[offset>>6]&setMap[offset&offsetMask] != 0
}

// left inclusive, right exclusive
func (m *BitMap64) Count(st, end int64) int64 {
	m.checkRange(st, end)
	end--
	var cnt int64
	sthi, stlo, endhi, endlo := st>>6, st&offsetMask, end>>6, end&offsetMask
	for (sthi < endhi || (sthi == endhi && stlo <= endlo)) && stlo != 64 {
		if m.bmap[sthi]&setMap[stlo] != 0 {
			cnt++
		}
		stlo = stlo + 1
		if stlo == 64 {
			sthi++
		}
	}
	for i := sthi; i < endhi; i++ {
		cnt += quickCount(m.bmap[i])
	}
	for i := int64(0); i < endlo; i++ {
		if m.bmap[endhi]&setMap[i] != 0 {
			cnt++
		}
	}
	return cnt
}

func count(n uint64) int64 {
	var cnt int64
	for ; n != 0; cnt++ {
		n &= n - 1
	}
	return cnt
}

func quickCount(n uint64) int64 {
	var cnt int64
	for i := uint(0); i < 8; i++ {
		num := (n & mask8bitsMap[i]) >> (i << 3)
		cnt += int64(countOneMap[num])
	}
	return cnt
}

func (m *BitMap64) Size() int64 {
	return m.size
}

func (m *BitMap64) checkRange(st, end int64) {
	if st >= end {
		panic(fmt.Errorf(errInvalidRangeFmt, st, end))
	}
	if st < 0 || st >= m.size {
		panic(fmt.Errorf(errInvalidRangeFmt, st, end))
	}
	if end < 0 || end > m.size {
		panic(fmt.Errorf(errInvalidRangeFmt, st, end))
	}
}

func (m *BitMap64) checkOffset(offset int64) {
	if offset < 0 || offset >= m.size {
		panic(fmt.Errorf(errOffsetOutOfBoundFmt, offset))
	}
}
