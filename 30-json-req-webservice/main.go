package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
	// Forms []struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
	// } `json:"forms"`
}

func main() {
	fmt.Println("Starting request to Pokemon API")
	fmt.Println("")
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/143")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(pokemon)
}
