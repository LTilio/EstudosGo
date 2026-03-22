package main

import "fmt"

func main() {
	fmt.Println("\nFunções")
	fmt.Println(soma(1, 5))

	fmt.Println("\nFunções retornando multiplos valores")
	fmt.Println(divide(10, 3))

	fmt.Println("\nFunções sem retorno")
	noReturn("world")

	fmt.Println("\nFunções Anonima")
	anon := func() int {
		return 10
	}

	fmt.Println(anon())

	fmt.Println("\nFunções Closure")
	nextOdd := oddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	NewnextOdd := oddGenerator()
	fmt.Println(NewnextOdd())

	fmt.Println("\nFunções Variaticas")
	fmt.Println(max(5, 63, 7, 99, 10))

}

func soma(a, b int) int {
	return a + b
}

func divide(a, b int) (int, int) {

	quotient := a / b
	remainder := a % b
	return quotient, remainder

}

func noReturn(name string) {
	fmt.Println("Hello ", name)
}

func oddGenerator() func() int {
	i := -1
	return func() int {
		i += 2
		return i
	}
}

func max(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	bigger := numbers[0]
	for _, num := range numbers {
		if num > bigger {
			bigger = num
		}
	}
	return bigger
}
