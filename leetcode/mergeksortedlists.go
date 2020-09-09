package main

import (
	"fmt"
)

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
	if lists == nil || len(lists) == 0 {
		return &ListNode{}
	}
	// start at the second list
	for i := 1; i <= len(lists)-1; i++ {
		for {
			headAtZero, headAtI := lists[0], lists[i]
			// if the head at i is nil break
			if headAtI == nil {
				break
			}
			if headAtZero == nil {
				return headAtI
			}

			// compare heads values and swap if first is greater
			if headAtZero.Val >= headAtI.Val {
				lists[i] = headAtI.Next
				headAtI.Next = headAtZero
				lists[0] = headAtI
			} else {
				// traverse the first list
				for headAtZero.Next != nil {
					if headAtZero.Next.Val > headAtI.Val {
						lists[i] = headAtI.Next
						headAtI.Next = headAtZero.Next
						headAtZero.Next = headAtI
						break
					}
					headAtZero = headAtZero.Next

					if headAtZero.Next == nil {
						lists[i] = headAtI.Next
						headAtI.Next = nil
						headAtZero.Next = headAtI
						headAtZero.Next.Next = nil
						break
					}
				}
			}
		}
	}
	return lists[0]
}

// Print traverse the linked list and print each element
func (node *ListNode) Print() {
	for ; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}

func main() {
	// y := &ListNode{Val: 8}
	// x := &ListNode{Val: 7, Next: y}
	//
	// b := &ListNode{Val: 4, Next: x}
	//
	// a := &ListNode{Val: 3, Next: b}
	//
	// d := &ListNode{Val: 2}
	// c := &ListNode{Val: 1, Next: d}
	//
	// f := &ListNode{Val: 6}
	// e := &ListNode{Val: 5, Next: f}
	// arr := []*ListNode{a, c, e}
	// val := mergeKLists(arr)
	// val.Print()
	//
	// arr = []*ListNode{}
	// val = mergeKLists(arr)
	// val.Print()
	//
	// f = &ListNode{Val: 6}
	// arr = []*ListNode{nil, f}
	// val = mergeKLists(arr)
	// val.Print()
	//TODO case not working
	f := &ListNode{Val: 1}
	g := &ListNode{Val: 0}
	arr := []*ListNode{g, f}
	val := mergeKLists(arr)
	val.Print()
}
