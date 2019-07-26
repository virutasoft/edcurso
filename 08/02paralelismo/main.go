package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(3)
	var wg sync.WaitGroup

	//el paralelismo nos permite decirle a go que use una cantidad maxima de procesadores para  que trabaje.
	numbers := []uint32{142857, 25, 428571, 32584784, 65847885, 21455, 255544, 4, 3333333, 13}
	wg.Add(len(numbers))
	fmt.Println("Comienza el  proceso......")
	for _, v := range numbers {
		go func(a uint32) {
			defer wg.Done()
			fmt.Println(a, EsPrimo(a))
		}(v)

	}
	wg.Wait()
	fmt.Println("Terminó wl proceso...")
}

// EsPrimo nos devuelve unarespuesta, si el numero es primo o no
func EsPrimo(a uint32) bool {
	c := 0
	var i uint32
	for i = 1; i <= a; i++ {
		if a%i == 0 {
			c++
		}
	}
	if c == 2 {
		return true
	}
	return false
}
