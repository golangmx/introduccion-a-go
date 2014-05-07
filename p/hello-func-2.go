package main

import "fmt"

func main() {
	// La instrucción de asignación ":=" infiere el tipo de variable del // HL
	// contexto. El ejemplo siguiente declara a <a> como una nueva variable // HL
	// del mismo tipo que "1" (int). Es equivalente a: // HL
	// // HL
	// var a int // HL
	// a = 1 // HL
	a := 1
	b := 2

	fn := func(a, b int) (int, int) { // <a> y <b> son ambos int, podemos omitir el tipo de <a>.
		return b, a
	}

	a, b = fn(a, b)

	// Imprime: a = 2, b = 1
	fmt.Printf("a = %d, b = %d\n", a, b)
}
