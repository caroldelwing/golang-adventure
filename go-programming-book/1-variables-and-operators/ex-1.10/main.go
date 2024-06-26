package main

import "fmt"

// Shorthand operators

func main() {

	count := 5
	count += 5 // Adds 5 to count
	fmt.Println(count)

	count++ // Increment by 1
	fmt.Println(count)

	count--
	fmt.Println(count)

	count -= 5
	fmt.Println(count)

	// String shorthand
	name := "Bob"
	name += " Marley" // Nice :)
	fmt.Println("Hello, ", name)

}