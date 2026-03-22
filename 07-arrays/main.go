package main

import "fmt"

func main() {

	var numbers [5]int

	fmt.Println("Arrays")
	fmt.Println("Array criado, por default leva o valro de 0")
	fmt.Println(numbers)
	fmt.Println("Array alterado")
	numbers[0] = 10
	numbers[1] = 1
	numbers[2] = 50
	fmt.Println(numbers)
	fmt.Println("Array criado com os valores")
	dias := [3]string{"seg", "ter", "quar"}
	fmt.Println(dias)

	fmt.Println("Array criado SEM COLOCAR O TAMANHO DELE")
	notas := [...]int{55, 20, 45, 55}
	fmt.Println(notas)

}
