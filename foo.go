package foo

import (
	"fmt"
	"foo/bar"
	"foo/baz"
)

func init() {
	bar.X++
	fmt.Println("foo", bar.X)
}

func HelloWorld() string {
	return bar.HelloWorld() + baz.HelloWorld()
}
