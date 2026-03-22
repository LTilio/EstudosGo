package main

import "fmt"

func readMessage(ch <-chan string) {

	msg1 := <-ch
	msg2 := <-ch
	fmt.Println("menssagens recebidos, ", msg1, " e ", msg2)

}

func main() {
	ch := make(chan string, 2)

	ch <- "M1"
	ch <- "M2"

	readMessage(ch)

}
