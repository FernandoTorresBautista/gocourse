package main

import "fmt"

// mapas
func main() {
	paises := make(map[string]string)
	fmt.Println(paises)
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises, "\n", paises["Mexico"])

	//
	campeonato := map[string]int{
		"Barcelona":   10,
		"Real Madrid": 8,
		"Cruz Azul":   5,
	}
	campeonato["River Plate"] = 5
	campeonato["Chivas"] = 7
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)
	for team, points := range campeonato {
		fmt.Printf("El equipo %s tiene un puntaje de %d \n", team, points)
	}

	//
	points, exist := campeonato["Santos"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", points, exist)
	points2, exist2 := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", points2, exist2)
}
