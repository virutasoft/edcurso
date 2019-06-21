package main

import "fmt"

func main() {
	//funciones con return
	// var n1, n2 int8
	// n1 = 15
	// n2 = 100
	// r := suma(n1, n2)
	// fmt.Println(r)
	// var edad uint8
	// edad = 58
	//retornoEdad := tipoEdad(edad)
	//fmt.Println(retornoEdad)

	//o de esta forma mas abreviada
	//fmt.Println(tipoEdad(edad))
	n := []int8{58, 24, 32, 45, 12, 35, 96, 110, 20, -5, -9}
	maximo, minimo := maxymin(n)
	fmt.Println("maximo: ", maximo)
	fmt.Println("minimo: ", minimo)
}

//func nombreFuncion(argumentos+tipo de datos)tipo de dato q va a retornar(si tiene retorno){
// ejecución

// func suma(a, b int8) int8 /*se conoce como la firma de la funcion*/ {
// 	return a + b
// }

// func tipoEdad(edad uint8) string {
// 	var tipo string

// 	switch {
// 	case edad < 12:
// 		tipo = "niñez"
// 	case edad < 18:
// 		tipo = "adolescencia"
// 	case edad <= 50:
// 		tipo = "adultez"
// 	case edad > 50:
// 		tipo = "ancianidad"
// 	default:
// 		tipo = "mentira!!!"

// 	}
// 	return tipo
// }

// RETORNO DE MULTIPLES VALORES, si la funcion va a retornar varios valores, debemos colocarlos en tre paréntesis
/*
	func maxymin(numeros []int8) (int8, int8) {
		var max, min int8
		for _, v := range numeros {
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}
		return max, min
	}
*/

//de forma mas sencilla se hace lo siguiente
func maxymin(numeros []int8) (max int8, min int8) {
	// var max, min int8  no es necesario declararlos porque ya se declararon en la firma de la funcion
	for _, v := range numeros {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return //no le colocamos los valores max y min porque ya se declaaron en la firma de la funcion
}
