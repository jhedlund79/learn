package main

import (
	"fmt"
	"strconv"
)

//Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
//Notice that the solution set must not contain duplicate triplets.
//
//
//
//Example 1:
//
//Input: nums = [-1,0,1,2,-1,-4]
//Output: [[-1,-1,2],[-1,0,1]]
//Example 2:
//
//Input: nums = []
//Output: []
//Example 3:
//
//Input: nums = [0]
//Output: []
//
//
//Constraints:
//
//0 <= nums.length <= 3000
//-105 <= nums[i] <= 105
func sortInts(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}

	}
}
func threeSum(nums []int) [][]int {
	sortInts(nums)
	res := [][]int{}
	m := map[string]int{}
	for i := 0; i < len(nums)-1; i++ {
		left := i + 1
		right := len(nums) - 1
		for right > left {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				rs := strconv.Itoa(nums[i]) + strconv.Itoa(nums[left]) + strconv.Itoa(nums[right])
				if _, ok := m[rs]; !ok {
					res = append(res, []int{nums[i], nums[left], nums[right]})
					m[rs] = 1
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}

		}
	}

	return res
}
func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
