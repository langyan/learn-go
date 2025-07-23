package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// 存储键值对
	m.Store("name", "Alice")
	m.Store("age", 30)
	m.Store("country", "USA")

	// 加载值
	if name, ok := m.Load("name"); ok {
		fmt.Println("Name:", name) // 输出: Name: Alice
	}

	// 加载或存储
	if actual, loaded := m.LoadOrStore("age", 31); loaded {
		fmt.Println("Existing age:", actual) // 输出: Existing age: 30
	}

	// 删除键
	m.Delete("country")

	// 遍历所有键值对
	fmt.Println("All entries:")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})
	// 可能输出:
	// name: Alice
	// age: 30

	// 并发安全的计数器示例
	var counter sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 使用原子操作增加计数器
			if val, ok := counter.Load("count"); ok {
				counter.Store("count", val.(int)+1)
			} else {
				counter.Store("count", 1)
			}
		}()
	}

	wg.Wait()
	if count, ok := counter.Load("count"); ok {
		fmt.Println("Final count:", count) // 输出: Final count: 100
	}
}
