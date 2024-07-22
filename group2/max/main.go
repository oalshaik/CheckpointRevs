package main

import (
	"fmt"
)

// Max returns the maximum value in a slice of integers
func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	maxValue := a[0]          // starting max value at the first element of the slice
	for _, value := range a { // iterate over the slice
		if value > maxValue { // if the current value is greater than the current max value
			maxValue = value // replace max value with current value
		}
	}
	return maxValue
}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max) // Should print 123
}
