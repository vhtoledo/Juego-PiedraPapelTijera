package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hola Mundo")
	})

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}