package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

// printStr prints a string using z01.PrintRune
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

// printInt prints an integer using z01.PrintRune
func printInt(n int) {
	printStr(strconv.Itoa(n))
}

func main() {
	if len(os.Args) != 2 {
		printStr("Usage: go run . <number>\n")
		return
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil || number <= 0 {
		printStr("Please provide a valid positive integer.\n")
		return
	}

	for i := 1; i <= 9; i++ {
		printInt(i)
		printStr(" x ")
		printInt(number)
		printStr(" = ")
		printInt(i * number)
		z01.PrintRune('\n')
	}
}
