package main

import (
	"fmt"
	"time"
)

// Implement a short variable declaration. It is only valid INSIDE the function.
// It requires an initial value

func main () {

	Debug := false
	LogLevel := "info"
	startUpTime := time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)
}