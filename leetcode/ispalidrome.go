package main

import (
	"fmt"
	"strings"
)

/*

Given a string s, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.


Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

func isPalindrome(s string) bool {
	s = stripNonAlpha(strings.ToLower(s))
	if s == reverseStr([]byte(s)) {
		return true
	}
	return false
}

func stripNonAlpha(s string) string {
	acceptedChars := "abcdefghijklmnopqrstuvwxyz1234567890"
	var res string
	for _, c := range s {
		if strings.Contains(acceptedChars, string(c)) {
			res += string(c)
		}
	}
	return res
}

func reverseStr(s []byte) string {
	l := len(s)

	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
	return string(s)
}

func main() {
	fmt.Println(isPalindrome("racecar"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
