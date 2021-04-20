package main

import "fmt"

// Fibonacci seq 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55…
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-2) + fib(n-1)
}

func fibIter(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	prevprevNum := 0
	previousNum := 1
	currentNum := 1
	for i := 2; i < n; i++ {
		prevprevNum = previousNum
		previousNum = currentNum
		currentNum = prevprevNum + previousNum
	}
	return currentNum
}

func main() {
	fmt.Println(fib(8))
	fmt.Println(fibIter(8))
}
