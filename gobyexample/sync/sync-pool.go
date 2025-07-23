package main

import (
	"fmt"
	"sync"
	"time"
)

func withoutPool() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每次创建新对象
			_ = make([]byte, 1024)
		}()
	}
	wg.Wait()
	fmt.Printf("Without Pool: %v\n", time.Since(start))
}

func withPool() {
	pool := &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}
	
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 从池中获取
			buf := pool.Get().([]byte)
			// 使用后放回
			defer pool.Put(buf)
		}()
	}
	wg.Wait()
	fmt.Printf("With Pool: %v\n", time.Since(start))
}

func main() {
	withoutPool()
	withPool()
}