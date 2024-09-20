package main

import "fmt"

func main() {
	b := make([]byte, 12)
	fmt.Println("Byte slice:", b)

	// 字符串转换为字节切片
	b = []byte("Byte slice €")
	fmt.Println("Byte slice:", b)

	fmt.Printf("Byte slice as text: %s\n", b)
	// 字节切片转换为字符串
	fmt.Println("Byte slice as text:", string(b))
	// 字节数
	fmt.Println("Length of b:", len(b)) // 14
}
