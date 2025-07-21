package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

type Payment struct {
	Data []byte
}

func processPayments(count int) []Payment {
	payments := make([]Payment, count)
	for i := 0; i < count; i++ {
		// Simulate memory allocation
		payments[i] = Payment{Data: make([]byte, 1024*1024)} // 1MB per payment
		time.Sleep(1 * time.Millisecond)
	}
	return payments
}

func printMemoryStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Heap Alloc: %v MB\n", m.HeapAlloc/1024/1024)
	fmt.Printf("Total Alloc: %v MB\n", m.TotalAlloc/1024/1024)
}

func main() {
	f, err := os.Create("mem_profile.prof")
	if err != nil {
		fmt.Printf("Failed to create profile: %v\n", err)
		return
	}
	defer f.Close()

	start := time.Now()
	payments := processPayments(1000)
	runtime.GC() // Force garbage collection
	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Printf("Failed to write profile: %v\n", err)
		return
	}
	printMemoryStats()
	fmt.Printf("Processed %d payments in %v\n", len(payments), time.Since(start))
}
