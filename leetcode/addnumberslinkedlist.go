package main

import (
	"github.com/davecgh/go-spew/spew"
)

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Example:
//
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry = 0
	curr := &ListNode{}
	arr := []*ListNode{}
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
			l2 = l2.Next
		}

		v := carry + x + y
		carry = v / 10
		curr.Next = &ListNode{Val: v % 10}
		curr = curr.Next
		if carry > 0 {
			curr.Next = &ListNode{Val: carry}
		}
		arr = append(arr, curr)
	}
	return arr[0]
}
func main() {
	c := &ListNode{Val: 3}
	b := &ListNode{Val: 4, Next: c}
	a := &ListNode{Val: 2, Next: b}
	z := &ListNode{Val: 4}
	y := &ListNode{Val: 6, Next: z}
	x := &ListNode{Val: 5, Next: y}
	val := addTwoNumbers(a, x)
	spew.Dump("val: ", val)

	//y1 := &ListNode{Val: 5}
	//x1 := &ListNode{Val: 5, Next: y1}
	//val1 := addTwoNumbers(x1, y1)
	//spew.Dump("val: ", val1)
	//
	//y2 := &ListNode{Val: 9}
	//x2 := &ListNode{Val: 9, Next: y1}
	//val2 := addTwoNumbers(x2, y2)
	//spew.Dump("val: ", val2)
	//
	//y3 := &ListNode{Val: 0}
	//x4 := &ListNode{Val: 8}
	//x3 := &ListNode{Val: 1, Next: x4}
	//val3 := addTwoNumbers(x3, y3)
	//spew.Dump("val: ", val3)
}
