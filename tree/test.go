package tree

import (
	"github.com/tiandi111/ds"
	"github.com/tiandi111/ds/list"
	"reflect"
)

type ValidateFunc func(Node) bool

func InOrderValidate(root Node, validateFunc ...ValidateFunc) bool {
	if reflect.ValueOf(root).IsNil() {
		return true
	}

	stack := list.NewGenericStack(8)
	pushAllLeftNode(root, stack)

	var last ds.Comparable

	for !stack.IsEmpty() {
		cur := stack.Pop().(Node)

		if last != nil && cur.Value().CompareTo(last) < 0 {
			return false
		}
		for _, f := range validateFunc {
			if !f(cur) {
				return false
			}
		}

		last = cur.Value()

		pushAllLeftNode(cur.Right(), stack)
	}

	return true
}

func pushAllLeftNode(node Node, stack list.Stack) {
	for node != nil {
		stack.Push(node)
		node = node.Left()
	}
}
