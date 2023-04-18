package main

import (
	"fmt"
)

const PI = 3.14159

// This is a "Typed Constant" declared with a definite type
const A int = 1

// This is an "Untyped Constant" declared without a type.
const B = 33

// You can group constants!
const (
	C int = 3928
    D = 99.99
    E = "Let's do this!"
)
func main() {
	fmt.Println(PI)
	fmt.Println(A)
	fmt.Println(B)

	// This is for those grouped constants.

	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)

}