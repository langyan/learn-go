package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Producting : ", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for item := range ch {
		fmt.Println("Consuming:", item)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int, 3)
	go producer(ch)
	consumer(ch)
}
