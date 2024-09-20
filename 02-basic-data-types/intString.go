package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Print provide an integer.")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// strconv.Itoa()
	input := strconv.Itoa(n)
	fmt.Printf("string() %s of type %T\n", input, input)

	// strconv.FormatInt()
	input = strconv.FormatInt(int64(n), 10)
	fmt.Printf("string() %s of type %T\n", input, input)

	// string()
	input = string(n) // int -> rune
    fmt.Println(input)
	fmt.Printf("string() %s of type %T\n", input, input)
}
