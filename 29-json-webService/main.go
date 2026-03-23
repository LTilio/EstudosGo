package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request received")
	var pokemon Pokemon
	// Decode lê JSON do *corpo* da requisição. GET no navegador/curl costuma vir sem body → EOF.
	err := json.NewDecoder(r.Body).Decode(&pokemon)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("pokemon received: ", pokemon)
	fmt.Fprintf(w, "pokemon received: %+v\n", pokemon)
}

func main() {

	fmt.Println("starting server on port 8080")
	http.HandleFunc("/sub", handle)
	http.ListenAndServe(":8080", nil)

}
