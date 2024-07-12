package main

import (
	"fmt"
	"time"
)

// Create and print a pointer

func main() {

	var count1 *int // Declare a pointer. It will have a nil value
	
	count2 := new(int) // Create a pointer variable using new, which gets some memory for a type and returns a pointer to that address

	countTemp := 5 // You cannot take the address of a literal number

	count3 := &countTemp // Create a pointer from the variable

	t := &time.Time{} // Create a pointer from time struct directly without a temp variable

	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time: %#v\n", t)

}