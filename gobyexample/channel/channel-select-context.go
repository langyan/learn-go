package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine stopped")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}

}

func main() {

	cxt, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go longRunningTask(cxt)
	time.Sleep(3 * time.Second)
}
