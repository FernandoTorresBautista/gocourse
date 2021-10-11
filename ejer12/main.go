package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// Manejo de archivos
func main() {
	// leer archivos
	leoArchivo()
	leoArchivo2()
	// grabar archivos
	graboArchivo()
	graboArchivo2()
}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2() {
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Println("Linea > " + registro)
		}
	}
	archivo.Close() // cerrar el archivo
}

func graboArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, "Nueva linea")
	archivo.Close()
}

func graboArchivo2() {
	fileName := "./nuevoArchivo.txt"
	if !Append(fileName, "\nNueva linea 2") {
		fmt.Println("Error en la 2da linea")
	}
}

func Append(archivo, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error abriendo el archivo")
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error escribiendo en el archivo")
		return false
	}
	return true
}
