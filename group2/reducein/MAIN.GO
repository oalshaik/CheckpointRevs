package main

import "fmt"

// ReduceInt applies the function f to each element of the slice a
func ReduceInt(a []int, f func(int, int) int) {
	if len(a) == 0 {
		return
	}

	result := a[0]
	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}

	fmt.Println(result)
}

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul) // 500 * 2 = 1000
	ReduceInt(as, sum) // 500 + 2 = 502
	ReduceInt(as, div) // 500 / 2 = 250
}
