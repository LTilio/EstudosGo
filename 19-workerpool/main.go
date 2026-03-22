package main

import (
	"fmt"
	"net/http"
	"sync"
)

func fetchUrl(workerID int, url string) {
	fmt.Printf("\nWorker: %d\nComeçando uma reequisição para %s\n", workerID, url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Worker %d: Error ao acesso %s: %v\n", workerID, url, err)
	}
	defer resp.Body.Close()
	fmt.Printf("Worker %d: concluido para %s com status %d\n", workerID, url, resp.StatusCode)
}

func worker(id int, urls <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urls {
		fetchUrl(id, url)
	}
}

func main() {

	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.golang.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
	}

	const numWorker = 3 //numero de workers no pool

	urlChannel := make(chan string, len(urls))
	var wg sync.WaitGroup

	for i := 1; i <= numWorker; i++ {
		wg.Add(1)
		go worker(i, urlChannel, &wg)
	}

	for _, url := range urls {
		urlChannel <- url
	}

	close(urlChannel)

	wg.Wait()
	fmt.Println("Todas as req. foram concluidas")

}
