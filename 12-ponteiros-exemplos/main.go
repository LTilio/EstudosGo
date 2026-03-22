package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) showName() {
	p.Name = "Gaby"
	fmt.Println(p.Name)
	fmt.Println(&p.Name)
}

func main() {
	fmt.Println("Exemplos de Ponteiros")
	fmt.Println("& - obtem o enderço de uma variavel")
	fmt.Println("* - Acessa o valor armazenado no enderço que o ponteiro aponta")

	fmt.Println()
	fmt.Println("Ponteiros com função")
	age := 33
	fmt.Println(&age)
	showAge(&age)
	fmt.Println(age)

	fmt.Println()
	fmt.Println("Ponteiros com metodos")
	fmt.Println()

	p := Person{
		Name: "Tilio",
	}

	fmt.Println(p.Name)
	fmt.Println(&p.Name)
	fmt.Println()
	p.showName()

}

//Passagenm por valor = age
//passagem por referencias = age *int
func showAge(age *int) {
	fmt.Println(age)  // endereço
	fmt.Println(*age) // valor
	*age = 45
}
