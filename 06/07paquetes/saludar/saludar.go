package saludar

import "fmt"

// MeVes es para utilizar desde otro paquete
var MeVes string

var noMeVes string

// Saludar saluda a una persona
func Saludar(nombre string) {
	//se usa la S y significa q va a ser una funcion exportada y que se puede usar desde otro paqwuete
	fmt.Println("Hola ", nombre)
}

func noVisible() {
	fmt.Println("No me ven desde otro paquete porque la variable no inicia en mayúscula")
}
