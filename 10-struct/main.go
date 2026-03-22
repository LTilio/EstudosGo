package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	Email    string
	Endereco Endereco
	baiscType
}

type Endereco struct {
	Rua    string
	Bairro string
	Cidade string
	Cep    string
}

type baiscType struct {
	Id string
}

func NewPerson(name, email, rua, bairro, cidade, cep, id string, age int) Person {
	p := Person{
		Name:  name,
		Age:   age,
		Email: email,
		Endereco: Endereco{
			Rua:    rua,
			Bairro: bairro,
			Cidade: cidade,
			Cep:    cep,
		},
	}
	p.baiscType.Id = id

	return p
}

func (p Person) Greet() string {
	return "Hello World, from " + p.Name
}

func main() {

	p := NewPerson("Tilio", "t@email.com", "bla", "centro", "Petropolis", "20202-55", "asijdoashdiuashda", 33)

	fmt.Println("Structs")

	p2 := Person{
		Name:  "Gaby",
		Age:   32,
		Email: "g@email.com",
		Endereco: Endereco{
			Rua:    "bla",
			Bairro: "centro",
			Cidade: "Petropolis",
			Cep:    "20202-55",
		},
	}

	persons := []Person{p, p2}

	fmt.Println(p)
	fmt.Println(p.Id)
	fmt.Println()
	fmt.Println("Nome: ", p.Name)
	fmt.Println("Idade: ", p.Age)
	fmt.Println("Email: ", p.Email)
	fmt.Println("Endereço: ", p.Endereco)
	fmt.Println()
	fmt.Println(p.Greet())
	fmt.Println()
	fmt.Println(persons)
	fmt.Println()

	for _, p := range persons {
		fmt.Println(p.Greet())
	}

}
