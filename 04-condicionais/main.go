package main

import (
	"fmt"
)

func main() {

	fmt.Println("\n Estruturas condicionais")
	fmt.Println("\n IF ELSE e IF")
	age := 20
	if age >= 13 {
		fmt.Println("Você é uma criança")
	} else if age < 18 {
		fmt.Println("Você é um adolecente")
	} else {
		fmt.Println("Você é um adulto")
	}

	day := ""

	switch day {
	case "domingo":
		fmt.Println("Final de semana")
	case "segunda":
		fmt.Println("Dia da semana")
	default:
		fmt.Println("N/A")
	}

}
