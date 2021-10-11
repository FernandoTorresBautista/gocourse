package main

import (
	"fmt"
	"time"
)

// Channels en GO - dialogo entre go rutines
func main() {
	// un canal es un espacio de memoria
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Llegue hasta aqui")

	// escuchar el canal1
	msg := <-canal1
	fmt.Println("Duration bucle: ", msg)
}

// canal y su tipo
func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 100000000; i++ {
	}
	final := time.Now()
	// asignar un valor a canal1
	canal1 <- final.Sub(inicio) // Duration entre inicio y fin
}
