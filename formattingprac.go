package main

import (
	"fmt"
)

func main() {
	
	// These are the general formatting verbs.
	var i = 69.99
	var txt = "This is awesome!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	// These are how you format integers.
	var p = 88888888

	fmt.Printf("%b\n", p)
	fmt.Printf("%d\n", p)
	fmt.Printf("%+d\n", p)
	fmt.Printf("%o\n", p)
	fmt.Printf("%O and also this\n", p)
	fmt.Printf("%x\n", p)
	fmt.Printf("%X\n", p)
	fmt.Printf("%#x\n", p)
	fmt.Printf("Pad %4d pad\n", p)
	fmt.Printf("Pad %-4d pad\n", p)
	fmt.Printf("Pad %04d pad\n", p)

	// Now let's format with strings.
	var words = "SpecialtySauce"

	fmt.Printf("%s\n", words)
	fmt.Printf("%q\n", words)
	fmt.Printf("%8s\n", words)
	fmt.Printf("%-8s\n", words)
	fmt.Printf("%x\n", words)
	fmt.Printf("% x\n", words)

	// Boolean time
	var j = true
	var k = false

	fmt.Printf("%t\n", j)
	fmt.Printf("%t\n", k)

	// It's time to float!
	var f = 0.99

	fmt.Printf("%e\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%.20f\n", f) // %.xf where x is how many digits out you want.
	fmt.Printf("%6.2f\n", f)
	fmt.Printf("%g\n", f)

}