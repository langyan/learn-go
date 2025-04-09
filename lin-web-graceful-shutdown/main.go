package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func runServer(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("正在正常关闭...")
			return
		default:
			fmt.Println("正在处理...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go runServer(ctx)
	<-ctx.Done()
	fmt.Println("服务已停止。")
}
