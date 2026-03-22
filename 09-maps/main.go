package main

import "fmt"

func main() {
	fmt.Println("Maps")

	fmt.Println("Forma basica de declaração")
	var m map[string]int
	m = make(map[string]int)
	m["Tilio"] = 33
	m["Gaby"] = 32

	fmt.Println("Pegando um valro no map")
	fmt.Println(m["Tilio"])

	age, exists := m["Tilio"]
	if exists {
		fmt.Println("Tilio nao existe no map")
		fmt.Println(age)
	}

	fmt.Println("Deletando um map")
	fmt.Println("Se passo uma chave que nao existe no map ele nao faz nada so ignora")
	delete(m, "sdasdasd")
	fmt.Println(m)

	fmt.Println("Forma mais ultilizda")
	capitais := make(map[string]string)
	capitais["sao paulo"] = "sao paulo"
	capitais["ceara"] = "fortaleza"
	capitais["pernanbuco"] = "recife"

	fmt.Println(capitais)
	fmt.Println(capitais["pernanbuco"])

	for key, value := range capitais {
		fmt.Println()
		fmt.Println(key)
		fmt.Println(value)
		fmt.Println()
	}

	fmt.Println("Declaração e inicialização com chaves e valores")
	idade := map[string]int{"Tilio": 33, "Gaby": 32}
	fmt.Println(idade)
	fmt.Println()

	fmt.Println("Declaração anilhados")
	innerMaps := make(map[string]map[string]int)
	innerMaps["outer"] = make(map[string]int)
	innerMaps["outer"]["inner"] = 30
	fmt.Println(innerMaps["outer"])
	fmt.Println(innerMaps["outer"]["inner"])
	fmt.Println(innerMaps["inner"])
	fmt.Println()

	s := make(map[string]string, 30)
	fmt.Println(len(s))

}
