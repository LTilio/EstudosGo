package main

import (
	"fmt"
	"time"
)

func makeCoffe() {
	fmt.Println("Fazendo o café")
	time.Sleep(2 * time.Second)
	fmt.Println("Café pronto")
}

func makebread() {
	fmt.Println("Fazendo o pão")
	time.Sleep(3 * time.Second)
	fmt.Println("pão pronto")
}

func fryEggs() {
	fmt.Println("Fazendo o ovos")
	time.Sleep(1 * time.Second)
	fmt.Println("ovos pronto")
}

func main() {
	fmt.Println("Go routines")
	fmt.Println()

	start := time.Now()
	//rodando de maneira concorrente
	go makeCoffe()
	go makebread()
	go fryEggs()

	time.Sleep(4 * time.Second)

	fmt.Println("Café da manhã pronto!")
	fmt.Printf("Tempo total: %v\n", time.Since(start))

}
