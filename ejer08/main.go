package main

import "fmt"

var tabla [10]int

// arreglos estaticos & slices
func main() {
	tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)

	tabla2 := [10]int{5, 6, 7, 8, 9, 10, 25, 36, 20, 10}
	// for i := 0; i < len(tabla); i++ {
	// 	fmt.Println(tabla2[i])
	// }
	fmt.Println(tabla2)

	// matriz
	var matriz [5][7]int
	matriz[3][5] = 1
	fmt.Println(matriz)

	// slice
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	variante2()

	variante3()

	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:]
	fmt.Println(porcion)
}

func variante3() {
	// slice de 5 elementos iniciales, con espacio reservado para 20
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d\n", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i) // append tiene su costo en proceso
	}
	fmt.Println(nums)
	// capacidad 128 -> reserva 2^n cada vez que sobrepasa el tamaño actual
	// 1,2,4,8,16,32,62,128*
	fmt.Printf("Largo %d, Capacidad %d\n", len(nums), cap(nums))
}
