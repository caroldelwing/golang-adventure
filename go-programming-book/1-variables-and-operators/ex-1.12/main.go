package main

import (
	"fmt"
	"time"
)

// Define some variables without an initial value and check its zero value

func main() {

	var count int
	fmt.Printf("Count : %#v \n", count)


	var discount float64
	fmt.Printf("Discount : %#v \n", discount)

	var debug bool
	fmt.Printf("Debug : %#v \n", debug)

	var message string
	fmt.Printf("Message : %#v \n", message)

	var emails []string // Colection of strings
	fmt.Printf("Emails : %#v \n", emails)

	var startTime time.Time // Struct
	fmt.Printf("Start : %#v \n", startTime)
}