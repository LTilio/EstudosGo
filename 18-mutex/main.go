package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func addCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {

	var wg sync.WaitGroup
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go addCounter(&wg)
	}
	wg.Wait()
	fmt.Println(counter)

}
