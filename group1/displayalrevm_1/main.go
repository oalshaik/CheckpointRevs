package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for i := 'z'; i >= 'a'; i-- {
		if (i-'a')%2 == 0 {
			z01.PrintRune(i - 'a' + 'A') // Even letters in uppercase
		} else {
			z01.PrintRune(i) // Odd letters in lowercase
		}
	}
	z01.PrintRune('\n')
}
