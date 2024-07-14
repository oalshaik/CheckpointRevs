package main

import (
	"os"

	"github.com/01-edu/z01"
)

// printStr prints a string using z01.PrintRune
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	// Check if there are more than one argument (os.Args[0] is the program name)
	if len(os.Args) > 1 {
		// Retrieve the last argument
		lastArg := os.Args[len(os.Args)-1]
		// Print the last argument
		printStr(lastArg)
	}
}
