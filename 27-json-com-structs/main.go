package main

import (
	"encoding/json"
	"fmt"
)

type Pokemon struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Weight    float64  `json:"weight"`
	Abilities []string `json:"abilities"`
	Stats     Stats    `json:"stats"`
}

type Stats struct {
	HP             int `json:"hp"`
	Attack         int `json:"attack"`
	Defense        int `json:"defense"`
	SpecialAttack  int `json:"special-attack"`
	SpecialDefense int `json:"special-defense"`
	Speed          int `json:"speed"`
}

func validateData(data Pokemon) error {

	if data.Name == "" {
		return fmt.Errorf("name is required")
	}
	if data.Type == "" {
		return fmt.Errorf("type is required")
	}
	if data.Weight == 0 {
		return fmt.Errorf("weight is required")
	}

	return nil
}

func main() {

	jsonData := `{
		"name": "snorlax", 
		"type": "normal",
		"weight": 460.0,
		"abilities": ["glutonous", "thick-fat"],
		"stats": {
			"hp": 160,
			"attack": 110,
			"defense": 65,
			"special-attack": 65,
			"special-defense": 110,
			"speed": 30
		}
	}`

	var data Pokemon
	err := json.Unmarshal([]byte(jsonData), &data)

	if err != nil {
		fmt.Println("Erro ao deserializar o JSON: ", err)
		return
	}

	err = validateData(data)
	if err != nil {
		fmt.Println("Erro ao validar os dados: ", err)
		return
	}

	fmt.Println(data)
}
