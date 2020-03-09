package test

import (
	"reflect"
	"testing"
)

// code in this package comes from etcd

func Assert(t *testing.T, e, g interface{}, msg ...string) {
	if (e == nil || g == nil) && (isNil(e) && isNil(g)) {
		return
	}
	if reflect.DeepEqual(e, g) {
		return
	}
	s := ""
	for _, m := range msg {
		s += m
	}
	t.Errorf("%s expected %v, got %v", s, e, g)
}

func AssertNonNil(t *testing.T, got interface{}) {
	if got == nil {
		t.Errorf("expected non-nil value, got %v", got)
	}
}

func AssertNil(t *testing.T, got interface{}, msg ...string) {
	Assert(t, nil, got, msg...)
}

func AssertTrue(t *testing.T, got bool, msg ...string) {
	Assert(t, true, got, msg...)
}

func AssertFalse(t *testing.T, got bool, msg ...string) {
	Assert(t, false, got, msg...)
}

func isNil(v interface{}) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	return rv.Kind() != reflect.Struct && rv.IsNil()
}
