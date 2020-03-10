package list

import "fmt"

var (
	ErrInvalidSkipListP            = fmt.Errorf("expected p in [0,1]")
	ErrInvaldSkipListHeadNodeLevel = fmt.Errorf("invalid level for skip list head node, must > 0")
)

func ErrIndexOutOfBound(n int) error {
	return fmt.Errorf("index out of bound: %d", n)
}

func errInvalidSkipListP() error {
	return ErrInvalidSkipListP
}

func errInvaldSkipListHeadNodeLevel() error {
	return ErrInvaldSkipListHeadNodeLevel
}
