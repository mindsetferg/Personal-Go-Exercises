package main

import "fmt"

func main() {
	// Plays with slices.

	// With no [...] in the box, this can be a slice.
	pizza_slice := []string{"Pepperoni", "Cheese", "Ham", "Sausage"}

	// Also, better to initialize with an {values} section.
	pizza_default := []int{}
	
	// Initialize it by adding values.
	pizza_sections := []int{4,8,12}

	fmt.Print(len(pizza_slice),"\n",cap(pizza_slice),"\n",pizza_slice,"\n")
	/* Note: if you use Println() for the above, it will print with weird spaces.
	Print() requires more manipulation than Prentln()*/
	//fmt.Println(cap(pizza_slice))
	//fmt.Println(pizza_slice)

	fmt.Print(len(pizza_default),"\n",cap(pizza_default),"\n",pizza_default,"\n")
	//fmt.Println(cap(pizza_default))
	//fmt.Println(pizza_default)

	fmt.Println(len(pizza_sections))
	fmt.Println(cap(pizza_sections))
	fmt.Println(pizza_sections)

	pizza_prices := [8]float32{10.99, 11.99, 12.99, 13.99, 20.99, 21.99, 22.99, 23.99}
	my_pizza_price := pizza_prices[0:3]
	fmt.Printf("my_pizza_price = %v\n", my_pizza_price)
	fmt.Printf("prices = %d\n", len(my_pizza_price))
	fmt.Printf("different prices = %d\n", cap(my_pizza_price))
	
	// We can also append here.
	my_pizza_price = pizza_prices[2:6]
	fmt.Printf("my_pizza_price = %v\n", my_pizza_price)
	fmt.Printf("length = %d\n", len(my_pizza_price))
	fmt.Printf("capacity = %d\n", cap(my_pizza_price))

	// Now append to my_pizza_price
	my_pizza_price = append(my_pizza_price, 31.99, 23.95, 27.95)
	fmt.Printf("my_pizza_price = %v\n", my_pizza_price)
	fmt.Printf("my_pizza_price = %d\n", len(my_pizza_price))
	fmt.Printf("capacity = %d\n", cap(my_pizza_price))
	fmt.Println(pizza_prices)

	/* Know the arrays and slices formats:
	var myarray = [length]datatype{values}
	aslice := []datatype{values}
	with make() function:
	slicename := make([]type, length, capacity)*/

	potential_appetizers := make([]int,5,10) // When setting a length, it will add 0's.
	fmt.Printf("potential_appetizers = %v\n", potential_appetizers)
	fmt.Printf("length = %d\n", len(potential_appetizers))
	fmt.Printf("capacity = %d\n", cap(potential_appetizers))

	//Now if I don't add a capacity:
	potential_apps := make ([]int,5)
	fmt.Printf("potential_apps = %v\n", potential_apps)
	fmt.Printf("length = %d\n", len(potential_apps))
	fmt.Printf("capacity = %d\n", cap(potential_apps))

	//We can now start accessing some values:

	// Easy access format: (accessing pizza_prices)
	fmt.Println(pizza_prices[4])
	fmt.Println(pizza_prices[7])

	// Easy to change:
	pizza_slice[3] = "Veggie"
	fmt.Println(pizza_slice[0])
	fmt.Println(pizza_slice[3])

	// Easy to add/append (at the end):
	pizza_slice = append(pizza_slice, "Sausage", "Mushroom")
	fmt.Printf("pizza_slice = %v\n", pizza_slice)
	fmt.Printf("length = %d\n", len(pizza_slice))
	fmt.Printf("capacity = %d\n", cap(pizza_slice))

	// Append a slice to a slice: note that you can't append two different types of slices possibly
	pizza_and_sections := append(pizza_default, pizza_sections...)
	fmt.Printf("pizza_and_sections=%v\n", pizza_and_sections)
	fmt.Printf("length=%d\n", len(pizza_and_sections))
	fmt.Printf("capacity=%d\n", cap(pizza_and_sections))

	// Practice with memory.
	customers := []int{0,5,10,15,20,25,30,35,40,45,50,55,60,65,70,75,80,85,90,95,100}
	fmt.Printf("customers = %v\n", customers)
	fmt.Printf("length = %d\n", len(customers))
	fmt.Printf("capacity = %d\n", cap(customers))

	//Now just use the needed numbers.
	currentCustomers := customers[:len(customers)-3]
	customersCopy := make([]int, len(currentCustomers))
	copy(customersCopy, currentCustomers)

	fmt.Printf("customersCopy = %v\n", customersCopy)
	fmt.Printf("length = %d\n", len(customersCopy))
	fmt.Printf("capacity = %d\n", cap(customersCopy))
}