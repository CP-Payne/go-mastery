package main

import (
	"fmt"
)

// 1
var packageLevel int

// 8
const pi float32 = 3.14159

func main() {
	// 2
	var name string = "Charl"
	// 3
	var age int = 24
	// 4
	var height float64 = 80.0
	// 5
	isCodingFun := true
	// 6
	var city, country, population string
	// 7
	favouriteLanguage := "Golang"

	// 9
	x := 1
	y := 2
	x, y = y, x

	// 10
	fmt.Printf("Type: %T\n", x)

	// 11
	var x int
	fmt.Printf("Zero valued: %d\n", x)

	// 12
	var p *int
	p = &x
	*p = 999

	// 13 Continue
	a, b, c := ReturnMultipleValues()

	// 14
	someFloat := float64(a)
	someIntConverted := int(a)

	// 15
	const someInt = 33
	var diffTyped float64 = someInt

	// 17
	someIntValue := new(int)
	*someIntValue = 44

	// 18
	if someVal := 1; x == 33 {
		// Do something here, perhaps with someval...
	}
}

// 13
func ReturnMultipleValues() (int, int, int) {
	return 1, 2, 3
}

// 16
func ShadowedVar(x int) {
	iter := 5
	for x := 0; x <= iter; x++ {
		fmt.Println("x is being shadowed")
	}

	fmt.Println("x: ", x)
}

// 19
func OuterFunc() func() {
	counter := 0
	return func() {
		counter += 1
		fmt.Printf("Counter: %d", counter)
	}
}
