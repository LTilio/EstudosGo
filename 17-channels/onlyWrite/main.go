package main

import "fmt"

func sendMessage(ch chan<- string) {
	ch <- "msg 01"
	ch <- "msg 02"
}

func main() {
	ch := make(chan string, 2)

	go sendMessage(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
