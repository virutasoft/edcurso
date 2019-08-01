package main

import (
	"log"
	"net/http"
)

//import("net/http")
func main() {
	/*
		//crear un servidor de archivos estaticos
		http.Handle("/", http.FileServer(http.Dir("public")))
		log.Println("Ejecutando server en http://localhost:8080")
		log.Println(http.ListenAndServe(":8080", nil))
	*/
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	log.Println("Ejecutando servidor desde el multiplexor en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
	
}
