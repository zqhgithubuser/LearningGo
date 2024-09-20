package main

import (
	"fmt"
	"unsafe"
)

func byteToString(bStr []byte) string {
	if len(bStr) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(bStr), len(bStr))
}

func stringToByte(str string) []byte {
	if str == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

func main() {
	var hi string = "Hello, Go!"
	myByteSlice := stringToByte(hi)
	fmt.Printf("myByteSlice: %v\n", myByteSlice)
	fmt.Printf("myByteSlice type: %T\n", myByteSlice) // []uint8

	myStr := byteToString(myByteSlice)
	fmt.Printf("myStr: %v\n", myStr)
	fmt.Printf("myStr type: %T\n", myStr) // string
}
