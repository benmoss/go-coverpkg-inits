package baz

import (
	"fmt"
	"foo/bar"
)

func init() {
	bar.X++
	fmt.Println("baz", bar.X)
}
func HelloWorld() string {
	return "baz"
}
