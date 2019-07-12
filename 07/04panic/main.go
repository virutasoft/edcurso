package main

import "fmt"

//panic muestra la fila de seguimiento de un problema inesperado, detien la ejecucion del programa pero deja que DEFER haga su trabajo siempre y cuando esten antes del PANIC
func main() {
	division(3, 0)
}

func division(dividendo, divisor int) {
	defer fmt.Println("Estos se ejecutará pase lo q pase")
	if divisor == 0 {
		panic("No se puede dividir por cero")
	}
	fmt.Println(dividendo / divisor)
}
