// https://leetcode.com/problems/linked-list-cycle/description/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	tortoise := head
	hare := head

	for {
		if hare == nil || hare.Next == nil {
			return false
		}

		tortoise = tortoise.Next
		hare = hare.Next.Next

		if tortoise == hare {
			return true
		}
	}
}
