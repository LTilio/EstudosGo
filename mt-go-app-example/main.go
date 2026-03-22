package main

import (
	"fmt"

	"github.com/Rhymond/go-money"
	"github.com/ltilio/mt-go-app-example/utils"
)

func main() {
	fmt.Println("Hello World...")
	fmt.Println("Server up!")

	fmt.Println("func do mesmo modulo")
	// Chamando a func do mesmo modulo
	fmt.Println(sum(5, 5))

	fmt.Println("func do modulo utils")
	// Chamando a func do modulo utils
	// (para a func se 'publica sem que estar com a priemra letra maiuscula)
	fmt.Println(utils.Sum(5, 4))

	fmt.Println("func de lib externa")
	fmt.Println(6.50)

	moeda := money.NewFromFloat(6.50, money.BRL)
	fmt.Println(moeda.Display())

}
