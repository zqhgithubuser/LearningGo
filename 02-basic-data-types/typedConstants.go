package main

import "fmt"

const (
	typedConstant   = int16(100)
	untypedConstant = 100
)

func main() {
	i := int(1)
	fmt.Println("unTyped:", i*untypedConstant)
	// invalid operation: i * typedConstant (mismatched types int and int16)
	// fmt.Println("Typed:", i * typedConstant)
	fmt.Println("Typed:", int16(i)*typedConstant)
}
