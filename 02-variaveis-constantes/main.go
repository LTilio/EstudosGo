package main

import "fmt"

const name5 = "Constante"

func main() {

	var name1 string
	name1 = "Tilio | valor passado apois iniciar a variavel"

	var name2 string = "Tilio | variavel já iniciada com valor"

	var name3 = "Tilio | varivel sem declarar o tipo"

	name4 := "Tilio | varivel sem declarar o tipo e sem o var"

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)
	fmt.Println(name5)

	age := 33
	valor := 3.14
	ativo := true

	fmt.Println("\nAqui vem os tipo;\n")
	fmt.Println(age, "int")
	fmt.Println(valor, "float64")
	fmt.Println(ativo, "boolean")

	inline1, inline2, inline3 := "A", 5, false

	fmt.Println("\nDeclarando varias variaveis in line\n")
	fmt.Println(inline1)
	fmt.Println(inline2)
	fmt.Println(inline3)

}
