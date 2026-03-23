package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{
	"name": "snorlax", 
	"type": "normal",
	"weight": 460.0,
	"stats": {
		"hp": 160,
		"attack": 110,
		"defense": 65,
		"special-attack": 65,
		"special-defense": 110,
		"abilities": ["gluttony", "thick-fat"]
		}
	}`

	// var data map[string]interface{}
	data := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonData), &data)

	if err != nil {
		fmt.Println("Erro ao deserializar o JSON: ", err)
		return
	}

	fmt.Println("Nome: ", data["name"])
	fmt.Println("Tipo: ", data["type"])

	stats := data["stats"].(map[string]interface{})
	fmt.Println("abilities: ", stats["abilities"])
	fmt.Println("hp: ", stats["hp"])
	fmt.Println("defense: ", stats["defense"])
	fmt.Println("special-attack: ", stats["special-attack"])
	fmt.Println("special-defense: ", stats["special-defense"])
	fmt.Println("speed: ", stats["speed"])

}
