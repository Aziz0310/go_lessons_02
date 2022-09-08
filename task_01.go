package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		FizzBuzz(i)
	}
}

func FizzBuzz(i int) {
	if i%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if i%5 == 0 {
		fmt.Println("Buzz")
	} else if i%3 == 0 {
		fmt.Println("Fizz")
	} else {
		fmt.Println(i)
	}
}

// WRITE YOUR CODE HERE
//
