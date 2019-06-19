package main

import "fmt"

func main() {
	// //for clasico
	// for i := 10; i > 0; i-- {
	// 	fmt.Println(i)
	// }
	// for j := 10; j > 0; j-- {
	// 	if j == 3 || j == 6 {
	// 		continue // nos permite saltar el resto de instrucciones de la iteracion actual
	// 	}
	// 	fmt.Println("el valor es: ", j)
	// }
	// for k := 0; k < 10; k++ {
	// 	if k == 7 {
	// 		break // nos permite romper el ciclo y parar la ejecucion
	// 	}
	// 	fmt.Println("el valor de k es: ", k)
	// }

	//for con estructuras de datos
	//var matriz [3][3]
	matriz := [3][3]int{}
	valor := 0
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			valor++
			matriz[i][k] = valor
		}
	}

	for fila := 0; fila < 3; fila++ {
		for columna := 0; columna < 3; columna++ {
			fmt.Println(matriz[fila][columna])
		}

	}

}
