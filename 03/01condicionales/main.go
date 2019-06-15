package main

import "fmt"

func main() {
	/*
		isValid := true

		if 5 < 10 || 1 > 7 {
			//siempre la instruccion se va a ejecutar es si  la condicion es verdad
			fmt.Println("esto está en el bloque de verdadero")
		}
		fmt.Println("fin del programa")
		if isValid {
			fmt.Println("esto es verdadero")
		}

		//otra forma de hacer lo del if
		if isValid1 := true; isValid1 { //se usa para comparar los resultados que retorna la vairable
			fmt.Println("esto es del bloque de verdad")
		}
		var arroz int
		arroz = 8

		switch arroz {
		case 1:
			fmt.Println("el valor es 1")
			//no es aconsejable colocar el break
		case 2:
			fmt.Println("el valor es 2")
		default:
			fmt.Println("que le den")
		}
	*/
	//

	// switch a {
	// case 1:
	// 	fmt.Println("lunes")
	// case 2:
	// 	fmt.Println("martes")
	// case 3:
	// 	fmt.Println("miercoles")
	// case 4:
	// 	fmt.Println("jueves")
	// case 5:
	// 	fmt.Println("viernes")
	// case 6:
	// 	fmt.Println("sabado")
	// case 7:
	// 	fmt.Println("domingo")
	// default:
	// 	fmt.Println("lo siento, esta muerto")
	// }

	// switch a {
	// case 1:
	// 	fallthrough //significa si es este valor, siga evaluando, continue con la siguiemte
	// case 2:
	// 	fallthrough
	// case 3:
	// 	fallthrough
	// case 4:
	// 	fallthrough
	// case 5:
	// 	fmt.Println("estas entre semana")
	// case 6:
	// 	fallthrough
	// case 7:
	// 	fmt.Println("estas en fin de semana")
	// default:
	// 	fmt.Println("estas muerto")
	// }
	//el switch nos permite asignar una variable o varias dentro del mismo switch, pero no se puede ahi mismo, colocarle la variable q vamos a velidar
	switch a := 6; {
	case a >= 0 && a <= 5:
		fmt.Println("entre semana")
	case a >= 6 && a <= 7:
		fmt.Println("estas en fin de semana")
	default:
		fmt.Println("estas muerto")
	}
}
