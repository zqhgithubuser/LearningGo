package main

import (
	"fmt"
	"math"
)

func main() {
    i := math.MaxInt - 100
	for {
		if i == math.MaxInt {
			break
		}
		i++
	}
    fmt.Println("Max:", i)

    i = math.MinInt + 1000
    for {
        if i == math.MinInt {
            break
        }
        i--
    }
    fmt.Println("Min:", i)
}
