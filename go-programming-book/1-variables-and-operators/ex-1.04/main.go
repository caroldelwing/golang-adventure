package main

import (
	"fmt"
	"time"
)

// Skip the optional initial values or type declarations from the variable declaration

var (
	Debug bool // False is a zero value, so it is set by default and you don't need to set it
	LogLevel = "info" // You don't need to declare the type
	startUpTime = time.Now()
)

func main () {
	fmt.Println(Debug, LogLevel, startUpTime)
}