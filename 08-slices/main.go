package main

import "fmt"

func main() {

	fmt.Println("Slices")
	fmt.Println("Slices inicia sem numero de posições")

	fmt.Println("Declaração basico de slices")
	var numbers []int
	fmt.Println(numbers)

	fmt.Println("Slice ja iniciado com valores")
	dias := []string{"seg", "ter", "quar"}
	fmt.Println(dias)

	fmt.Println("Slice ja iniciado com valores")
	numbers2 := []int{2, 3, 5}
	fmt.Println(numbers2)
	fmt.Println("Slice com valores adicionados com append")
	numbers2 = append(numbers2, 60)
	fmt.Println(numbers2)

	fmt.Println("Usando um 'parte' do slice")
	numbers3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numbers3Sub := numbers3[3:6]
	fmt.Println(numbers3Sub)

	fmt.Println("Pegando a partir da posição 3")
	numbers3Sub1 := numbers3[:3]
	fmt.Println(numbers3Sub1)

	fmt.Println("Pegando da posição 5 em diante")
	numbers3Sub2 := numbers3[5:]
	fmt.Println(numbers3Sub2)

	fmt.Println("Deletando uam posição atraves de sub slices")
	numbersDelete3Index := append(numbers3[:2], numbers3[3:]...)
	fmt.Println(numbersDelete3Index)

	fmt.Println("Slices multidimencionas")
	matriz := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matriz)

	fmt.Println("Função make")
	numbers4 := make([]int, 5)
	fmt.Println(numbers4)
	numbers4[0] = 9
	fmt.Println(numbers4)
	fmt.Println("Tamanho do slice")
	fmt.Println(len(numbers4))
	fmt.Println("Capacidade do slice")
	fmt.Println(cap(numbers4))

	fmt.Println("Extrapolando a capacidade do slice")
	numbers5 := make([]int, 5, 10)
	fmt.Println(cap(numbers5))

	for i := 0; i < 6; i++ {
		numbers5 = append(numbers5, i)
	}
	fmt.Println(numbers5)
	fmt.Println(len(numbers5))
	fmt.Println(cap(numbers5))

}
