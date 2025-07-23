package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	name string
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{name: "My Singleton"}
		fmt.Println("Singleton instance created")
	})
	return instance
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			singleton := GetInstance()
			fmt.Printf("Singleton address: %p, name: %s\n", singleton, singleton.name)
		}()
	}

	wg.Wait()
	/*
	可能的输出:
	Singleton instance created
	Singleton address: 0x123456, name: My Singleton
	Singleton address: 0x123456, name: My Singleton
	Singleton address: 0x123456, name: My Singleton
	Singleton address: 0x123456, name: My Singleton
	Singleton address: 0x123456, name: My Singleton
	*/
}