package main

import (
	"fmt"
	"math"
)

// 全局变量
var Global int = 1234
var AnothaGlobal = -5678

func main() {
	var j int
	// 短赋值语句
	i := Global + AnothaGlobal
	fmt.Println("Initial j value:", j)

	j = Global
	k := math.Abs(float64(AnothaGlobal))
	fmt.Printf("Global=%d, i=%d, j=%d k=%.2f.\n", Global, i, j, k)
}
