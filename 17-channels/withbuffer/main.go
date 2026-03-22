package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 3)

	go func() {

		fmt.Println("escrevendo no channel com o buffer")
		ch <- "msg 1"
		fmt.Println("msg 1 --- ENVIADA")
		ch <- "msg 2"
		fmt.Println("msg 2 --- ENVIADA")
		ch <- "msg 3"
		fmt.Println("msg 3 --- ENVIADA")
		ch <- "msg 4"
		fmt.Println("msg 4 --- ENVIADA")
	}()

	time.Sleep(2 * time.Second)

	go func() {
		for i := 0; i < 4; i++ {
			msg := <-ch
			fmt.Printf("Recebido: %s\n", msg)
			time.Sleep(1 * time.Second)

		}
	}()

	time.Sleep(10 * time.Second)

}
