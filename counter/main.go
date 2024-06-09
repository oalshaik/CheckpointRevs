package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:] // os.Args[0] is the program name, so we skip it
	count := 0

	for _, arg := range args {
		// Try to convert the argument to an integer
		if _, err := strconv.Atoi(arg); err == nil {
			count++
		}
	}

	fmt.Println(count)
}
