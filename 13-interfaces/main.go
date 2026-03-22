package main

import "fmt"

type Mamifero interface {
	saudacao()
	hobbie()
}

type Humano struct {
	Nome string
	Prof string
}

func (h Humano) saudacao() {
	fmt.Println("Olá, meu nome é " + h.Nome + " e eu sou " + h.Prof)
}

func (h Humano) hobbie() {
	fmt.Println("Video games")
}

type Cachorro struct {
	Nome string
}

func (c Cachorro) saudacao() {
	fmt.Println("au au")
}

func (c Cachorro) hobbie() {
	fmt.Println("latir")
}

type Gato struct {
	Nome        string
	TempodeSono int
}

func (g Gato) saudacao() {
	fmt.Println("miau")
}

func (g Gato) hobbie() {
	fmt.Println("dormir por ", g.TempodeSono)
}

func main() {
	fmt.Println("Interfaces")
	fmt.Println()

	eu := Humano{
		Nome: "Tilio",
		Prof: "dev",
	}

	gato := Gato{
		Nome:        "Pastel",
		TempodeSono: 12,
	}

	cachorro := Cachorro{
		Nome: "tunico",
	}

	apr(eu)
	apr(gato)
	apr(cachorro)

}

func apr(m Mamifero) {
	m.saudacao()
	m.hobbie()
}
