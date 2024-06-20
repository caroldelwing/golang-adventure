package main

// Create a medical form for a doctor’s office to capture a patient’s name, age, and whether they have a peanut allergy

import (
	"fmt"
)

func main () {
	var name string = "Carol"
	var family_name string = "Delwing"
	var age int = 27
	var is_peanut bool = false
	// Use printf when you need to format the data
	fmt.Printf("%v\n%v\n%v\n%v\n", name, family_name, age, is_peanut)
}