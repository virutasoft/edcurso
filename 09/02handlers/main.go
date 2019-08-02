package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

func (m messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}
func main() {
	//handlers: nos permiten crear los manejadores para las rutas que diseñemos en lña app, es una interfaz q contiene un mmetodo que nos permite recibir como paramaetro unj responsewriter y qw tambien recibe un request
	mux := http.NewServeMux()
	m1 := &messageHandler{message: "<h1>Hola desde un handler</h1>"}
	m2 := &messageHandler{message: "<h1>Estos son perritos</h1>"}
	m3 := &messageHandler{message: "<h1>Estos son gatitos</h1>"}

	mux.Handle("/saludar", m1)
	mux.Handle("/perros", m2)
	mux.Handle("/gatos", m3)
	log.Println("Ejecutando servidor desde el multiplexor en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
