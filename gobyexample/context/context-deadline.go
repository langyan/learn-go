package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(4*time.Second))
	defer cancel()
	go Watch(ctx, "goroutine1")
	go Watch(ctx, "goroutine2")

	time.Sleep(6 * time.Second)

	fmt.Printf("end watching")
	time.Sleep(1 * time.Second)

}

func Watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s exit !\n", name)
			return
		default:
			fmt.Printf("%s watching ...\n", name)
			time.Sleep(time.Second)
		}
	}
}
