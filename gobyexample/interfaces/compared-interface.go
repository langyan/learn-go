package main

import (
	"fmt"
	"reflect"
)

func main() {
	equal()
	differentDynamicTypes()
	differentDynamicTypes2()
	UncomparableTypes()

}

func isEqual(a, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}
	return a == b
}

func equal() {
	var a interface{} = 10
	var b interface{} = 10
	fmt.Println(" Same Dynamic Type and Value: ", a == b) // true
}

func differentDynamicTypes() {
	type MyInt int
	var a interface{} = 10
	var b interface{} = MyInt(10)
	fmt.Println("1Different Dynamic Types: ", a == b) // false (different dynamic types)
}

func differentDynamicTypes2() {
	type MyInt int
	a := 10
	b := MyInt(10)
	fmt.Println(" 2Different Dynamic Types: ", isEqual(a, b)) // false (different dynamic types)
}

func UncomparableTypes() {
	var a interface{} = []int{1, 2, 3}
	var b interface{} = []int{1, 2, 3}
	fmt.Println("Uncomparable Types", isEqual(a, b)) // panic: comparing uncomparable type []int
}
