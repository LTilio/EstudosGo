package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("erro, nao pode dividir por 0")
	}
	return a / b, nil
}

func openfile(name string) error {
	if name == "" {
		return fmt.Errorf("Erro ao abrir o arquivo: %w\n", errors.New("nome do aquivo vazio"))

	}
	return nil
}

func main() {

	// result, err := divide(10, 0)

	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }

	// fmt.Println(result)
	// fmt.Println(err)

	err := openfile("")
	if err != nil {
		fmt.Println(err)
		return
	}

}
