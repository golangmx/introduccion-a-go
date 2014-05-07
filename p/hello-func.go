// Éste es nuestro hello-world.go

// Nombre del paquete. // HL
package main

// Instrucción para usar la librería "fmt". // HL
import "fmt"

func foo() {
	// Función "foo" sin argumentos y sin tipo de retorno. // HL
}

func multiplica(número int, veces int) int {
	// Función "multiplica": con argumentos tipo <int> y valor de retorno tipo <int>. // HL
	return número * veces
}

func intercambia(a int, b int) (int, int) {
	// Función "intercambia": acepta dos <int> y los regresa, cambiando el órden. // HL
	return b, a
}

// La función main() contiene instrucciones para ejecutar en el momento // HL
// que el programa comienza. // HL
func main() {

	var a int
	var b int

	a = 1
	b = 2

	a, b = intercambia(a, b) // Múltiple retorno y asignación. // HL

	fmt.Printf("a = %d, b = %d.\n", a, b)
}
