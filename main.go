package main

import (
	//"fmt"
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {
	// Crear enrutador
	router := http.NewServeMux()

	// Configurar rutas
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}