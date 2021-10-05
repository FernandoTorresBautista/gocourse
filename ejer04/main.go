package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

// mostrar y aceptar datos
func main() {
	fmt.Println("Ingrese numero 1: ")
	// "%d" -> verbo
	// en windows toma el \n como parte del numero 1 y \r como parte del numero 2
	// fmt.Scanf("%d", &numero1)
	fmt.Scanln(&numero1) // se resuelve leyendo la linea
	fmt.Println("Ingrese numero 2: ")
	// fmt.Scanf("%d", &numero2)
	fmt.Scanln(&numero2)
	fmt.Printf("%d\n", numero1+numero2)
	fmt.Println(numero1, numero2)
	//
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()
		fmt.Println("leyenda: ", leyenda)
	}
	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)
}
