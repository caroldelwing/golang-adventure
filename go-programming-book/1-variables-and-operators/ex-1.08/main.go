package main

import (
	"fmt"
)

// Change the value of multiple variables at once


func main() {
	query, limit, offset := "bat", 10, 0 // Declare the variables
	query, limit, offset = "ball", offset, 20 // Change the values
	fmt.Println(query, limit, offset)
}
