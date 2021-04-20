package main

import "fmt"

//A non-empty array A consisting of N integers is given. The array contains an odd number of elements, and each element of the array can be paired with another element that has the same value, except for one element that is left unpaired.
//
//For example, in array A such that:
//
//  A[0] = 9  A[1] = 3  A[2] = 9
//  A[3] = 3  A[4] = 9  A[5] = 7
//  A[6] = 9
//the elements at indexes 0 and 2 have value 9,
//the elements at indexes 1 and 3 have value 3,
//the elements at indexes 4 and 6 have value 9,
//the element at index 5 has value 7 and is unpaired.
//Write a function:
//
//func Solution(A []int) int
//that, given an array A consisting of N integers fulfilling the above conditions, returns the value of the unpaired element.
//
//For example, given array A such that:
//
//  A[0] = 9  A[1] = 3  A[2] = 9
//  A[3] = 3  A[4] = 9  A[5] = 7
//  A[6] = 9
//the function should return 7, as explained in the example above.
//
//Write an efficient algorithm for the following assumptions:
//
//N is an odd integer within the range [1..1,000,000];
//each element of array A is an integer within the range [1..1,000,000,000];
//all but one of the values in A occur an even number of times.perm
func findTheSingle(arr []int) int {
	m := map[int]int{}
	for _, num := range arr {
		if _, ok := m[num]; !ok {
			m[num] = 0
		}
		m[num] += 1
	}
	for k, v := range m {
		fmt.Println("k, v: ", k, v)
		if v == 1 {
			return k
		}
	}
	return 0
}
func main() {
	n := []int{}
	for i := 0; i < 201; i++ {
		if i == 56 {
			continue
		}
		n = append(n, i)
	}
	for i := 0; i < 201; i++ {
		n = append(n, i)
	}
	fmt.Println(findTheSingle(n))

	n1 := []int{}
	for i := 0; i < 10000000; i++ {
		if i == 56 {
			continue
		}
		n1 = append(n1, i)
	}
	for i := 0; i < 10000000; i++ {
		n1 = append(n1, i)
	}
	fmt.Println(findTheSingle(n1))
	fmt.Println(findTheSingle([]int{1, 2, 1, 3, 4, 4, 3, 5, 5, 6, 7, 7, 8, 8, 6, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}))
	fmt.Println(findTheSingle([]int{9, 3, 9, 3, 9, 7, 9}))
}
