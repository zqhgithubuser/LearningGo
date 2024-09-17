package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enough arguments")
		return
	}

	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _, v := range arguments[1:] {
		_, err := strconv.Atoi(v)
		if err == nil {
			total++
			nInts++
			continue
		}

		_, err = strconv.ParseFloat(v, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}

		invalid = append(invalid, v)
	}

	fmt.Println("#read:", total, "#ints:", nInts, "#floats:", nFloats)
	if len(invalid) > total {
		fmt.Println("Too much invalid input:", len(invalid))
		for _, s := range invalid {
			fmt.Println(s)
		}
	}
}
