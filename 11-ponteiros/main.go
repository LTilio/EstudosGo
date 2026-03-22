package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	fmt.Println("& - obtem o enderço de uma variavel")
	fmt.Println("* - Acessa o valor armazenado no enderço que o ponteiro aponta")

	name := "Tilio"
	var pointer *string

	fmt.Println(name)
	fmt.Println(pointer)

	fmt.Println("Endereço de memoria das variaveis")
	fmt.Println(&name)
	fmt.Println(&pointer)

	fmt.Println("\npointer referenciando name")
	pointer = &name
	fmt.Println(pointer)
	fmt.Println(*pointer)

	fmt.Println("\nAltereção do valor de name pelo ponteiro")
	*pointer = "Tilio Muniz"
	fmt.Println(name)
	fmt.Println(*pointer)

	fmt.Println("\nAltereção do valor de name pela propria variavel name")
	name = "Leandro Tilio Muniz"
	fmt.Println(name)
	fmt.Println(*pointer)
	fmt.Println(&name)
	fmt.Println(pointer)
	fmt.Println("\nO endereço de memória de pointer é diferente do de name porque pointer é uma variável que armazena o endereço de name, e não o valor diretamente.")
	fmt.Println(&pointer)

}
