package bar

import "fmt"

var X int

func init() {
	X++
	fmt.Println("bar", X)
}

func HelloWorld() string {
	return "bar"
}
