package main

//Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.
//
//Example:
//
//Input: 1->2->4, 1->3->4
//Output: 1->1->2->3->4->4

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeKLists([]*ListNode{l1, l2})
}
