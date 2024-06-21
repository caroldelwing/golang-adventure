package main

import (
	"fmt"
	"time"
)

// Call a function that returns multiple values, and assign each value to a new variable

func getConfig() (bool, string, time.Time) { // The function returns 3 values
	return false, "info", time.Now()
} 

func main() {
	Debug, LogLevel, startUpTime := getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
}
