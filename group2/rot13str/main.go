package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	output := rot13(input)
	fmt.Println(output)
}

// rot13 applies the ROT13 cipher to a string
func rot13(s string) string {
	var result strings.Builder

	for _, r := range s {
		switch {
		case 'a' <= r && r <= 'z':
			result.WriteRune(((r - 'a' + 13) % 26) + 'a')
		case 'A' <= r && r <= 'Z':
			result.WriteRune(((r - 'A' + 13) % 26) + 'A')
		default:
			result.WriteRune(r)
		}
	}

	return result.String()
}
