package main

import (
	"fmt"

	"github.com/golang-es/edcurso-go/06/07paquetes/despedida"
	"github.com/golang-es/edcurso-go/06/07paquetes/saludar"
)

//se coloca la ruta   a partir de la carpeta SRC q se agregó en el PATH y ya
func main() {
	//paquetes

	//SIEMPRE SE RECOMIENDA USAR EL MISMO NOMBRE PARA LA CARPETA DEL PAQUETE TANTO COMO PARA E ARCHIVO PRINCIPAL DE ESTE PAQUETE

	nombre := "RoglinetekTrolinatik"
	saludar.Saludar("perra")
	saludar.Saludar(nombre)
	saludar.MeVes = "Esto es un texto usado desde el main"
	fmt.Println(saludar.MeVes)

	despedida.Despedirse("zorra")
}
