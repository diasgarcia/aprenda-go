package main

import "fmt"

// TIPOS PRIMITIVOS / BÁSICOS:
// bool, string
// int, int8, int16, int32, int64
// uint, uint8, uint16, uint32, uint64, uintptr
// byte, rune
// float32, float64
// complex64, complex128

// TIPOS COMPOSTOS:
// array, slice, map, struct, pointer, function, interface, channel

func main() {
	var ativo bool = true
	var nome string = "Rafael"
	var idade int = 21
	var altura float64 = 1.73

	fmt.Println(ativo, nome, idade, altura)
}
