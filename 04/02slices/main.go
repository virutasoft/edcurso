package main

import "fmt"

func main() {
	//SLICES
	// representan una secuencia de elelmentos de tamaño variable de elementos del mismo tipo
	//tienen un apuntador a un array
	//a diferencia de los arrays, el tamaño es indeterminado
	// 1. apuntador a un array
	// 2. siempre tiene un tamaño
	// 3. tiene una capacidad

	// al crear el slice, la capacidad siempre va a ser el doble de lo asignado 1→2, 3→6 y si sobrepasa en algun momento la capacidad del slice, este duplica la cantidad de memoria de lo q se sobrepasó
	//tenen¿mos un slice de 3 que duplica a 6, pero si yo al rellenar el slice, supero la cantidad de 6 elementos en 1 el slice tendría ya no 6 sino 7 elementos, ahí el slice duplica la cantidad y pasa a tener 14 espacios, y así en adelante

	//var nombres []string

	//hay otra forma mas legal de contruir los slices, es con la funcion prefabricada:
	//make(tipo de datos del slice, tamaño inicial, capacidad inicial)
	// nombres := make([]string, 0)
	// //nombres= append(nombre slice, "que info se agrega")
	// nombres = append(nombres, "perra")
	// fmt.Printf("el tamaño del slice es: %d , su capacidad es de: %d\n", len(nombres), cap(nombres))
	// nombres = append(nombres, "zorra")
	// fmt.Printf("el tamaño del slice es: %d , su capacidad es de: %d\n", len(nombres), cap(nombres))
	// nombres = append(nombres, "mapache")
	// fmt.Printf("el tamaño del slice es: %d , su capacidad es de: %d\n", len(nombres), cap(nombres))
	// nombres = append(nombres, "rogleton")
	// fmt.Printf("el tamaño del slice es: %d , su capacidad es de: %d\n", len(nombres), cap(nombres))
	// nombres = append(nombres, "bebe")
	// fmt.Printf("el tamaño del slice es: %d , su capacidad es de: %d\n", len(nombres), cap(nombres))
	// nombres = append(nombres, "viruta")
	// fmt.Printf("el tamaño del slice es: %d , su capacidad es de: %d\n", len(nombres), cap(nombres))
	// nombres = append(nombres, "tregué")
	// fmt.Printf("el tamaño del slice es: %d , su capacidad es de: %d\n", len(nombres), cap(nombres))
	// nombres = append(nombres, "lamp")
	// fmt.Printf("el tamaño del slice es: %d , su capacidad es de: %d\n", len(nombres), cap(nombres))
	// nombres = append(nombres, "vaselinicon")
	// fmt.Printf("el tamaño del slice es: %d , su capacidad es de: %d\n", len(nombres), cap(nombres))

	// nombres[4] = "frozinatic"

	//fmt.Println(nombres)

	//otra forma de crear slices pero ya con los valores ingresados es asi:
	nombres := []string{
		"alavaro",
		"zorra",
		"lamedusicona",
	}
	fmt.Println(nombres)
}
