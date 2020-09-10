package main

import "C"
import "fmt"

// 23. Merge k Sorted Lists
// Hard
//
// 5299
//
// 303
//
// Add to List
//
// Share
// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
//
// Merge all the linked-lists into one sorted linked-list and return it.
//
//
//
// Example 1:
//
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
//  1->4->5,
//  1->3->4,
//  2->6
// ]
// merging them into one sorted list:
// 1->1->2->3->4->4->5->6
// Example 2:
//
// Input: lists = []
// Output: []
// Example 3:
//
// Input: lists = [[]]
// Output: []

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	vals := []int{}
	// grab all the values out of the nodes
	for _, list := range lists {
		if list == nil {
			continue
		}
		curr := list
		for curr.Next != nil {
			vals = append(vals, curr.Val)
			curr = curr.Next
		}
		if curr.Next == nil {
			vals = append(vals, curr.Val)
		}
	}

	// sort the values
	for i := 1; i < len(vals); i++ {
		for j := 0; j < i; j++ {
			if vals[j] > vals[i] {
				vals[j], vals[i] = vals[i], vals[j]
			}
		}
	}

	// construct new nodes in a linked list
	var head *ListNode
	if len(vals) > 0 {
		head = &ListNode{Val: vals[0]}
		curr := head
		for i := 1; i < len(vals); i++ {
			curr.Next = &ListNode{Val: vals[i]}
			curr = curr.Next
		}
	}
	return head
}

// Print traverse the linked list and print each element
func (node *ListNode) Print() {
	for ; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
	fmt.Println()
}

func main() {
	y := &ListNode{Val: 8}
	x := &ListNode{Val: 7, Next: y}

	b := &ListNode{Val: 4, Next: x}

	a := &ListNode{Val: 3, Next: b}

	d := &ListNode{Val: 2}
	c := &ListNode{Val: 1, Next: d}

	f := &ListNode{Val: 6}
	e := &ListNode{Val: 5, Next: f}
	arr := []*ListNode{a, c, e}
	val := mergeKLists(arr)
	val.Print()

	fmt.Println("empty list")
	arr = []*ListNode{}
	val = mergeKLists(arr)
	val.Print()

	fmt.Println("nil list")
	f = &ListNode{Val: 6}
	arr = []*ListNode{nil, f}
	val = mergeKLists(arr)
	val.Print()

	fmt.Println("one empty list")
	f = &ListNode{Val: 1}
	arr = []*ListNode{&ListNode{}, f}

	val = mergeKLists(arr)
	val.Print()

	fmt.Println("list with zero value")
	f = &ListNode{Val: 1}
	g := &ListNode{Val: 0}
	arr = []*ListNode{g, f}
	val = mergeKLists(arr)
	val.Print()
}
