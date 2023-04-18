// Declare the package.
package main

// Import the package.
import (
	"fmt"
)

// Using var allows putting variables outside of func main()
var d int
var e int = 2
var f = 3

func main() {
	var student1 string = "John" // string type variable
	var student2 = "Jane"        // inferred type

	x := 2 // also inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	// Assign variable as a string and establish value after
	var a string
	a = "Sizzle"

	// Assign variables types which will return their default values
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Assigned this as int earlier before func main() and assigning value now
	d = 79

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	// Easily assign multiple variables in a single block.
	var g, h, i, j int = 99, 54, 33, 77

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

	var k, l = 23, "Jordan"
	m, n := 55, "Matsui"

	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)

	var (
		o int
		p int    = 1
		q string = "that's a wrap!"
	)

	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(q)
}
