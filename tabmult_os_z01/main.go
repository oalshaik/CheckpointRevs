package main

import (
	"os"

	"github.com/01-edu/z01"
)

// printStr prints a string using z01.PrintRune
func printStr(s string) {
	// Iterate over each rune in the string and print it
	for _, r := range s {
		z01.PrintRune(r)
	}
}

// printInt converts an integer to a string and prints it using z01.PrintRune
func printInt(n int) {
	// Handle the special case where the number is 0
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := []rune{} // Initialize an empty slice to store the digits
	// Extract each digit from the number and store it in reverse order
	for n > 0 {
		// Prepend the current digit to the digits slice
		digits = append([]rune{rune('0' + n%10)}, digits...)
		n /= 10 // Remove the last digit from the number
	}

	// Print each digit in the correct order
	for _, d := range digits {
		z01.PrintRune(d)
	}
}

func main() {
	// Check if the number of arguments is exactly 1 (excluding the program name)
	if len(os.Args) != 2 {
		printStr("Usage: go run . <number>\n")
		return
	}

	arg := os.Args[1] // Get the argument (the input number as a string)
	number := 0
	// Convert the string argument to an integer
	for _, r := range arg {
		if r < '0' || r > '9' { // Check if the character is not a digit
			printStr("Please provide a valid positive integer.\n")
			return
		}
		// Update the number by shifting its digits left and adding the new digit
		number = number*10 + int(r-'0')
	}

	// Check if the number is a positive integer
	if number <= 0 {
		printStr("Please provide a valid positive integer.\n")
		return
	}

	// Print the multiplication table for the number from 1 to 9
	for i := 1; i <= 9; i++ {
		printInt(i)          // Print the multiplier
		printStr(" x ")      // Print the " x " string
		printInt(number)     // Print the input number
		printStr(" = ")      // Print the " = " string
		printInt(i * number) // Print the product
		z01.PrintRune('\n')  // Print a newline character
	}
}
