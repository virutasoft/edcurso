package main

import "fmt"

//RECOVER..es una funcion que recupera el control de la aplicacion a pesar de que haya un error dificil de recuperar
func main() {
	f()
}
func f() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Printf("%T\n",r)
			fmt.Println("RECUPERADO EN R: ", r)
		}
	}()
	fmt.Println("llamando a g()")
	g(7)
	fmt.Println("Retorna normalmente desde g()")
}

func g(i int) {
	if i > 3 {
		fmt.Println("AAAAAAAAAAAHHHHHH")
		panic("I ES MAYOR QUE TRES")
	}
	defer fmt.Println("Defer en la funcion g()")
	fmt.Println("Imprimiendo en g()", i)
}
// LOS RECOVER SE UTILIZAN DENTRO DE LAS FUNCIOnES DEFER
