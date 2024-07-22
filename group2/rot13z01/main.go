package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	rot13(input)
}

// rot13 applies the ROT13 cipher to a string and prints the result
func rot13(s string) {
	for _, r := range s {
		switch {
		case 'a' <= r && r <= 'z':
			// If the rune is a lowercase letter, apply ROT13 and print it
			z01.PrintRune(((r - 'a' + 13) % 26) + 'a')
		case 'A' <= r && r <= 'Z':
			// If the rune is an uppercase letter, apply ROT13 and print it
			z01.PrintRune(((r - 'A' + 13) % 26) + 'A')
		default:
			// If the rune is not a letter, print it unchanged
			z01.PrintRune(r)
		}
	}
	// Print a newline character at the end
	z01.PrintRune('\n')
}
