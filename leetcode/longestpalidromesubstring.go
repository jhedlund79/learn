package main

import (
	"fmt"
	"math"
)

//Given a string s, return the longest palindromic substring in s.
//
//
//
//Example 1:
//
//Input: s = "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.
//Example 2:
//
//Input: s = "cbbd"
//Output: "bb"
//Example 3:
//
//Input: s = "a"
//Output: "a"
//Example 4:
//
//Input: s = "ac"
//Output: "a"
//
//
//Constraints:
//
//1 <= s.length <= 1000
//s consist of only digits and English letters (lower-case and/or upper-case),

func longestPalidromeSubstr(s string) string {
	if s == "" || len(s) == 0 {
		return ""
	}
	var start, end int
	for i := range s {
		l1 := expandAroundCenter(s, i, i)
		l2 := expandAroundCenter(s, i, i+1)
		max := int(math.Max(float64(l1), float64(l2)))
		if max > end-start {
			start = i - (max-1)/2
			end = i + max/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func main() {
	fmt.Println(longestPalidromeSubstr("babad"))
	fmt.Println(longestPalidromeSubstr("cbbd"))
	fmt.Println(longestPalidromeSubstr("abacdfgdcabaabacdfgdcaba"))
}
