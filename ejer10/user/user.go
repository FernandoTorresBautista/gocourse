package user

import "time"

// scope: visible desde fuera
type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

// metodos (funciones)

func (user *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool) {
	user.Id = id
	user.Nombre = nombre
	user.FechaAlta = fechaalta
	user.Status = status
}
