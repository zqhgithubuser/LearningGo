package main

import (
	"fmt"
)

func main() {
	slice := make([]byte, 3)
	// 指向底层数组的指针
	arrayPtr := (*[3]byte)(slice)
	fmt.Println("Print array pointer:", arrayPtr) // &[0 0 0]
	fmt.Printf("Data type: %T\n", arrayPtr)       // *[3]uint8
	fmt.Println("arrayPtr[0]:", arrayPtr[0])

	slice2 := []int{-1, -2, -3}
	// 转换为底层数组
	array := [3]int(slice2)
	fmt.Println("Print array contents:", array) // [-1 -2 -3]
	fmt.Printf("Data type: %T\n", array)        // [3]int
}
