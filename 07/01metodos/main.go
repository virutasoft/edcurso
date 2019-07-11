package main

import "fmt"

// Persona... es una structura
type Persona struct {
	nombre string
	edad   int8
}

type Numero int

func (n Numero) presentarse() {
	fmt.Println("Hola, yo soy un número")
}
func (p Persona) saludar() {
	fmt.Println("Hola, soy una persona")
}
func (p *Persona) asignarEdad(i int8) {
	if i >= 0 {
		p.edad = i

	} else {
		fmt.Println("NO ES VALIDA LA EDAD NEGATIVA")
	}
}

func main() {
	var persona Persona
	// var numero Numero
	// numero.presentarse()
	persona.saludar()
	persona.asignarEdad(-35)
	fmt.Println(persona)
	//si no se le llama a persona medinte un puntero, la funcion va a mostrar el valor que tiene dentro de la funcion pero como copia, por eso se le apunta con un puntero(*)
}
