package main

import (
	"fmt"
)

type hotdog int

var h hotdog = 10

func main() {

	x := 10
	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v", h, h)

	// h tem o mesmo valor de x, mas tem o tipo diferente
	// h = x
}
