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
	if len(os.Args) > 1 {
		firstArg := os.Args[1]
		printStr(firstArg)
	}
}
