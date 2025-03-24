package main

import "fmt"

func main() {
	printSimple()
}

func printSimple() {
	var a interface{} = 42             // stores int
	var b interface{} = "hello"        // stores string
	var c interface{} = []int{1, 2, 3} // stores slice

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
