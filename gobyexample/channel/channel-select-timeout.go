package main

import (
	"fmt"
	"time"
)

func slowFunction(done chan<- string) {
	time.Sleep(3 * time.Second)
	done <- "done"
}
func main() {
	done := make(chan string)
	go slowFunction(done)
	select {
	case result := <-done:
		fmt.Println("Success:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Time out")

	}
}
