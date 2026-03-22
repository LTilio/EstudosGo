package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Chanel sem buffer")

	ch := make(chan string) // declaração do chanel

	go func() {
		fmt.Println("Escrevendo no channel")
		ch <- "message 1" // escredndo no channel
		fmt.Println("menassagem enviada")
	}()

	time.Sleep(2 * time.Second)

	go func() {
		msg := <-ch
		fmt.Printf("Recebido: %s\n", msg)
	}()

	time.Sleep(2 * time.Second)

}
