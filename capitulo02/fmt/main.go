package main

import "fmt"

func main() {

	// interpreted string literals
	x := "oi bom dia\ncomo vai?\tespero \"que\" tudo bem."
	fmt.Println(x)

	// raw string literals.
	y := `oi bom dia\ncomo vai?\tespero \"que\" tudo bem.`
	fmt.Println(y)

	// string print, retorna uma nova string
	a := "oi "
	b := "bom dia!"
	z := fmt.Sprint(a, b)

	fmt.Println(z)
}
