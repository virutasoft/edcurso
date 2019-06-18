package main

import "fmt"

//creacion de MAPS
func main() {
	// idiomas["es"]= "español"
	// // forma mas comun
	// idiomas := make(map[string]string)
	// idiomas["es"] = "Español"
	// idiomas["en"] = "ingles"
	// idiomas["fr"] = "frances"
	// fmt.Println(idiomas)
	/*
		//declarar todo en una sola linea espaciada
		idioma := map[string]string{
			"es": "españolete",
			"en": "english",
			"fr": "franchute",
		}
		//FORMA 2 DE DECLARAR LOS MAPS

		//asignando directamente el valor
		idiomas := map[string]string{"es": "Español", "en": "ingles", "fr": "Francés", "pt": "portuguese"}

		fmt.Println(idiomas)
		fmt.Println(idiomas["es"])
		fmt.Println(idioma)
		//eliminar pos de un mapa
		delete(idioma, "en")
		fmt.Println(idioma)

		// los mapas devuelven el valor de la clave y un booleano, indicando si la clave existe o no existe
		if idioma1, ok := idiomas["pt"]; ok {
			fmt.Println("la posicion si existe y vale", idioma1)
		} else {
			fmt.Println("la posicion no existe")
		}
		fmt.Println(idioma["pt"])
	*/

	numeros := map[int]string{
		1:    "uno es un numero chiquito",
		2:    "esto es otro numero",
		2016: "un numero mas grande",
		-4:   "aqui es negativo",
	}
	fmt.Println(numeros)
}
