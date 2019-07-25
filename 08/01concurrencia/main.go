package main

import (
	"fmt"
	"time"
)

func main() {
	var h string

	//creamos la gorutine
	go MostrarNumeros()
	fmt.Println("Digita algo... ")
	fmt.Scan(&h)
	fmt.Println("Digitaste: ", h)
}

// MostrarNumeros muestra numeros del 1 al 10
func MostrarNumeros() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
