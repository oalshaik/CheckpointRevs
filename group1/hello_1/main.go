package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// Define the message to be printed
	message := "Hello World!"

	// Print each character in the message
	for _, r := range message {
		z01.PrintRune(r)
	}

	// Print a newline character
	z01.PrintRune('\n')
}
