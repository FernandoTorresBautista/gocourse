package main

import (
	"fmt"
)

// Middlewares en GO
// Son interceptores que permiten ejecutar instrucciones comunes a
// varias funciones que reciben y devuelven los mismos tipos variables

var result int

func main() {
	fmt.Println("Inicio")
	result = operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(10, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de Operación")
		return f(a, b)
	}
}
