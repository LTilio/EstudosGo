package main

import "fmt"

type ErrorCusTom struct {
	TaErrado string
}

func (ec *ErrorCusTom) Error() string {
	return fmt.Sprintf("\nErrou burrão! %s\n", ec.TaErrado)
}

func pegaError(qualquerCoisa string) error {
	if qualquerCoisa == "" {
		return &ErrorCusTom{""}
	}
	return nil
}

func main() {

	err := pegaError("")
	if err != nil {
		fmt.Println(err)
	}
}
