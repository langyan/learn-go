package main

import "fmt"

func zeroval(ival int) {
	ival = 0
	fmt.Println("zeroval-pointer-ival:", &ival)
}

func zeroptr(iptr *int) {
	*iptr = 0

	fmt.Println("zeroval-pointer-iptr:", iptr)
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	fmt.Println("initial-pointer:", &i)

	zeroval(i)
	fmt.Println("zeroval:", i)
	fmt.Println("zeroval-pointer:", &i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}
