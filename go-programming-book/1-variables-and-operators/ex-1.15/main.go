package main

import "fmt"

// Function design with pointers

func add5Value (count int) {
	count += 5
	fmt.Println("add5Value :", count)
}

func add5Point (count *int) {
	*count += 5
	// Print out the updated value of count and dereference it
	fmt.Println("add5Point :", *count)
}

func main() {
	var count int
	add5Value(count)
	fmt.Println("add5Value post:", count)

	add5Point(&count) // Use & to pass a pointer ti the variable
	fmt.Println("add5Point post:", count)
}

// When passing by value, the changes you make to the value in a function do not affect the value of the variable that’s passed to the function, while passing a pointer to a value does change the value of the variable passed to the function.