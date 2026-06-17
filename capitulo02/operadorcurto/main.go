package main

import (
	"fmt"
)

var y2 = "olá\n"

func main() {

	fmt.Println(y2)

	x := 10
	y := "olá bom dia"
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x, z := 20, 30.
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n\n", z, z)
}
