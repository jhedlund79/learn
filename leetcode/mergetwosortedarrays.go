package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1, nums2 = nums1[:m], nums2[:n]
	nums1 = append(nums1, nums2...)
	for i := 1; i < len(nums1); i++ {
		for j := 0; j < i; j++ {
			if nums1[j] > nums1[i] {
				nums1[j], nums1[i] = nums1[i], nums1[j]
			}
		}
	}
	fmt.Println(nums1)
}

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	merge(a, 3, b, 3)
}
