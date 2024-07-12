// Swap the values of a and b. The swap function only accepts pointers and doesn’t return anything
package main

import "fmt"

func main() {
  a, b := 5, 10
  // call swap here
  swap(&a, &b) // & calls the address of a (the pointer)
  fmt.Println(a == 10, b == 5)
}
func swap(a *int, b *int) { // * says the argument is a pointer
  // swap the values here
  temp := *a;
  *a = *b; // *a is the value pointed by a
  *b = temp;
}