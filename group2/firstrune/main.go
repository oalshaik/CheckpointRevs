package main

import (
	"github.com/01-edu/z01"
)

// FirstRune returns the first rune of a string
func FirstRune(s string) rune {
	runes := []rune(s) // Convert the string to a slice of runes
	if len(runes) == 0 {
		return 0 // Return 0 if the string is empty
	}
	return runes[0] // Return the first rune
}

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}
