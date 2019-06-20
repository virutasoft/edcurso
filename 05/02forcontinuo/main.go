package main

import "fmt"

func main() {
	//for continuo
	//solo contien la expresion de condicion
	a := 5 // va por fuera de este for
	for a > 0 {
		a-- // va por dentro del for
		fmt.Println(a)
	}
}
