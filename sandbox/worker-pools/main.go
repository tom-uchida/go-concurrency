package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

const (
	baseURL    = "https://openlibrary.org"
	endpoint   = "/api/books"
	numWorkers = 2
)

type Job struct {
	URL string
}

func main() {
	c := new(http.Client)

	urls := []string{
		baseURL + endpoint + "?bibkeys=ISBN:9780134190440&format=json",
		baseURL + endpoint + "?bibkeys=ISBN:978-1491941195&format=json",
		baseURL + endpoint + "?bibkeys=ISBN:0134494164&format=json",
	}

	jobs := make(chan Job, len(urls))
	results := make(chan string, len(urls))
	var wg sync.WaitGroup

	// Start workers
	for w := 0; w < numWorkers; w++ {
		go worker(&wg, c, jobs, results)
	}

	// Sending jobs to the worker pool
	wg.Add(len(urls))
	for _, url := range urls {
		jobs <- Job{URL: url}
	}
	close(jobs)
	wg.Wait()

	// Collecting results
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-results)
	}
}

func worker(wg *sync.WaitGroup, c *http.Client, jobs <-chan Job, results chan<- string) {
	for job := range jobs {
		req, _ := http.NewRequest(http.MethodGet, job.URL, nil)
		resp, _ := c.Do(req)

		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		var response interface{}
		json.Unmarshal(body, &response)

		results <- fmt.Sprintf("%d: %s", resp.StatusCode, response)
		wg.Done()
	}
}
