package main

import (
	"github.com/01-edu/z01"
)

func LastWord(s string) string {
	start := -1
	end := -1

	// Find the last word by iterating from the end of the string
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' && end == -1 {
			end = i
		}
		if s[i] == ' ' && end != -1 {
			start = i + 1
			break
		}
	}

	// If there is no word found
	if end == -1 {
		return "\n"
	}

	// If start is not set, it means the whole string until end is the last word
	if start == -1 {
		start = 0
	}

	// Extract the last word
	word := s[start : end+1]
	return word + "\n"
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	printString(LastWord("this        ...       is sparta, then again, maybe    not"))
	printString(LastWord(" lorem,ipsum "))
	printString(LastWord(" "))
}
