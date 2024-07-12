package main

import (
	"fmt"
	"time"
)

// Get a value from a pointer

func main() {

	var count1 *int // Declare a pointer. It will have a nil value
	
	count2 := new(int) // Create a pointer variable using new, which gets some memory for a type and returns a pointer to that address

	countTemp := 5 // You cannot take the address of a literal number

	count3 := &countTemp // Create a pointer from the variable

	t := &time.Time{} // Create a pointer from time struct directly without a temp variable

	// Add nil checks and use dereferencing to get the values of the pointers
	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}
	if t != nil {
		fmt.Printf("time: %#v\n", *t)
	}

	fmt.Printf("time: %#v\n", t.String())

}