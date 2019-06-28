package main

import "fmt"

func main() {
	//FUNCIONES ANOINIMAS : SON UN TIPO DE DATO, una FA mantiene disponible el valor de sus variables para la ejecución de si misma, incluso puede ser auto ejecutada, como en js

	//SINTAXIS FUNCIÓN ANÓNIMA
	//func (){
	//contenido
	//}()
	//COMO NO SE PUEDE LLAMAR EN ALGUN LADO PORQUE ES ANÓNIMA, SE ESCRIBEN LOS PARÉNTESIS AL FINAL PARA QUE SE EJECUTE SOLA o SE LE PUEDE ASIGNAR A UNA VARIABLE Y QUITARLE LOS 2 PARÉNTESIS

	// func() {
	// 	fmt.Println("Me imprimo")
	// }()

	//---------------------------------------------

	//se la asignamos a una variable
	/*
		anonima := func() {
			fmt.Println("hola perra")
		}

		//imprimimos la funcion en el scope de ella o donde se encuentre y la ejecuto

		anonima()
	*/
	secuencia1 := secuencia()
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println("---------------------------")
	secuencia2 := secuencia()
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
}

func secuencia() func() int32 {
	var x int32
	return func() int32 {
		x++
		return x
	}
}
