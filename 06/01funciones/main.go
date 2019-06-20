package main

import "fmt"

func main() {
	saludar("perra", 35)
}

// func saludar() {
// 	fmt.Println("hola")
// }

//funciones con parametro
func saludar(nombre string, edad uint8) {
	fmt.Println("hola", nombre, "tienes", edad, "años")
	fmt.Printf("hola %s, tienes %d años de edad\n", nombre, edad)
	if edad > 25 {
		return
	}
	fmt.Println("eres joven")
	//funciones con return

}
