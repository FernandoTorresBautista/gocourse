package main

import "fmt"

// variable global
var numero int
var texto string
var status bool = true

// variables
func main() {
	//fmt.Println(numero, "-", texto, "-", status)
	fmt.Println(numero, "-", texto, "-")
	var numero2 int
	numero3 := 4
	fmt.Println(numero2, numero3)
	numero4, numero5, status := 50, "numero50", false
	fmt.Println(numero4, numero5, status)
	mostrarStatus()
}

func mostrarStatus() {
	fmt.Println(status)
}
