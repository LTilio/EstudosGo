package main

import "fmt"

func main() {

	fmt.Println("Operadores Aritimeticos:")

	sum := 10 + 3
	sub := 30 - 5
	div := 45 / 7
	mul := 56 * 23

	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println(div)
	fmt.Println(mul)

	fmt.Println("Operadores Relacionais:")
	age := 25
	fmt.Println(34 > 4)
	fmt.Println(34 < 4)
	fmt.Println(age > 18) // >, >=, ==, <, <=

	fmt.Println("Operadores de Atribuição:")

	x := 10
	x += 5
	x *= 5
	y := 20
	y -= 5
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("Operadores logicos:")
	usepass := "123"
	userAdm := true

	if usepass == "123" && userAdm == true {
		fmt.Println("Logado")
	} else {
		fmt.Println("negado")
	}

	// bit a bit
	fmt.Println("Operadores bit a bit:")
	a := 10 //1010 em binario
	b := 3  // 0011 em binario

	result := a & b
	fmt.Println(result)

}
