package main

import "fmt"

// funciones anonimas y closures

var Calculo func(int, int) int = func(i1 int, i2 int) int {
	return i1 + i2
}

func main() {
	// funciones anonimas
	fmt.Println("funciones anonimas")
	fmt.Printf("Sumo 5 + 7 = %d \n", Calculo(5, 7))

	// Restamos (redefinimos la funcion)
	Calculo = func(n1, n2 int) int {
		return n1 - n2
	}
	fmt.Printf("Restamos 6 - 4 = %d \n", Calculo(6, 4))

	// Dividimos
	Calculo = func(n1, n2 int) int {
		return n1 / n2
	}
	fmt.Printf("Dividimos 12 / 3 = %d \n", Calculo(12, 3))

	Operaciones()

	fmt.Printf("\nClosures \n")
	// closures
	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}

}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

// Closures - devuelve una funcion
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
