package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola este es un handlerfunc</14>")
}

//son un tipo de datos q permite convertir funciones en manejadores http.
func messageHtCustom(message string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, message)
		},
	)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", messageHandlerFunc)
	mux.Handle("/saludar", messageHtCustom("<h1>Hola perra</h1>"))
	mux.Handle("/despedirse", messageHtCustom("<h1>Chao perra</h1>"))

	log.Println("Ejecutando servidor desde el multiplexor en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
