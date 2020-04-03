package util

import "fmt"

func CheckIndexInt64(index, lo, hi int64) {
	if index < lo || index >= hi {
		panic(fmt.Errorf("index out of bound: %d", index))
	}
}
