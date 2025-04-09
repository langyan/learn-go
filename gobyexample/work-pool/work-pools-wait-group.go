package main

import (
	"fmt"
	"sync"
	"time"
)

const maxWorkers = 5

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d \n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
		fmt.Printf("Workder %d finished job %d \n", id, j)
	}
}
func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	var wg sync.WaitGroup

	for w := 1; w < maxWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j < 10; j++ {
		jobs <- j
	}

	close(jobs)
	wg.Wait()
	close(results)
	for res := range results {
		fmt.Println("Results:", res)
	}
}
