package main

import "fmt"

func main() {
	//arrays
	//var nombres [3]string
	//caracteristicas de los arrays en go:
	//tdos los valores deben ser del mismo tipo
	//tamño fijo
	// nombres[0] = "perra"
	// nombres[1] = "zorra"
	// nombres[2] = "mapache"
	// nombres[2] = "mapache"
	// fmt.Println(nombres)

	//otra forma de crear un array
	nombres := [4]string{"annie",
		"alvaro",
		"alexys",
		"perra",
	}
	// nombre := nombres[2]
	// fmt.Println(nombre)
	// fmt.Println(nombres[1])
	//tamaño
	size := len(nombres)
	fmt.Println("el tamaño del arreglo es: ", size)
	//conocer una posiciones determinadas del arreglo
	//fmt.Println(nombres[inicio:final(excluyente)])
	//fmt.Println(nombres[1:3])
	fmt.Println(nombres)

}
