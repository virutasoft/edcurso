package main

import (
	"errors"
	"fmt"
)

func main() {
	//ERRORES
	//son un tipo de datos como los enteros o como los string
	//NIL valor 0 de tipo error es como si fuera nulo
	r, err := division(24, 0)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Println(r) // lo raro de que el resultado de la división por cero sea cero es que el valor por defeecto del tipo de dato resultado es cero
	//fmt.Printf("%T\n", err)

}

func division(dividendo, divisor float64) (resultado float64, err error) {
	if divisor == 0 {
		err = errors.New("NO SE PUEDE DIVIDIR POR CERO")
	} else {
		resultado = dividendo / divisor
	}
	return
}
