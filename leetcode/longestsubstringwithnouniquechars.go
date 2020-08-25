package main

import (
	"fmt"
	"math"
)

// Given a string, find the length of the longest substring without repeating characters.
//
// Example 1:
//
// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:
//
// Input: "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:
//
// Input: "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
//             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
func lengthOfLongestSubstring(s string) int {
	index := [128]int{}
	var i, j, ans int
	for j = 0; j < len(s); j++ {
		fmt.Println("comparing for i: ", float64(index[s[j]]), ":", float64(i))
		i = int(math.Max(float64(index[s[j]]), float64(i)))
		fmt.Println("comparing for ans: ", float64(ans), ":", float64(j-i+1))
		ans = int(math.Max(float64(ans), float64(j-i+1)))
		index[s[j]] = j + 1
	}
	return ans
}
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
