package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//runeCountInStringFun() // 输出: 字节长度: 13

	fmt.Println("-------------------------------------------")

	//runeFun() // 输出: 字节长度: 13
	fmt.Println("-------------------------------------------")

	const s = "你好"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func runeCountInStringFun() {
	a := "你好, world!"
	runeCount := utf8.RuneCountInString(a)
	fmt.Println("Rune 数量:", runeCount) // 输出: Rune 数量: 9
	fmt.Println("字节长度:", len(a))
}

func runeFun() {
	b := "你好, world!"
	runes := []rune(b)
	runeCount := len(runes)
	fmt.Println("Rune 数量:", runeCount) // 输出: Rune 数量: 9
	fmt.Println("字节长度:", len(b))
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
