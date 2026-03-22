package main

import "fmt"

func unsafeop() {
	panic("algo errado ocorreu")
}
func safeop() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperando server...", r)

		}
	}()
	panic("algo errado ocorreu")
}

func main() {

	fmt.Println("Iniciou o server...")
	// unsafeop()
	safeop()
	fmt.Println("server down")

}
