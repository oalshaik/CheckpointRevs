package main

import (
	"github.com/01-edu/z01"
)

// Rot14 applies the ROT14 cipher to a string
func Rot14(s string) string {
	var result []rune

	for _, r := range s {
		switch {
		case 'a' <= r && r <= 'z':
			// For lowercase letters
			result = append(result, ((r-'a'+14)%26)+'a')
		case 'A' <= r && r <= 'Z':
			// For uppercase letters
			result = append(result, ((r-'A'+14)%26)+'A')
		default:
			// For non-alphabetic characters
			result = append(result, r)
		}
	}

	return string(result)
}

func main() {
	result := Rot14("hiiIIIiih")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
