package main

import "fmt"

//DEFER: se ejecutan en el orden inverso al que fueron agregados, y sirven para diferir la ejecución de un proceso o un método, aún cuando el programa haya entrado en ´pánico
func main() {
	fmt.Println("Conectando a la base de datos...")
	defer cerrarDB()

	fmt.Println("Consultando la información del set de datos")
	defer cerrarSetDeDatos()

}

func cerrarDB() {
	fmt.Println("Cerrar la base de datos")

}
func cerrarSetDeDatos() {
	fmt.Println("Cerrar set de datos")
}
