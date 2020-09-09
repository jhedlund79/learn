package main

//
import "fmt"

// Given two sorted arrays nums1 and nums2 of size m and n respectively.
//
// Return the median of the two sorted arrays.
//
// Follow up: The overall run time complexity should be O(log (m+n)).
//
//
//
// Example 1:
//
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:
//
// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
// Example 3:
//
// Input: nums1 = [0,0], nums2 = [0,0]
// Output: 0.00000
// Example 4:
//
// Input: nums1 = [], nums2 = [1]
// Output: 1.00000
// Example 5:
//
// Input: nums1 = [2], nums2 = []
// Output: 2.00000
//
//
// Constraints:
//
// nums1,length == m
// nums2,length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
func sortSlices(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
	return nums
}

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sorted := sortSlices(merged)
	if len(merged) == 1 {
		return float64(sorted[0])
	}

	a := len(merged) / 2
	if len(merged)%2 != 0 {
		return float64(merged[a])
	}
	return float64(merged[a-1]+merged[a]) / 2.0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 5, 6}, []int{4, 7, 8, 9}))
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 5, 6}, []int{4, 7, 8, 9, 10}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{2}))
	fmt.Println(findMedianSortedArrays(nil, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{3}, []int{-2, -1}))
}
