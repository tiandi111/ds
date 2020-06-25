package test

import (
	"bytes"
	"testing"
)

func TestSerializableCpb(t *testing.T) {
	var buf bytes.Buffer
	sc := &SerializableCpb{Val: 1}
	err := sc.Serialize(&buf)
	AssertNil(t, err)
	sc1 := new(SerializableCpb)
	err = sc1.Deserialize(&buf)
	AssertNil(t, err)
	Assert(t, sc1.Val, 1)
}
