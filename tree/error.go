package tree

import "fmt"

var (
	ErrInvalidRangeFindArgs = fmt.Errorf("from.CompareTo(to) <= 0 must hold")
)

func errInvalidRangeFindArgs() error {
	return ErrInvalidRangeFindArgs
}
