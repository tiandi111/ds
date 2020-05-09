package util

import (
	"fmt"
	"reflect"
)

func CheckIndexInt64(index, lo, hi int64) {
	if index < lo || index >= hi {
		panic(fmt.Errorf("index out of bound: %d", index))
	}
}

func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	return rv.Kind() != reflect.Struct && rv.IsNil()
}
