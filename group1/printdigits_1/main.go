package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// Define the range of digits from '0' to '9'
	for r := '0'; r <= '9'; r++ {
		// Print each digit using z01.PrintRune
		z01.PrintRune(r)
	}
	// Print a newline character at the end
	z01.PrintRune('\n')
}
