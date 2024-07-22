package main

import (
	"github.com/01-edu/z01"
)

// LastRune returns the last rune of a string
func LastRune(s string) rune {
	runes := []rune(s) // Convert the string to a slice of runes
	if len(runes) == 0 {
		return 0 // Return 0 if the string is empty
	}
	return runes[len(runes)-1] // Return the last rune
}

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}
