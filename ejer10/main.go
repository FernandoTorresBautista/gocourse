package main

import (
	"fmt"
	"time"

	us "GoCourse/ejer10/user" // importa el paquete (la carpeta user)
)

// type usuario struct {
// 	Id        int
// 	Nombre    string
// 	FechaAlta time.Time
// 	Status    bool
// }

type example struct {
	us.Usuario
}

// POO en GO - a traves de estructuras
func main() {

	// User := new(usuario)
	// User.Id = 10
	// User.Nombre = "Fernando"
	// User.Status = true
	// User.FechaAlta = time.Now()

	// User := new(us.Usuario)
	User := new(example)
	User.AltaUsuario(10, "Fernando", time.Now(), true)
	fmt.Println(User.Usuario)

}

//
