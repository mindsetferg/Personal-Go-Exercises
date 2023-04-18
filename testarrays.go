package main

import "fmt"

/* Let's declare an array in two ways: using var and :=
with var, it goes:
var array_name = [Length]datatype{values} with length defined
or
var array_name = [...]datatype{values} with length inferred

with :=

array_name := [Length]datatype{values} with length defined
or
array_name :\ [...]datatype{values} length inferred*/

func main() {
	var arrayone = [5]int{1,2,2} // If length is longer, it will add 0's.
	arraytwo := [17]float32{2.55, 3.78, 5.98, 3.22}

	fmt.Println(arrayone)
	fmt.Println(arraytwo)

	var arraythree = [...]int{11, 22, 33, 10} // Notice inferred vs declared differences with length
	arrayfour := [...]int{44, 55, 66, 77, 88} // The term is "initialized."

	fmt.Println(arraythree)
	fmt.Println(arrayfour)

	// We'll do one more with strings. Declaring also adds " " rather than 0.
	var cars = [10]string{"Pontiac", "Volkswagen", "Honda", "Porsche"}
	fmt.Print(cars, "\n")

	//Now let's access these.
	prices := [7]float32{29.99, 19.99, 23} //You can include integers in float32 but not int.

	fmt.Println(prices[0])
	fmt.Println(prices[5])

	// And now to change one of these values. (Notice the order of execution)
	prices[5] = 37
	fmt.Println(prices)

	// We can also update arrays from here.
	prices2 := [7]float32{0:19.99, 2:79.99}
	fmt.Println(prices2)

	// And of course we now need to find the lengths!

fmt.Println(len(arrayone))
fmt.Println(len(arraytwo))
fmt.Println(len(arraythree))
fmt.Println(len(arrayfour))
fmt.Println(len(cars))
fmt.Println(len(prices))
fmt.Println(len(prices2))

}