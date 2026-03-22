package main

import (
	"fmt"
	"os"
)

func writeFile() {
	file, err := os.Create("example.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fprintln(file, "Hello World!")
	fmt.Println("Arquivo escrito corretamente")
}

func main() {
	writeFile()
}
