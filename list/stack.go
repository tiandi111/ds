package list

import (
	"fmt"
	"github.com/tiandi111/ds"
)

type GenericStack struct {
	arr []interface{}
}

func NewGenericStack(cap int) *GenericStack {
	return &GenericStack{make([]interface{}, 0, cap)}
}

func (s *GenericStack) Push(v interface{}) {
	s.arr = append(s.arr, v)
}

func (s *GenericStack) Pop() interface{} {
	if s.IsEmpty() {
		panic(fmt.Errorf("stack is empty"))
	}
	obj := s.arr[s.Size()-1]
	s.arr = s.arr[:s.Size()-1]
	return obj
}

func (s *GenericStack) Top() interface{} {
	if s.IsEmpty() {
		panic(fmt.Errorf("stack is empty"))
	}
	return s.arr[s.Size()-1]
}

func (s *GenericStack) Size() int {
	return len(s.arr)
}

func (s *GenericStack) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *GenericStack) NewIterator() ds.Iterator {
	return nil
}
