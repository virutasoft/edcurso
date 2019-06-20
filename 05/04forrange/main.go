package main

import "fmt"

func main() {
	//for range seria equivalente al foreach de otros lenguajes de programacion
	//numeros := []string{"cero", "uno", "dos", "tres", "cuatro", "cinco"} //slice de strings
	//nombres := map[string]string{"es": "españa", "co": "colombia", "ru": "rumania"}
	// for i, v := range numeros {
	// 	fmt.Println(i, v)
	// }
	//para ignorar la variable y q no bote error, puedo colocar cun "_" en lugar de esta y así no va a botar error al compilar, así no se imprimirían los índices
	// for _, v := range numeros {
	// 	fmt.Println(v)
	// }
	//se puede OMITIR el valor, mas no las posiciones de los valores
	// for i := range numeros {
	// 	fmt.Println(i)
	// }

	//vamos a recorrer el mapa
	// 	for i, v := range nombres {
	// 		fmt.Println(i, v)
	// 	}
	// tambien podemos iterar frases o palabras

	// frase := "hola mundo"
	// for posicion, letra := range frase {
	// 	fmt.Println(posicion, string(letra))
	// }
	//PODEMOS AGREGAR UN SLICE O UN TIPO DE DATO VÁLIDO PARA EL FOR RANGE
	// for _, entero := range []int{15, 24, 36, 58, 12} {
	// 	fmt.Println(entero)
	// }
	for _, letra := range "CURSO GO DESDE CERO" {
		fmt.Println(string(letra))
	}
}
