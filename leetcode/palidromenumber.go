package main

import "fmt"

// Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
//
// Example 1:
//
// Input: 121
// Output: true
// Example 2:
//
// Input: -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:
//
// Input: 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revNum := 0
	for x > revNum {
		revNum = revNum*10 + x%10
		x /= 10
	}
	return x == revNum || x == revNum/10
}

func main() {
	fmt.Println(isPalindrome(1001))
	fmt.Println(isPalindrome(2020))
}
