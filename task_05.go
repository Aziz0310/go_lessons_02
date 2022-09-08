package main

import "fmt"

func main() {
	n := 12
	// Read n from input
	DisplayMinimumNumberFunction(n)
}

// https://www.hackerrank.com/contests/w30/challenges/find-the-minimum-number
func DisplayMinimumNumberFunction(n int) {
	x := "min(int, int)"
	for i := 3; i <= n; i++ {
		x = fmt.Sprintf("min(int, %s)", x)
	}
	fmt.Println(x)
	//
	// WRITE YOUR CODE HERE
	//min := values[0]
	// for _, n := range values {
	// 	if n < min {
	// 		min = n
	// 	}
	// }

	// fmt.Println(min)
}
