package easy

// 876. Middle of the Linked List
// https://leetcode.com/problems/middle-of-the-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	middle := head
	end := head

	for end != nil && end.Next != nil {
		middle = middle.Next
		end = end.Next.Next
	}

	return middle
}
