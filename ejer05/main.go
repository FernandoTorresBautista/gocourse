package main

import "fmt"

// ciclos
func main() {
	// for
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//
	// for j := 0; j < 10; j++ {
	// 	fmt.Println(j)
	// }

	//
	// numero := 0
	// for {
	// 	fmt.Println("continue")
	// 	fmt.Println("ingrese un numero")
	// 	fmt.Scanln(&numero)
	// 	if numero == 100 {
	// 		break
	// 	}
	// }

	// var i = 0
	// for i < 10 {
	// 	if i == 5 {
	// 		fmt.Printf("\n Multiplicar por 2 \n")
	// 		i = i * 2
	// 		continue
	// 	}
	// 	fmt.Printf("\n Valor de i: %d", i)
	// 	i++
	// }

	var i int = 0
RUTINA: // marca una etiqueta
	for i < 10 {
		if i == 4 {
			// i = i + 1
			i = i * 2
			// fmt.Println("voy a RUTINA sumando 2 a i")
			fmt.Println("voy a RUTINA multiplicando i por 2")
			goto RUTINA
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}
}
