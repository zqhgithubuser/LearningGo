package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MIN = 0
var MAX = 94

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myrand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myrand))
		temp += newChar
		if len == i {
			break
		}
		i++
	}
	return temp
}

func main() {
	var LENGTH int64 = 8

	seed := time.Now().Unix()
	rand.Seed(seed)
	fmt.Println(getString(LENGTH))
}
