package main

import (
	"fmt"
)

func main() {
	var a,b,c string = "Bring","It","In"

	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c)
	
	// Or with new lines...
	fmt.Print(" ", a, "\n")
	fmt.Print(b, "\n")
	fmt.Print(c, "\n")

	// Or multi variables in one Print()
	fmt.Print(a, "\n", b, " ",c)

	var d,e = 99,88
	fmt.Print("If the arguments are not strings such as an int/float32 \n aka numbers, Print() auto-adds a space like this...", d, e)

	// Try with Println()
	fmt.Println(a,b,c)

	// Testing out Printf()
	var f string = "Spectactular"
	var g float32 = 89.99

	fmt.Printf("d has value: %v and type: %T\n", f, f)
	fmt.Printf("e has value: %v and type: %T", g, g)
}