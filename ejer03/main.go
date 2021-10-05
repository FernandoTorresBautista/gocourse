package main

import "fmt"

var estado bool

// condiciones
func main() {
	estado = true
	if estado = false; estado { // asignar en la misma linea del if, estado == true
		fmt.Println(estado)
	} else {
		fmt.Println("Es falso")
	}
	//
	switch numero := 10; numero {
	case 1:
		fmt.Println("1", numero)
	case 2:
		fmt.Println("2", numero)
	case 3:
		fmt.Println("3", numero)
	case 4:
		fmt.Println("4", numero)
	case 5:
		fmt.Println("5", numero)
	default:
		fmt.Println("Mayor a 5")
	}
}
