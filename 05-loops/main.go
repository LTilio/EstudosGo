package main

import "fmt"

func main() {

	fmt.Println("\nLoops")
	fmt.Println("Loop for exemplo basico")
	for i := 0; i < 5; i++ {
		fmt.Println("Valor de 'I' no for: ", i)
	}

	fmt.Println("\nLoop for como se fosse while")

	i := 0
	for i < 5 {
		fmt.Println("Valor de 'I' no for(como se fosse while): ", i)
		i++
	}

	fmt.Println("Loop infinito")
	for {
		fmt.Println("Loop infinito executando...")
		break
	}

	nums := []int{10, 20, 30}

	fmt.Println("\nLoop com range com slice")
	fmt.Println("OBS: para omitir uma variavel o padroa no go é '_'")
	fmt.Println("ex;")
	fmt.Println("for index, value := range nums {} | com o index")
	fmt.Println("for _, value := range nums {} | sem o index")
	for index, value := range nums {
		fmt.Printf("Indece: %d, Valor: %d\n", index, value)
	}

	fmt.Println("\nLoop com range com MAP")
	capitais := map[string]string{
		"Brasil": "Brasilia",
		"França": "Paris",
		"Japão":  "tokio",
	}

	for pais, capital := range capitais {
		fmt.Printf("Indece: %s, Valor: %s\n", pais, capital)

	}

}
