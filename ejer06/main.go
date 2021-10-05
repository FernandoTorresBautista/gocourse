package main

import "fmt"

func main() {
	fmt.Println(uno(10))
	number, status := dos(1)
	fmt.Println(number, status)

	//
	fmt.Println("calculo: ", Calculo(1, 2, 3))
	fmt.Println("calculo: ", Calculo(10, 20, 30, 40))
	fmt.Println("calculo: ", Calculo(30, 1))
}

func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	}
	return 10, true
}

// parametros variables
func Calculo(numero ...int) int {
	total := 0
	// _ es el index
	for _, v := range numero {
		total += v
	}
	return total
}
