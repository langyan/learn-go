package main

import (
	"fmt"
	"time"
)

func main() {
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	defer close(jobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	for i := 1; i <= numJobs; i++ {
		<-results
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
