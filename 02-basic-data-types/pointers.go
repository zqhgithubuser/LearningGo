package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)
	// 指向f的指针
	fP := &f
	fmt.Println("Memory address of f:", fP)
	fmt.Println("Value of f:", *fP) // 12.123

	processPointer(fP)
	fmt.Printf("Value of f: %.2f\n", f) // 146.97

	x := returnPointer(f)
	fmt.Printf("Value of f: %.2f\n", f)  // 146.97
	fmt.Printf("Value of x: %.2f\n", *x) // 293.93

	xx := bothPointers(fP)
	fmt.Printf("Value of f: %.2f\n", f)
	fmt.Printf("Value of x: %.2f\n", *xx) // 293.93

    var k *aStructure
    fmt.Println(k) // <nil>
    if k == nil {
        k = new(aStructure)
    }
    fmt.Printf("%+v\n", k)
    if k != nil {
        fmt.Println("k is not nil!")
    }
}
