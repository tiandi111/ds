package tree

import "fmt"

var (
	errNilTreeNode = fmt.Errorf("operate no nil tree node")
)

func errNilPointer() error {
	return errNilTreeNode
}
