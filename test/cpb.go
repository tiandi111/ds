package test

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/tiandi111/ds"
)

// Cpb implements ds.Comparable
// only used for testing
type Cpb struct {
	Val int
}

func (c Cpb) CompareTo(other ds.Comparable) int {
	return c.Val - other.(Cpb).Val
}

func (c Cpb) String() string {
	return fmt.Sprintf("%v", c.Val)
}

type SerializableCpb struct {
	Val int
}

func (sc *SerializableCpb) CompareTo(other ds.Comparable) int {
	return sc.Val - other.(*SerializableCpb).Val
}

func (sc *SerializableCpb) Serialize(w io.Writer) error {
	data := []byte(fmt.Sprintf("%v", sc))
	n, err := w.Write(data)
	if err != nil {
		return err
	}
	if n != len(data) {
		return errors.New("incomplete serialization")
	}
	return nil
}

func (sc *SerializableCpb) Deserialize(r io.Reader) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	val, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	sc.Val = val
	return nil
}

func (sc *SerializableCpb) String() string {
	return fmt.Sprintf("%v", sc.Val)
}

func SerializableCpbDeserializer(r io.Reader) (ds.Element, error) {
	sc := new(SerializableCpb)
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	val, err := strconv.Atoi(string(data))
	if err != nil {
		return nil, err
	}
	sc.Val = val
	return sc, nil
}
