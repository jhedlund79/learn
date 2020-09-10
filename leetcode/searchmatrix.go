package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	// i number of rows
	i := len(matrix) - 1
	// n is the number of cols
	n := len(matrix[0])
	j := 0
	for i >= 0 && j < n {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			//change rows
			i--
		} else {
			// continue scanning row
			j++
		}
	}
	return false
}

func main() {
	m := [][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}}
	fmt.Println(searchMatrix(m, 23))

	m2 := [][]int{[]int{}}
	fmt.Println(searchMatrix(m2, 0))

	m3 := [][]int{[]int{1}}
	fmt.Println(searchMatrix(m3, 1))
}
