package main

import (
	"fmt"
	"strings"
	"time"
)

// Go Rutines - ejecucion asincrona de promesas
func main() {
	// go - ejecutar de forma asincrona
	go miNombreLento("Fernando")
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x) // con esto vemos que si escribimos termina la ejecucion no espera la ejecucion asincrona
}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
