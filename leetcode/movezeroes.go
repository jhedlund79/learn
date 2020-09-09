package main

import "fmt"

func moveZeroes(nums []int)  {
	for _, _ = range nums{
		for j, num := range nums {
			if j+1 == len(nums) {
				break
			}
			if num == 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}