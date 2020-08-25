package main

import (
	"fmt"
	"strings"
)

// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
// Example 1:
//
// Input: ["flower","flow","flight"]
// Output: "fl"
// Example 2:
//
// Input: ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
// Note:
//
// All given inputs are in lowercase letters a-z.
func longestCommonPrefix(strs []string) string {
	if 0 == len(strs) {
		return ""
	}

	short := strs[0]
	for _, s := range strs {
		if len(short) >= len(s) {
			short = s
		}
	}

	prefixArr := []string{}
	prefix := ""
	oldPrefix := ""
	for i := 0; i < len(short); i++ {
		prefixArr = append(prefixArr, string(short[i]))
		prefix = strings.Join(prefixArr, "")

		for _, s := range strs {
			if !strings.HasPrefix(s, prefix) {
				return oldPrefix
			}
		}
		oldPrefix = prefix

	}
	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
