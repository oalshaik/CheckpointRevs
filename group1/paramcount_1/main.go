package main

import (
	"os"

	"github.com/01-edu/z01"
)

// printInt prints an integer using z01.PrintRune
func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		digits := []rune{}
		for n > 0 {
			digits = append([]rune{rune('0' + n%10)}, digits...)
			n /= 10
		}
		for _, d := range digits {
			z01.PrintRune(d)
		}
	}
	z01.PrintRune('\n')
}

func main() {
	// The number of arguments is the length of os.Args minus 1 (because os.Args[0] is the program name)
	argCount := len(os.Args) - 1
	printInt(argCount)
}
