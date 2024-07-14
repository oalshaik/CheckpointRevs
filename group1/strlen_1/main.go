package main

import (
	"fmt"
)

// in the test you will only write this one
// StrLen counts the number of runes in a string and returns the count
func StrLen(s string) int {
	c := 0
	for range s {
		c++
	}
	return c
}

func main() {
	// Test the StrLen function
	l := StrLen("Hello World!")
	fmt.Println(l) // Should print 12
}
