package main

import "fmt"

func main() {
	aString := "Hello World! €"
	fmt.Println("First byte", string(aString[0])) // 索引访问

	r := '€'                             // A rune
	fmt.Println("As an int32 value:", r) // 8364
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	for _, v := range aString {
		fmt.Printf("%x ", v) // rune（十六进制）
	}
	fmt.Println()

	for _, v := range aString {
		fmt.Printf("%c ", v) // 字符
	}
	fmt.Println()
}
