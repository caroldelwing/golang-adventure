package main

import (
	"fmt"
	"time"
)

// Declare and print multiple variables using one var statement

var (
	Debug bool = false
	LogLevel string = "info"
	startUpTime time.Time = time.Now()
)

func main () {
	fmt.Println(Debug, LogLevel, startUpTime)
}