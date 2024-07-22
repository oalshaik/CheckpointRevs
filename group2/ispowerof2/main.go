package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if there is exactly one argument (excluding the program name)
	if len(os.Args) != 2 {
		return
	}

	// Convert the argument to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}

	// Check if the number is a power of 2
	if isPowerOf2(n) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

// isPowerOf2 determines if a number is a power of 2
func isPowerOf2(n int) bool {
	// A number is a power of 2 if it is greater than 0
	// and has exactly one bit set in its binary representation
	return n > 0 && (n&(n-1)) == 0
}
