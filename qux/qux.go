package qux

import (
	"fmt"
	"foo/bar"
)

func init() {
	bar.X++
	fmt.Println("qux", bar.X)
}

func Qux() int {
	return 1 + 1
}
