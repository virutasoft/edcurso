package main

import "fmt"

func main() {
	//FUNCIONES VARIATICAS
	//recibe un numero variable de parametros del mismo tipo
	saludarVarios(18, "perra", "zorra", "lampedusa", "raccoon", "rata", "rastrorrea")
}

// func saludarVarios(nombres ...string) {
// 	//los tres puntos hacen a go inferir q recibe un slice
// 	//verificamos el formato de lo q retorna la funcion
// 	// fmt.Printf("%T\n", nombres)
// 	for _, v := range nombres { //recorremos el slice nombres con un for range
// 		fmt.Println("hola", v)
// 	}
// }

//las funciones variaticas tienen como restriccion q solo pueden recibir un parametro variatico al final de los demás

func saludarVarios(edad uint8, nombres ...string) {
	//los tres puntos hacen a go inferir q recibe un slice
	//verificamos el formato de lo q retorna la funcion
	// fmt.Printf("%T\n", nombres)
	for _, v := range nombres { //recorremos el slice nombres con un for range
		fmt.Println("hola", v, "tienes", edad, "años de edad")
	}
}
