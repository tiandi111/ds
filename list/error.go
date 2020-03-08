package list

import "fmt"

func ErrIndexOutOfBound(n int) error {
	return fmt.Errorf("index out of bound: %d", n)
}
