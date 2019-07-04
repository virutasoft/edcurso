package main

import "fmt"

func main() {
	//PUNTEROS
	//ES LA DIRECCION EN MEMORIA A UNA VARIABLE DE UN TIPO DETERMINADO

	a := 3
	// fmt.Println(&a) // para acceder a las direccion en memoria, uso el caracter (&) antepuesto a la variable o lo q necesite

	//el valor 0 de un puntero en NIL

	fmt.Println("Antes de duplicar, a= ", a)
	duplicar(&a)
	fmt.Println("Después de duplicar, a= ", a)

}
func duplicar(x *int) {
	*x = *x * 2
	/*		 *x *= *2 lo que sea que tenga ese valr, lo va a multiplicar x 2     */
	fmt.Println("DENTRO DE LA FUNCIÓN DUPLICAR x VALE ", *x)
}


//cuando le pasamos los valres a una funcion, se llama PASO POR VALOR
// cuando le pasamos punteros la func recibe punteros y el valor se llama PASO POR REFERENCIA