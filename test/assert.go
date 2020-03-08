package test

import (
	"testing"
)

func Assert(t *testing.T, expected, got interface{}) {
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func AssertNonNil(t *testing.T, got interface{}) {
	if got == nil {
		t.Error("expected non-nil value, got nil")
	}
}

func AssertNil(t *testing.T, got interface{}) {
	if got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}
