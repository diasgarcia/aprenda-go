package main

import (
	"fmt"
)

type hotdog int

var h hotdog = 101

func main() {

	x := 10
	fmt.Printf("%v\n", x)

	x = int(h)
	fmt.Printf("%v\n", x)
}
