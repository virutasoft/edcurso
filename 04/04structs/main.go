package main

import "fmt"

type Persona struct {
	Nombre string   //almacena valor string
	Edad   uint8    //almacena enteros
	Email  []string //un slice con los emails de la persona
}

func main() {
	//structs: tipos de datos
	// var persona1 Persona
	// persona1.Nombre = "eddie"
	// persona1.Edad = 20

	// fmt.Println(persona1.Nombre) //fmt imprime los valores de la estructura en elmismo orden en el que esta declarado

	//shortcut

	// persona2 := Persona{}
	// persona2.Nombre = "lucrecia"
	// persona2.Edad = 64
	// fmt.Println(persona2)

	//otra forma de declarar structs
	//se puede agregar slices u otros tipo de estructuras al struct

	// otra forma de agregar un slice al struct es declarando antes el mismo

	emails := []string{"a@b.com", "c@d.com", "e@f.com"}
	persona2 := Persona{
		Nombre: "rogletek",
		Edad:   128,
		Email:  []string{"a@perra.com", "c@zorra.com", "e@mapache.com"},
	}
	fmt.Println(persona2)

	persona3 := Persona{
		"sparky",
		18,
		[]string{"g@h.com"},
	}
	fmt.Println(persona3)

	persona4 := Persona{
		Nombre: "rogletek",
		Edad:   128,
		Email:  emails,
	}
	fmt.Println(persona4)
}
