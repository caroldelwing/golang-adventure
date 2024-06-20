package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Print a random number, between 1 and 5, of stars (*) to the console
func main () {

	rand.Seed(time.Now().UnixNano()) // It seeds the random number generator
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)
}