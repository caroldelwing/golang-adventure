package main

import (
	"fmt"
)

var foo string = "bar" // Package level

func main () {
	var baz string = "qux" // Function level
	fmt.Println(foo, baz)
}