package main

import "fmt"

type SerVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool // ser vivo
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	estaVivo() bool // ser vivo
}

type vegetal interface {
	ClasificacionVegetal() string
}

// genero humano
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	// edad       int
	// altura     float32
	// peso       float32
	// respirando bool
	// pensando   bool
	// comiendo   bool
	hombre // heredando del hombre (en teoria)
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre {
		return "Hombre"
	}
	return "Mujer"
}
func (h *hombre) estaVivo() bool { return h.vivo }

// func (m *mujer) respirar()    { m.respirando = true }
// func (m *mujer) comer()       { m.comiendo = true }
// func (m *mujer) pensar()      { m.pensando = true }
// func (m *mujer) sexo() string {  return "Mujer"}

func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un(a) %s y ya estoy repirando \n", hu.sexo())
}

// genero animal
type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) EsCarnivoro() bool { return p.carnivoro }
func (p *perro) estaVivo() bool    { return p.vivo }

func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Printf("Soy un animal y estoy respirando")
}

func AnimalEsCarnivoro(an animal) int {
	if an.EsCarnivoro() {
		return 1
	}
	return 0
}

// serVivo
func estoyVivo(v SerVivo) bool {
	return v.estaVivo()
}

// interfaces
func main() {

	// humanos
	fmt.Println("Humanos")
	Pedro := new(hombre)
	Pedro.esHombre = true
	HumanosRespirando(Pedro)

	Maria := new(mujer)
	HumanosRespirando(Maria)

	// animales
	fmt.Println("Animales")
	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	totalCarnivoros = +AnimalEsCarnivoro(Dogo)
	fmt.Printf("Total carnivoros %d\n", totalCarnivoros)
	fmt.Printf("Estoy vivo = %t", estoyVivo(Dogo))
	//
}

//
