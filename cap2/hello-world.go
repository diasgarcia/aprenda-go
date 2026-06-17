package main

import (
	"fmt"
)

func main() {
	x := 16
	y := "string"
	z := true

	fmt.Println(x, y, z)

	_, erros := fmt.Println("Hello World!", "Oi galera", 100)
	fmt.Println(erros)
}
