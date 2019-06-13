package main

import "fmt"

func main() {

	nombre, apellido := "alexys", "lozada"
	//otra forma de usar := al re-asignar un valor de la misma, esta se debe usar  o declarar con otra variable nueva.↓
	nombre, edad := "pedroNel", 20
	fmt.Println(nombre, "", apellido, "", edad)
}
