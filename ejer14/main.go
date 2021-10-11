package main

import (
	"fmt"
	"log"
	"os"
)

// defer, panic & recover
func main() {

	// ejemplo defer
	archivo := "prueba.txt"
	f, err := os.Open(archivo)
	if err != nil {
		fmt.Println("error abriendo el archivo")
		os.Exit(1)
	}
	defer f.Close() // se ejecuta al salir de la funcion

	// ejemplo panic
	ejemploPanic()

}

func ejemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil { // ocurrio un panic
			log.Fatalf("Ocurrio un error que genero Panic\n %v", reco) // sale con exit status 1
		}
	}()

	a := 1
	if a == 1 {
		panic("Se encontro con el valor de 1")
	}
}
