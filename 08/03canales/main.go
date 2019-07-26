package main

import "fmt"

//LOS CANALES SE CREAR PARA COMUNICAR LAS DIFERENTES GOROUTINES
//PERMITEN LA COMUNICACION ENTRE GORUTINAS PARA SINCRONIZARSE ENTRE ELLAS E INTERCAMBIEN INFORMACION
func main() {
	bodegaOrigen := []string{"PHP", "JAVASCRIPT", "HTML", "GO", "RUST", "RUBY", "TYPESCRIPT", "C++", "C#", "JAVA", "KOTLIN", "FORTRAN", "BASES DE DATOS", "GIT"}
	bodegaDestino := []string{}

	fotocopiadora := make(chan string) //por lo general los canales se crean con la funcion make()
	fotocopiado := make(chan string)

	go func() {
		for _, libro := range bodegaOrigen {
			fotocopiadora <- libro
		}
	}()

	//el FOREVER se usa principalmente para trabajar con canales
	//se encarga la fotocopiadora

	//se deja en la bodega destino
	go func() {
		for {
			//entrega el contenido de la fotocopiadora a la varable LIBRO
			libro := <-fotocopiadora
			fotocopiado <- libro

			//agregar el libro a la bodega destino

			bodegaDestino = append(bodegaDestino, libro)
			if len(bodegaOrigen) == len(bodegaDestino) {
				//cerrar el canal fotocopiado
				close(fotocopiado)
			}
		}
	}()

	fmt.Println("***** LISTADO DE LIBROS FOTOCOPIADOS ****")
	for {
		if libro, ok := <-fotocopiado; ok {
			fmt.Println(libro)
		} else {
			break
		}

	}

}
