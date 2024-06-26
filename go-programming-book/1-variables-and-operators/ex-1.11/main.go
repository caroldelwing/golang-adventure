package main

import "fmt"

// Test to see what level of membership a user has based on the number of visits they’ve had

func main() {

	visits := 15
	fmt.Println("First visit :", visits == 1) // Check if this is the first visit
	fmt.Println("Return visit :", visits != 1) // Check if this is a returning visit

	fmt.Println("Silver member :", visits >= 10) // Check if this is a silver member
	fmt.Println("Gold member :", visits > 20 && visits <= 30) // Check if this is a gold member
	fmt.Println("Platinum member :", visits > 30) // Check if this is a platinum member
}