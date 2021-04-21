package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(items []int) []int {
	l := len(items)
	if l == 1 || l == 0 {
		return items
	}

	mid := l / 2
	left := make([]int, mid)
	right := make([]int, l-mid)

	for i := 0; i < l; i++ {
		if i < mid {
			left[i] = items[i]
		} else {
			right[i-mid] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

// Generates a slice of size filled with random numbers
func randomSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func main() {
	slice := randomSlice(20)
	fmt.Println("--- unsorted --- \n", slice)
	fmt.Println("--- sorted ---\n", mergeSort(slice))
}
