package main

import (
	"fmt"
	"sync"
	"time"
)

func makeCoffe(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Fazendo o café")
	time.Sleep(2 * time.Second)
	fmt.Println("Café pronto")
}

func makebread(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Fazendo o pão")
	time.Sleep(3 * time.Second)
	fmt.Println("pão pronto")
}

func fryEggs(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Fazendo o ovos")
	time.Sleep(1 * time.Second)
	fmt.Println("ovos pronto")
}

func main() {
	fmt.Println("Go routines")
	fmt.Println()
	var wg sync.WaitGroup
	wg.Add(3)

	start := time.Now()
	//rodando de maneira concorrente
	go makeCoffe(&wg)
	go makebread(&wg)
	go fryEggs(&wg)

	wg.Wait()

	fmt.Println("Café da manhã pronto!")
	fmt.Printf("Tempo total: %v\n", time.Since(start))

}
