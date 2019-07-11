package main

//INTERFACES... SON UN CONT5RATO Q SE PUEDE IMPÑLETMENTAR EN CUALQUIER TIPO DE DATO EN GO
import "github.com/golang-es/edcurso-go/07/02interfaces/animales"

func main() {
	var p animales.Perro
	var g animales.Gato

	p.Nombre = "Rogleton"
	g.Nombre = "Willy"
	// AdoptarPerro(p)
	// AdoptarGato(g)
	AdoptarMascota(p)
	AdoptarMascota(g)
}

// func AdoptarPerro(p animales.Perro) {
// 	p.Comunicarse()
// }

// func AdoptarGato(g animales.Gato) {
// 	g.Comunicarse()
// }

func AdoptarMascota(m animales.Mascota) {
	m.Comunicarse()
}
