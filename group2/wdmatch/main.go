package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	if wdmatch(s1, s2) {
		printStr(s1)
		z01.PrintRune('\n')
	}
}

// wdmatch checks if it is possible to write s1 using characters from s2
// while respecting the order of characters in s2
func wdmatch(s1, s2 string) bool {
	j := 0
	for i := 0; i < len(s2) && j < len(s1); i++ {
		if s2[i] == s1[j] {
			j++
		}
	}
	return j == len(s1)
}

// printStr prints a string using z01.PrintRune
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
